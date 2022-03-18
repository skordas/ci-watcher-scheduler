package scheduleanalyzer

import (
	"time"

	"github.com/skordas/ci-watcher-scheduler/spreadsheets/engineer"
	"github.com/skordas/ci-watcher-scheduler/spreadsheets/holiday"
	"github.com/skordas/ci-watcher-scheduler/spreadsheets/schedule"
	"github.com/skordas/ci-watcher-scheduler/tools/logging"
)

const layoutUS = "1/2/2006"

// Adding one ponint to engineer activity for each assigned activity in
// CI watcher schedule.
func AddActivity(engineersMap map[string]engineer.Engineer, key string) {
	if eng, ok := engineersMap[key]; ok {
		eng.Activity++
		engineersMap[key] = eng
		logging.Debug("Adding activity for %s: now: %d", key, eng.Activity)
	}
}

// TODO instead checking all current schedule - check some static number of days
// ex. last 3 months.
// Checking current schedule for active engineers and add activity points in
// engeeners map.
func CountEngineersActivity(engineers map[string]engineer.Engineer, currentSchedule map[string]schedule.Schedule) {
	if len(currentSchedule) == 0 {
		logging.Info("No history of engineers activity!")
	} else {
		for day, scheduleForDay := range currentSchedule {
			logging.Info("Checking activity of engineers for date: %s", day)
			AddActivity(engineers, scheduleForDay.E2eWatcherY0)
			AddActivity(engineers, scheduleForDay.E2eWatcherY1)
			AddActivity(engineers, scheduleForDay.E2eWatcherY2)
			AddActivity(engineers, scheduleForDay.E2eWatcherY3)
			AddActivity(engineers, scheduleForDay.E2eWatcherY4)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY0)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY1)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY2)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY3)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY4)
		}
	}
}

// e2ey0 - e2e watcher for latest release
// e2ey1 - e2e watcher for latest - 1 release
// e2ey2 - e2e watcher for latest - 2 release
// e2ey3 - e2e watcher for latest - 3 release
// e2ey4 - e2e watcher for latest - 4 release
// upgry0 - e2e watcher for latest release
// upgry1 - e2e watcher for latest - 1 release
// upgry2 - e2e watcher for latest - 2 release
// upgry3 - e2e watcher for latest - 3 release
// upgry4 - e2e watcher for latest - 4 release
func GetWatcherFor(assignedVersion string, engineers map[string]engineer.Engineer, date string, holidays []holiday.Holiday) string {
	// coping map
	watchers := make(map[string]engineer.Engineer)
	for id, properties := range engineers {
		watchers[id] = properties
	}

	// Filtering map
	for id, properties := range engineers {
		// ignore managers
		if properties.Manager {
			logging.Debug("Removing %s - Manager is for managing, not for watching", id)
			delete(watchers, id)
			continue
		}
		// ignore not current week
		if properties.Week != weekNum(date) {
			logging.Debug("Removing %s as watcher - not current week", id)
			delete(watchers, id)
			continue
		}

		// TODO add ignoring PTO

		// ignore during holidays
		if itsHolidayInCountry(holidays, date, properties.Country) {
			logging.Debug("Romoving %s as watcher - it's holiday in %s", id, properties.Country)
			delete(watchers, id)
			continue
		}

		// ignore not assigned version
		switch assignedVersion {
		case "e2ey0":
			if !properties.E2eY0 {
				logging.Debug("Removing %s as watcher - not assigned version: E2E latest", id)
				delete(watchers, id)
				continue
			}
		case "e2ey1":
			if !properties.E2eY1 {
				logging.Debug("Removing %s as watcher - not assigned version: E2E latest - 1", id)
				delete(watchers, id)
				continue
			}
		case "e2ey2":
			if !properties.E2eY2 {
				logging.Debug("Removing %s as watcher - not assigned version: E2E latest - 2", id)
				delete(watchers, id)
				continue
			}
		case "e2ey3":
			if !properties.E2eY3 {
				logging.Debug("Removing %s as watcher - not assigned version: E2E latest - 3", id)
				delete(watchers, id)
				continue
			}
		case "e2ey4":
			if !properties.E2eY4 {
				logging.Debug("Removing %s as watcher - not assigned version: E2E latest - 4", id)
				delete(watchers, id)
				continue
			}
		case "upgry0":
			if !properties.UpgrY0 {
				logging.Debug("Removing %s as watcher - not assigned version: Upgrade latest", id)
				delete(watchers, id)
				continue
			}
		case "upgry1":
			if !properties.UpgrY1 {
				logging.Debug("Removing %s as watcher - not assigned version: Upgrade latest - 1", id)
				delete(watchers, id)
				continue
			}
		case "upgry2":
			if !properties.UpgrY2 {
				logging.Debug("Removing %s as watcher - not assigned version: Upgrade latest - 2", id)
				delete(watchers, id)
				continue
			}
		case "upgry3":
			if !properties.UpgrY3 {
				logging.Debug("Removing %s as watcher - not assigned version: Upgrade latest - 3", id)
				delete(watchers, id)
				continue
			}
		case "upgry4":
			if !properties.UpgrY4 {
				logging.Debug("Removing %s as watcher - not assigned version: Upgrade latest - 3", id)
				delete(watchers, id)
				continue
			}
		}
	}

	// sort engineers by activity
	var candidates = make(map[int][]string)
	for id, properties := range watchers {
		activity := properties.Activity
		if _, found := candidates[activity]; found {
			can := candidates[activity]
			can = append(can, id)
			candidates[activity] = can
		} else {
			can := []string{id}
			candidates[activity] = can
		}
	}

	// Getting watcher for specific version
	var theChosenOne string
	// TODO find some way, to not use hardcoded value '100' - probably number of days back * 10 (10 watchers per day)
	for i := 0; i < 100; i++ {
		if v, found := candidates[i]; found {
			theChosenOne = v[0]
			logging.Info("The Chosen One for %s: %s", assignedVersion, theChosenOne)
			break
		} else {
			continue
		}
	}
	return theChosenOne
}

func weekNum(date string) int {
	day, _ := time.Parse(layoutUS, date)
	_, weekNow := day.ISOWeek()
	logging.Debug("Week of the year: %d", weekNow)
	mod := weekNow % 2
	var weekNum int
	if mod == 1 {
		weekNum = 1
	} else {
		weekNum = 2
	}
	logging.Debug("Active week: %d", weekNum)
	return weekNum
}

// Return true if on given day and country it's a holiday
func itsHolidayInCountry(holidays []holiday.Holiday, date string, country string) bool {
	for _, h := range holidays {
		if h.Date == date && h.Country == country {
			logging.Debug("Founded holiday on %s in %s: %s", date, country, h.Name)
			return true
		}
	}
	return false
}

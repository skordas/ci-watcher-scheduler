package scheduleanalyzer

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/engineer"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/holiday"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/schedule"
)

// Adding one ponint to engineer activity for each assigned activity in
// CI watcher schedule.
func AddActivity(engineersMap map[string]engineer.Engineer, key string) {
	if eng, ok := engineersMap[key]; ok {
		eng.Activity++
		engineersMap[key] = eng
		log.WithFields(log.Fields{"engineer": key, "Activity": eng.Activity}).Debug("Adding new activity")
	}
}

// TODO instead checking all current schedule - check some static number of days
// ex. last 3 months.
// Checking current schedule for active engineers and add activity points in
// engeeners map.
func CountEngineersActivity(engineers map[string]engineer.Engineer, currentSchedule map[time.Time]schedule.Schedule) {
	if len(currentSchedule) == 0 {
		log.Info("No history of engineers activity!")
	} else {
		for day, scheduleForDay := range currentSchedule {
			log.WithFields(log.Fields{"Date": day}).Info("Checking activity of ongineers")
			AddActivity(engineers, scheduleForDay.E2eWatcherY0)
			AddActivity(engineers, scheduleForDay.E2eWatcherY1)
			AddActivity(engineers, scheduleForDay.E2eWatcherY2)
			AddActivity(engineers, scheduleForDay.E2eWatcherY3)
			AddActivity(engineers, scheduleForDay.E2eWatcherY4)
			AddActivity(engineers, scheduleForDay.E2eWatcherY5)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY0)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY1)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY2)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY3)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY4)
			AddActivity(engineers, scheduleForDay.UpgrWatcherY5)
		}
	}
}

// Values for assigned version:
// e2ey0 - e2e watcher for latest release
// e2ey1 - e2e watcher for latest - 1 release
// e2ey2 - e2e watcher for latest - 2 release
// e2ey3 - e2e watcher for latest - 3 release
// e2ey4 - e2e watcher for latest - 4 release
// e2ey5 - e2e watcher for latest - 5 release
// upgry0 - e2e watcher for latest release
// upgry1 - e2e watcher for latest - 1 release
// upgry2 - e2e watcher for latest - 2 release
// upgry3 - e2e watcher for latest - 3 release
// upgry4 - e2e watcher for latest - 4 release
// upgry5 - e2e watcher for latest - 5 release
func GetWatcherFor(assignedVersion string, engineers map[string]engineer.Engineer, date time.Time, holidays []holiday.Holiday) string {
	// coping map
	watchers := make(map[string]engineer.Engineer)
	for id, properties := range engineers {
		watchers[id] = properties
	}

	// Filtering map
	for id, properties := range engineers {
		// ignore managers
		if properties.Manager {
			log.WithField("Engineer", id).Debug("Manager is for managing, not for watching. Removing!")
			delete(watchers, id)
			continue
		}

		// ignore New to CI
		// TODO - check if engineer is a buddy - then add this engineer.
		if properties.NewToCi {
			log.WithField("Engineer", id).Debug("New To Ci - will be added with buddy. Removing!")
			delete(watchers, id)
		}

		// ignore not current week
		if properties.Week != weekNum(date) {
			log.WithField("Engineer", id).Debug("Not assigned to current week. Removing!")
			delete(watchers, id)
			continue
		}

		// TODO add ignoring PTO

		// ignore during holidays
		if itsHolidayInCountry(holidays, date, properties.Country) {
			log.WithFields(log.Fields{"Engineer": id, "Country": properties.Country}).Debug("It's Holiday for this engineery. Removing")
			delete(watchers, id)
			continue
		}

		// ignore not assigned version
		switch assignedVersion {
		case "e2ey0":
			if !properties.E2eY0 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to E2E latest. Removing!")
				delete(watchers, id)
				continue
			}
		case "e2ey1":
			if !properties.E2eY1 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to E2E latest - 1. Removing!")
				delete(watchers, id)
				continue
			}
		case "e2ey2":
			if !properties.E2eY2 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to E2E latest - 2. Removing!")
				delete(watchers, id)
				continue
			}
		case "e2ey3":
			if !properties.E2eY3 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to E2E latest - 3. Removing!")
				delete(watchers, id)
				continue
			}
		case "e2ey4":
			if !properties.E2eY4 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to E2E latest - 4. Removing!")
				delete(watchers, id)
				continue
			}
		case "e2ey5":
			if !properties.E2eY5 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to E2E latest - 5. Removing!")
				delete(watchers, id)
				continue
			}
		case "upgry0":
			if !properties.UpgrY0 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to Upgrade latest. Removing!")
				delete(watchers, id)
				continue
			}
		case "upgry1":
			if !properties.UpgrY1 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to Upgrade latest - 1. Removing!")
				delete(watchers, id)
				continue
			}
		case "upgry2":
			if !properties.UpgrY2 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to Upgrade latest - 2. Removing!")
				delete(watchers, id)
				continue
			}
		case "upgry3":
			if !properties.UpgrY3 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to Upgrade latest - 3. Removing!")
				delete(watchers, id)
				continue
			}
		case "upgry4":
			if !properties.UpgrY4 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to Upgrade latest - 4. Removing!")
				delete(watchers, id)
				continue
			}
		case "upgry5":
			if !properties.UpgrY5 {
				log.WithField("Engineer", id).Debug("Engineer not assigned to Upgrade latest - 5. Removing!")
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
			log.WithFields(log.Fields{"Assigned Version": assignedVersion, "The Chosen One": theChosenOne}).Info("We got the winner!")
			break
		} else {
			continue
		}
	}
	return theChosenOne
}

func weekNum(date time.Time) int {
	_, weekNow := date.ISOWeek()
	log.WithField("WeekNo", weekNow).Debug("Week of the year")
	mod := weekNow % 2
	var weekNum int
	if mod == 1 {
		weekNum = 1
	} else {
		weekNum = 2
	}
	log.WithField("Watch week", weekNum).Debug("Active week.")
	return weekNum
}

// Return true if on given day and country it's a holiday
func itsHolidayInCountry(holidays []holiday.Holiday, date time.Time, country string) bool {
	for _, h := range holidays {
		if h.Date == date && h.Country == country {
			log.WithFields(log.Fields{"Date": date, "Holiday name": h.Name}).Debug("Founded holiday!")
			return true
		}
	}
	return false
}

func GetDayToSchedule(schedule map[time.Time]schedule.Schedule) time.Time {
	dayToSchedule := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	for _, day := range schedule {
		if day.Date.After(dayToSchedule) {
			dayToSchedule = day.Date
		}
	}
	return dayToSchedule.Add(24 * time.Hour)
}

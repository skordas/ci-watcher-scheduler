package main

import (
	"os"
	"time"

	"github.com/skordas/ci-watcher-scheduler/spreadsheets"
	"github.com/skordas/ci-watcher-scheduler/spreadsheets/engineer"
	"github.com/skordas/ci-watcher-scheduler/spreadsheets/holiday"
	"github.com/skordas/ci-watcher-scheduler/spreadsheets/schedule"
	"github.com/skordas/ci-watcher-scheduler/tools/logging"
)

const layoutUS = "1/2/2006"

var currentSchedule = make(map[string]schedule.Schedule)
var engineers = make(map[string]engineer.Engineer)
var holidays = []holiday.Holiday{}

func main() {
	dayToSchedule := os.Getenv("DATE")
	engineers = spreadsheets.GetEngineers()
	holidays = spreadsheets.GetHolidays()
	currentSchedule = spreadsheets.GetCurrentSchedule()

	// TODO - move this outside main func
	// Counting activity of engineers

	if len(currentSchedule) == 0 {
		logging.Info("No history of engineers activity!")
	} else {
		for day, scheduleForDay := range currentSchedule {
			logging.Info("Checking activity of engineers for date: %s", day)
			addActivity(engineers, scheduleForDay.E2eWatcherY0)
			addActivity(engineers, scheduleForDay.E2eWatcherY1)
			addActivity(engineers, scheduleForDay.E2eWatcherY2)
			addActivity(engineers, scheduleForDay.E2eWatcherY3)
			addActivity(engineers, scheduleForDay.E2eWatcherY4)
			addActivity(engineers, scheduleForDay.UpgrWatcherY0)
			addActivity(engineers, scheduleForDay.UpgrWatcherY1)
			addActivity(engineers, scheduleForDay.UpgrWatcherY2)
			addActivity(engineers, scheduleForDay.UpgrWatcherY3)
			addActivity(engineers, scheduleForDay.UpgrWatcherY4)
		}
	}

	logging.Info("------ Creating schedule for date: %s ------", dayToSchedule)
	e2ey0Watcher := getWatcherFor("e2ey0", dayToSchedule)
	addActivity(engineers, e2ey0Watcher)
	e2ey1Watcher := getWatcherFor("e2ey1", dayToSchedule)
	addActivity(engineers, e2ey1Watcher)
	e2ey2Watcher := getWatcherFor("e2ey2", dayToSchedule)
	addActivity(engineers, e2ey2Watcher)
	e2ey3Watcher := getWatcherFor("e2ey3", dayToSchedule)
	addActivity(engineers, e2ey3Watcher)
	e2ey4Watcher := getWatcherFor("e2ey4", dayToSchedule)
	addActivity(engineers, e2ey4Watcher)
	upgry0Watcher := getWatcherFor("upgry0", dayToSchedule)
	addActivity(engineers, upgry0Watcher)
	upgry1Watcher := getWatcherFor("upgry1", dayToSchedule)
	addActivity(engineers, upgry1Watcher)
	upgry2Watcher := getWatcherFor("upgry2", dayToSchedule)
	addActivity(engineers, upgry2Watcher)
	upgry3Watcher := getWatcherFor("upgry3", dayToSchedule)
	addActivity(engineers, upgry3Watcher)
	upgry4Watcher := getWatcherFor("upgry4", dayToSchedule)
	addActivity(engineers, upgry4Watcher)

	scheduleToStore := schedule.New(dayToSchedule, "", e2ey0Watcher, e2ey1Watcher,
		e2ey2Watcher, e2ey3Watcher, e2ey4Watcher, upgry0Watcher, upgry1Watcher,
		upgry2Watcher, upgry3Watcher, upgry4Watcher)

	spreadsheets.StoreSchedule(scheduleToStore)
}

func addActivity(engineersMap map[string]engineer.Engineer, key string) {
	if eng, ok := engineersMap[key]; ok {
		eng.Activity++
		engineersMap[key] = eng
		logging.Debug("Adding activity for %s: now: %d", key, eng.Activity)
	}
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
func getWatcherFor(assignedVersion string, date string) string {
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
		// TODO add ignoring Holidays

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
	for i := 0; i < len(candidates); i++ {
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

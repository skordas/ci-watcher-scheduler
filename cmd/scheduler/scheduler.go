package main

import (
	"os"

	sa "github.com/skordas/ci-watcher-scheduler/internal/scheduleanalyzer"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/engineer"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/holiday"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/schedule"
	"github.com/skordas/ci-watcher-scheduler/tools/logging"
)

var currentSchedule = make(map[string]schedule.Schedule)
var engineers = make(map[string]engineer.Engineer)
var holidays = []holiday.Holiday{}

func main() {
	dayToSchedule := os.Getenv("DATE")
	engineers = spreadsheets.GetEngineers()
	holidays = spreadsheets.GetHolidays()
	currentSchedule = spreadsheets.GetCurrentSchedule()

	// Counting activity of engineers
	sa.CountEngineersActivity(engineers, currentSchedule)

	logging.Info("------ Creating schedule for date: %s ------", dayToSchedule)
	e2ey0Watcher := sa.GetWatcherFor("e2ey0", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey0Watcher)
	e2ey1Watcher := sa.GetWatcherFor("e2ey1", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey1Watcher)
	e2ey2Watcher := sa.GetWatcherFor("e2ey2", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey2Watcher)
	e2ey3Watcher := sa.GetWatcherFor("e2ey3", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey3Watcher)
	e2ey4Watcher := sa.GetWatcherFor("e2ey4", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey4Watcher)
	e2ey5Watcher := sa.GetWatcherFor("e2ey5", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey5Watcher)
	upgry0Watcher := sa.GetWatcherFor("upgry0", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry0Watcher)
	upgry1Watcher := sa.GetWatcherFor("upgry1", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry1Watcher)
	upgry2Watcher := sa.GetWatcherFor("upgry2", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry2Watcher)
	upgry3Watcher := sa.GetWatcherFor("upgry3", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry3Watcher)
	upgry4Watcher := sa.GetWatcherFor("upgry4", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry4Watcher)
	upgry5Watcher := sa.GetWatcherFor("upgry5", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry5Watcher)

	// store schedule in spreadsheet
	scheduleToStore := schedule.New(dayToSchedule, "", e2ey0Watcher, e2ey1Watcher,
		e2ey2Watcher, e2ey3Watcher, e2ey4Watcher, e2ey5Watcher, upgry0Watcher,
		upgry1Watcher, upgry2Watcher, upgry3Watcher, upgry4Watcher, upgry5Watcher)

	spreadsheets.StoreSchedule(scheduleToStore)
}
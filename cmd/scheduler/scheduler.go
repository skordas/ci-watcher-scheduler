package main

import (
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	sa "github.com/skordas/ci-watcher-scheduler/internal/scheduleanalyzer"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/engineer"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/holiday"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/schedule"
)

var currentSchedule = make(map[time.Time]schedule.Schedule)
var engineers = make(map[string]engineer.Engineer)
var holidays = []holiday.Holiday{}

func init() {
	if os.Getenv("DEBUG") == "true" {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	engineers = spreadsheets.GetEngineers()
	holidays = spreadsheets.GetHolidays()
	currentSchedule = spreadsheets.GetCurrentSchedule()
	dayToSchedule := sa.GetDayToSchedule(currentSchedule)

	// Loading weights
	var e2eY0WatcherWeight, _ = strconv.Atoi(os.Getenv("E2E_WATCHER_Y0_WEIGHT"))
	var e2eY1WatcherWeight, _ = strconv.Atoi(os.Getenv("E2E_WATCHER_Y1_WEIGHT"))
	var e2eY2WatcherWeight, _ = strconv.Atoi(os.Getenv("E2E_WATCHER_Y2_WEIGHT"))
	var e2eY3WatcherWeight, _ = strconv.Atoi(os.Getenv("E2E_WATCHER_Y3_WEIGHT"))
	var e2eY4WatcherWeight, _ = strconv.Atoi(os.Getenv("E2E_WATCHER_Y4_WEIGHT"))
	var e2eY5WatcherWeight, _ = strconv.Atoi(os.Getenv("E2E_WATCHER_Y5_WEIGHT"))
	var upgrY0WatcherWeight, _ = strconv.Atoi(os.Getenv("UPGR_WATCHER_Y0_WEIGHT"))
	var upgrY1WatcherWeight, _ = strconv.Atoi(os.Getenv("UPGR_WATCHER_Y1_WEIGHT"))
	var upgrY2WatcherWeight, _ = strconv.Atoi(os.Getenv("UPGR_WATCHER_Y2_WEIGHT"))
	var upgrY3WatcherWeight, _ = strconv.Atoi(os.Getenv("UPGR_WATCHER_Y3_WEIGHT"))
	var upgrY4WatcherWeight, _ = strconv.Atoi(os.Getenv("UPGR_WATCHER_Y4_WEIGHT"))
	var upgrY5WatcherWeight, _ = strconv.Atoi(os.Getenv("UPGR_WATCHER_Y5_WEIGHT"))

	// Counting activity of engineers
	sa.CountEngineersActivity(engineers, currentSchedule)

	log.WithField("dayToSchedule", dayToSchedule).Info("Creating schedule")
	e2ey0Watcher := sa.GetWatcherFor("e2ey0", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey0Watcher, e2eY0WatcherWeight)
	e2ey1Watcher := sa.GetWatcherFor("e2ey1", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey1Watcher, e2eY1WatcherWeight)
	e2ey2Watcher := sa.GetWatcherFor("e2ey2", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey2Watcher, e2eY2WatcherWeight)
	e2ey3Watcher := sa.GetWatcherFor("e2ey3", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey3Watcher, e2eY3WatcherWeight)
	e2ey4Watcher := sa.GetWatcherFor("e2ey4", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey4Watcher, e2eY4WatcherWeight)
	e2ey5Watcher := sa.GetWatcherFor("e2ey5", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, e2ey5Watcher, e2eY5WatcherWeight)
	upgry0Watcher := sa.GetWatcherFor("upgry0", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry0Watcher, upgrY0WatcherWeight)
	upgry1Watcher := sa.GetWatcherFor("upgry1", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry1Watcher, upgrY1WatcherWeight)
	upgry2Watcher := sa.GetWatcherFor("upgry2", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry2Watcher, upgrY2WatcherWeight)
	upgry3Watcher := sa.GetWatcherFor("upgry3", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry3Watcher, upgrY3WatcherWeight)
	upgry4Watcher := sa.GetWatcherFor("upgry4", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry4Watcher, upgrY4WatcherWeight)
	upgry5Watcher := sa.GetWatcherFor("upgry5", engineers, dayToSchedule, holidays)
	sa.AddActivity(engineers, upgry5Watcher, upgrY5WatcherWeight)

	// store schedule in spreadsheet
	scheduleToStore := schedule.New(dayToSchedule.Format(schedule.LayoutUS), "", e2ey0Watcher, e2ey1Watcher,
		e2ey2Watcher, e2ey3Watcher, e2ey4Watcher, e2ey5Watcher, upgry0Watcher,
		upgry1Watcher, upgry2Watcher, upgry3Watcher, upgry4Watcher, upgry5Watcher)

	spreadsheets.StoreSchedule(scheduleToStore)
}

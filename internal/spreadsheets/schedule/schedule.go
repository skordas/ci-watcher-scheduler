package schedule

import (
	log "github.com/sirupsen/logrus"
)

type Schedule struct {
	Date          string //  0 - Scheduled day
	Manager       string //  1 - ID of manager for week
	E2eWatcherY0  string //  2 - E2E CI-Watcher for v. 4.y
	E2eWatcherY1  string //  3 - E2E CI-Watcher for v. 4.y-1
	E2eWatcherY2  string //  4 - E2E CI-Watcher for v. 4.y-2
	E2eWatcherY3  string //  5 - E2E CI-Watcher for v. 4.y-3
	E2eWatcherY4  string //  6 - E2E CI-Watcher for v. 4.y-4
	E2eWatcherY5  string //  7 - E2E CI-Watcher for v. 4.y-5
	UpgrWatcherY0 string //  8 - Upgrade CI-Watcher for v. 4.y
	UpgrWatcherY1 string //  9 - Upgrade CI-Watcher for v. 4.y-1
	UpgrWatcherY2 string // 10 - Upgrade CI-Watcher for v. 4.y-2
	UpgrWatcherY3 string // 11 - Upgrade CI-Watcher for v. 4.y-3
	UpgrWatcherY4 string // 12 - Upgrade CI-Watcher for v. 4.y-4
	UpgrWatcherY5 string // 13 - Upgrade CI-Watcher for v. 4.y-5
}

func New(date interface{}, manager interface{}, e2eWatcherY0 interface{},
	e2eWatcherY1 interface{}, e2eWatcherY2 interface{},
	e2eWatcherY3 interface{}, e2eWatcherY4 interface{},
	e2eWatcherY5 interface{}, upgradeWatcherY0 interface{},
	upgradeWatcherY1 interface{}, upgradeWatcherY2 interface{},
	upgradeWatcherY3 interface{}, upgradeWatcherY4 interface{},
	upgradeWatcherY5 interface{}) Schedule {

	sch := Schedule{
		date.(string),
		manager.(string),
		e2eWatcherY0.(string),
		e2eWatcherY1.(string),
		e2eWatcherY2.(string),
		e2eWatcherY3.(string),
		e2eWatcherY4.(string),
		e2eWatcherY5.(string),
		upgradeWatcherY0.(string),
		upgradeWatcherY1.(string),
		upgradeWatcherY2.(string),
		upgradeWatcherY3.(string),
		upgradeWatcherY4.(string),
		upgradeWatcherY5.(string),
	}

	log.WithFields(log.Fields{
		"Date":                       sch.Date,
		"Manager":                    sch.Manager,
		"E2E Watcher latest":         sch.E2eWatcherY0,
		"E2E Watcher latest - 1":     sch.E2eWatcherY1,
		"E2E Watcher latest - 2":     sch.E2eWatcherY2,
		"E2E Watcher latest - 3":     sch.E2eWatcherY3,
		"E2E Watcher latest - 4":     sch.E2eWatcherY4,
		"E2E Watcher latest - 5":     sch.E2eWatcherY5,
		"Upgrade Watche latest":      sch.UpgrWatcherY0,
		"Upgrade Watcher latest - 1": sch.UpgrWatcherY1,
		"Upgrade Watcher latest - 2": sch.UpgrWatcherY2,
		"Upgrade Watcher latest - 3": sch.UpgrWatcherY3,
		"Upgrade Watcher latest - 4": sch.UpgrWatcherY4,
		"Upgrade Watcher latest - 5": sch.UpgrWatcherY5,
	}).Debug("Creating a new schedule")

	return sch
}

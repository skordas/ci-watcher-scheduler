package spreadsheets

import "time"

type schedule struct {
	Date          time.Date()  // Scheduled day
	Manager       string       // ID of manager for week
	E2eWatcherY0  string       // E2E CI-Watcher for v. 4.y
	E2eWatcherY1  string       // E2E CI-Watcher for v. 4.y-1
	E2eWatcherY2  string       // E2E CI-Watcher for v. 4.y-2
	E2eWatcherY3  string       // E2E CI-Watcher for v. 4.y-3
	E2eWatcherY4  string       // E2E CI-Watcher for v. 4.y-4
	UpgrWatcherY0 string       // Upgrade CI-Watcher for v. 4.y
	UpgrWatcherY1 string       // Upgrade CI-Watcher for v. 4.y-1
	UpgrWatcherY2 string       // Upgrade CI-Watcher for v. 4.y-2
	UpgrWatcherY3 string       // Upgrade CI-Watcher for v. 4.y-3
	UpgrWatcherY4 string       // Upgrade CI-Watcher for v. 4.y-4
}

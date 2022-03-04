package engineer

type engineer struct {
	kerberos     string // Kebreros ID
	team         string // Team
	country      string // Country
	week         int    // Number of week
	manager      bool   // true for managers
	releaseLead  string // true for release leaders
	teamLead     bool   // true fdo team leaders
	newToCi      bool   // true if CI-Watcher role is a new thing for engineer
	newToCiBuddy string // Kerberos of engineer which will be helping
	e2eY0        bool   // True for CI-watcher for E2E v. 4.y
	e2eY1        bool   // True for CI-watcher for E2E v. 4.y-1
	e2eY2        bool   // True for CI-watcher for E2E v. 4.y-2
	e2eY3        bool   // True for CI-watcher for E2E v. 4.y-3
	e2eY4        bool   // True for CI-watcher for E2E v. 4.y-4
	upgrY0       bool   // True for CI-watcher for Upgrade v. 4.y
	upgrY1       bool   // True for CI-watcher for Upgrade v. 4.y-1
	upgrY2       bool   // True for CI-watcher for Upgrade v. 4.y-2
	upgrY3       bool   // True for CI-watcher for Upgrade v. 4.y-3
	upgrY4       bool   // True for CI-watcher for Upgrade v. 4.y-4
}

func New(kerberos string, team string, country string, week int, manager bool, releaseLead string, teamLead bool, newToCi bool, newToCiBuddy string, e2eY0 bool, e2eY1 bool, e2eY2 bool, e2eY3 bool, e2eY4 bool, upgrY0 bool, upgrY1 bool, upgrY2 bool, upgrY3 bool, upgrY4 bool) engineer {
	e := engineer {
		kerberos,
		team,
		country,
		week,
		manager,
		releaseLead,
		teamLead,
		newToCi,
		newToCiBuddy,
		e2eY0,
		e2eY1,
		e2eY2,
		e2eY3,
		e2eY4,
		upgrY0,
		upgrY1,
		upgrY2,
		upgrY3,
		upgrY4
	}
	return e
}

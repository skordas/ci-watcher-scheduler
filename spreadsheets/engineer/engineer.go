package engineer

import (
	"strconv"

	"github.com/skordas/ci-watcher-scheduler/tools/logging"
)

type Engineer struct {
	Kerberos     string //  0 - Kebreros ID
	Team         string //  1 - Team
	Country      string //  2 - Country
	Week         int    //  3 - Number of week
	Manager      bool   //  4 - true for managers
	ReleaseLead  bool   //  5 - true for release leaders
	TeamLead     bool   //  6 - true fdo team leaders
	NewToCi      bool   //  7 - true if CI-Watcher role is a new thing for engineer
	NewToCiBuddy string //  8 - Kerberos of engineer which will be helping
	E2eY0        bool   //  9 - True for CI-watcher for E2E v. 4.y
	E2eY1        bool   // 10 - True for CI-watcher for E2E v. 4.y-1
	E2eY2        bool   // 11 - True for CI-watcher for E2E v. 4.y-2
	E2eY3        bool   // 12 - True for CI-watcher for E2E v. 4.y-3
	E2eY4        bool   // 13 - True for CI-watcher for E2E v. 4.y-4
	UpgrY0       bool   // 14 - True for CI-watcher for Upgrade v. 4.y
	UpgrY1       bool   // 15 - True for CI-watcher for Upgrade v. 4.y-1
	UpgrY2       bool   // 16 - True for CI-watcher for Upgrade v. 4.y-2
	UpgrY3       bool   // 17 - True for CI-watcher for Upgrade v. 4.y-3
	UpgrY4       bool   // 18 - True for CI-watcher for Upgrade v. 4.y-4
}

func New(kerberos interface{}, team interface{}, country interface{},
	week interface{}, manager interface{}, releaseLead interface{},
	teamLead interface{}, newToCi interface{}, newToCiBuddy interface{},
	e2eY0 interface{}, e2eY1 interface{}, e2eY2 interface{}, e2eY3 interface{},
	e2eY4 interface{}, upgrY0 interface{}, upgrY1 interface{},
	upgrY2 interface{}, upgrY3 interface{}, upgrY4 interface{}) Engineer {

	e := Engineer{
		kerberos.(string),
		team.(string),
		country.(string),
		parseInt(week.(string)),
		parseBool(manager.(string)),
		parseBool(releaseLead.(string)),
		parseBool(teamLead.(string)),
		parseBool(newToCi.(string)),
		newToCiBuddy.(string),
		parseBool(e2eY0.(string)),
		parseBool(e2eY1.(string)),
		parseBool(e2eY2.(string)),
		parseBool(e2eY3.(string)),
		parseBool(e2eY4.(string)),
		parseBool(upgrY0.(string)),
		parseBool(upgrY1.(string)),
		parseBool(upgrY2.(string)),
		parseBool(upgrY3.(string)),
		parseBool(upgrY4.(string)),
	}

	logging.Debug("------ Getting engineer from spreadsheet ------")
	logging.Debug("Kerberos: %s", e.Kerberos)
	logging.Debug("Team: %s", e.Team)
	logging.Debug("Country: %s", e.Country)
	logging.Debug("Week: %d", e.Week)
	logging.Debug("Manager: %t", e.Manager)
	logging.Debug("Release Lead: %t", e.ReleaseLead)
	logging.Debug("Team Lead: %t", e.TeamLead)
	logging.Debug("New to CI: %t", e.NewToCi)
	logging.Debug("New to CI buddy: %s", e.NewToCiBuddy)
	logging.Debug("E2E latest: %t", e.E2eY0)
	logging.Debug("E2E latest - 1: %t", e.E2eY1)
	logging.Debug("E2E latest - 2: %t", e.E2eY2)
	logging.Debug("E2E latest - 3: %t", e.E2eY3)
	logging.Debug("E2E latest - 4: %t", e.E2eY4)
	logging.Debug("Upgrade latest: %t", e.UpgrY0)
	logging.Debug("Upgrade latest - 1: %t", e.UpgrY1)
	logging.Debug("Upgrade latest - 2: %t", e.UpgrY2)
	logging.Debug("Upgrade latest - 3: %t", e.UpgrY3)
	logging.Debug("Upgrade latest - 4: %t", e.UpgrY4)

	return e
}

// parseInt is needed to deal with empty fields.
func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func parseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}

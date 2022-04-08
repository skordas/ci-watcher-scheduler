package engineer

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Engineer struct {
	Kerberos     string //  0 - Kebreros ID
	Team         string //  1 - Team
	Country      string //  2 - Country
	Week         int    //  3 - Number of week
	Manager      bool   //  4 - true for managers
	ReleaseLead  bool   //  5 - true for release leaders
	TeamLead     bool   //  6 - true for team leaders
	NewToCi      bool   //  7 - true if CI-Watcher role is a new thing for engineer
	NewToCiBuddy string //  8 - Kerberos of engineer which will be helping
	E2eY0        bool   //  9 - True for CI-watcher for E2E v. 4.y
	E2eY1        bool   // 10 - True for CI-watcher for E2E v. 4.y-1
	E2eY2        bool   // 11 - True for CI-watcher for E2E v. 4.y-2
	E2eY3        bool   // 12 - True for CI-watcher for E2E v. 4.y-3
	E2eY4        bool   // 13 - True for CI-watcher for E2E v. 4.y-4
	E2eY5        bool   // 14 - True for CI-watcher for E2E v. 4.y-5
	UpgrY0       bool   // 15 - True for CI-watcher for Upgrade v. 4.y
	UpgrY1       bool   // 16 - True for CI-watcher for Upgrade v. 4.y-1
	UpgrY2       bool   // 17 - True for CI-watcher for Upgrade v. 4.y-2
	UpgrY3       bool   // 18 - True for CI-watcher for Upgrade v. 4.y-3
	UpgrY4       bool   // 19 - True for CI-watcher for Upgrade v. 4.y-4
	UpgrY5       bool   // 20 - True for CI-watcher for Upgrade v. 4.y-5
	Activity     int    // 21 - Number of times Engineer was working as Watcher
}

func New(kerberos interface{}, team interface{}, country interface{},
	week interface{}, manager interface{}, releaseLead interface{},
	teamLead interface{}, newToCi interface{}, newToCiBuddy interface{},
	e2eY0 interface{}, e2eY1 interface{}, e2eY2 interface{}, e2eY3 interface{},
	e2eY4 interface{}, e2ey5 interface{}, upgrY0 interface{},
	upgrY1 interface{}, upgrY2 interface{}, upgrY3 interface{},
	upgrY4 interface{}, upgrY5 interface{}) Engineer {

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
		parseBool(e2ey5.(string)),
		parseBool(upgrY0.(string)),
		parseBool(upgrY1.(string)),
		parseBool(upgrY2.(string)),
		parseBool(upgrY3.(string)),
		parseBool(upgrY4.(string)),
		parseBool(upgrY5.(string)),
		0,
	}

	log.WithFields(log.Fields{
		"Kerberos":           e.Kerberos,
		"Team":               e.Team,
		"Country":            e.Country,
		"Week":               e.Week,
		"Manager":            e.Manager,
		"Release Lead":       e.ReleaseLead,
		"Team Lead":          e.TeamLead,
		"New to CI":          e.NewToCi,
		"New to CI Buddy":    e.NewToCiBuddy,
		"E2E latest":         e.E2eY0,
		"E2E latest - 1":     e.E2eY1,
		"E2E latest - 2":     e.E2eY2,
		"E2E latest - 3":     e.E2eY3,
		"E2E latest - 4":     e.E2eY4,
		"E2E latest - 5":     e.E2eY5,
		"Upgrade latest":     e.UpgrY0,
		"Upgrade latest - 1": e.UpgrY1,
		"Upgrade latest - 2": e.UpgrY2,
		"Upgrade latest - 3": e.UpgrY3,
		"Upgrade latest - 4": e.UpgrY4,
		"Upgrade latest - 5": e.UpgrY5,
		"Activity":           e.Activity,
	}).Debug("Getting engineer from spreadsheet")

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

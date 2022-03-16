package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"os"
	"time"

	"github.com/skordas/ci-watcher-scheduler/spreadsheets/engineer"
	"github.com/skordas/ci-watcher-scheduler/spreadsheets/holiday"
	"github.com/skordas/ci-watcher-scheduler/spreadsheets/schedule"
	"github.com/skordas/ci-watcher-scheduler/tools/logging"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const layoutUS = "1/2/2006"

var scheduleCurrent = make(map[string]schedule.Schedule)
var engineers = make(map[string]engineer.Engineer)
var holidays = []holiday.Holiday{}

func main() {
	// Get Environment Variables
	credentialsJson := os.Getenv("CREDENTIALS")
	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	dayToSchedule := os.Getenv("DATE")
	// TODO move sheets ranges to some properties file
	engineersRange := "Engineers!A2:S"
	holidaysRange := "Holidays!A2:C"
	scheduleRange := "CI_Watch_Schedule!A2:L"

	ctx := context.Background()
	credentials, err := ioutil.ReadFile(credentialsJson)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(credentials, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := config.Client(ctx)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrive Sheets client: %v", err)
	}

	// TODO - move getting engineers outside main func.
	// Getting engineers
	logging.Info("------ Getting list of engineers ------")
	respEngineers, err := srv.Spreadsheets.Values.Get(spreadsheetId, engineersRange).Do()
	if err != nil {
		log.Fatal("Unable to retrieve data from sheet: %v", err)
	}

	if len(respEngineers.Values) == 0 {
		logging.Warning("No data found in range: %s", engineersRange)
	} else {
		for _, row := range respEngineers.Values {
			e := engineer.New(row[0], row[1], row[2], row[3], row[4], row[5], row[6],
				row[7], row[8], row[9], row[10], row[11], row[12], row[13],
				row[14], row[15], row[16], row[17], row[18])
			engineers[e.Kerberos] = e
		}
	}

	// TODO - move getting holidays outside main func.
	// Getting holidays
	logging.Info("------ Getting list of holidays ------")
	respHolidays, err := srv.Spreadsheets.Values.Get(spreadsheetId, holidaysRange).Do()
	if err != nil {
		log.Fatal("Unable to retrive data from sheet: %v", err)
	}

	if len(respHolidays.Values) == 0 {
		logging.Warning("No data found in range: %s", holidaysRange)
	} else {
		for _, row := range respHolidays.Values {
			h := holiday.New(row[0], row[1], row[2])
			holidays = append(holidays, h)
		}
	}

	// TODO - move getting scheduler outside main func
	// TODO - getting schedule only specific number of days back (ex. last 30 days)
	// Getting current schedule.
	// TODO - for now date as a string - later as a date
	logging.Info("------ Getting current schedule ------")
	respSchedule, err := srv.Spreadsheets.Values.Get(spreadsheetId, scheduleRange).Do()
	if err != nil {
		log.Fatal("Unable to retrive data from sheet: %v", err)
	}

	if len(respSchedule.Values) == 0 {
		logging.Warning("No data found in range :s", scheduleRange)
	} else {
		for _, row := range respSchedule.Values {
			sch := schedule.New(row[0], row[1], row[2], row[3], row[4], row[5],
				row[6], row[7], row[8], row[9], row[10], row[11])
			scheduleCurrent[sch.Date] = sch
		}
	}

	// TODO - move this outside main func
	// Counting activity of engineers

	if len(scheduleCurrent) == 0 {
		logging.Info("No history of engineers activity!")
	} else {
		for day, scheduleForDay := range scheduleCurrent {
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
	fmt.Printf("Right now: %s : %d\n", engineers[e2ey0Watcher].Kerberos, engineers[e2ey0Watcher].Activity)
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

	// store in spreadsheet
	var vr sheets.ValueRange
	myval := []interface{}{
		scheduleToStore.Date,
		scheduleToStore.Manager,
		scheduleToStore.E2eWatcherY0,
		scheduleToStore.E2eWatcherY1,
		scheduleToStore.E2eWatcherY2,
		scheduleToStore.E2eWatcherY3,
		scheduleToStore.E2eWatcherY4,
		scheduleToStore.UpgrWatcherY0,
		scheduleToStore.UpgrWatcherY1,
		scheduleToStore.UpgrWatcherY2,
		scheduleToStore.UpgrWatcherY3,
		scheduleToStore.UpgrWatcherY4,
	}
	vr.Values = append(vr.Values, myval)
	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, "CI_Watch_Schedule!A1", &vr).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		logging.Error("Unable to store data in sheet. %v", err)
	}
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

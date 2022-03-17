package spreadsheets

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/skordas/ci-watcher-scheduler/spreadsheets/engineer"
	"github.com/skordas/ci-watcher-scheduler/spreadsheets/holiday"
	"github.com/skordas/ci-watcher-scheduler/spreadsheets/schedule"
	"github.com/skordas/ci-watcher-scheduler/tools/logging"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var credentialsJson string
var spreadsheetId string
var engineersRange string = "Engineers!A2:S"
var holidaysRange string = "Holidays!A2:C"
var scheduleRange string = "CI_Watch_Schedule!A2:L"

var srv *sheets.Service
var initiate bool = true

func initiateSrv() {
	if initiate {
		logging.Info("------ Staring sheets client ------")
		credentialsJson = os.Getenv("CREDENTIALS")
		logging.Info("Credentials path: %s", credentialsJson)

		spreadsheetId = os.Getenv("SPREADSHEET_ID")
		logging.Info("Spreadsheet ID: %s", spreadsheetId)

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
		srv, err = sheets.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Unable to retrive Sheets client: %v", err)
		}
		initiate = false
	}
}

func GetEngineers() map[string]engineer.Engineer {
	initiateSrv()

	logging.Info("------ Getting list of engineers ------")
	engineers := make(map[string]engineer.Engineer)
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, engineersRange).Do()
	if err != nil {
		log.Fatal("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		logging.Warning("No data found in range: %s", engineersRange)
	} else {
		for _, row := range resp.Values {
			e := engineer.New(row[0], row[1], row[2], row[3], row[4], row[5], row[6],
				row[7], row[8], row[9], row[10], row[11], row[12], row[13],
				row[14], row[15], row[16], row[17], row[18])
			engineers[e.Kerberos] = e
		}
	}
	return engineers
}

func GetHolidays() []holiday.Holiday {
	initiateSrv()

	logging.Info("------ Getting list of holidays ------")
	var holidays = []holiday.Holiday{}
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, holidaysRange).Do()
	if err != nil {
		log.Fatal("Unable to retrive data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		logging.Warning("No data found in range: %s", holidaysRange)
	} else {
		for _, row := range resp.Values {
			h := holiday.New(row[0], row[1], row[2])
			holidays = append(holidays, h)
		}
	}
	return holidays
}

func GetCurrentSchedule() map[string]schedule.Schedule {
	initiateSrv()

	logging.Info("------ Getting current schedule ------")
	var scheduleCurrent = make(map[string]schedule.Schedule)
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, scheduleRange).Do()
	if err != nil {
		log.Fatal("Unable to retrive data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		logging.Warning("No data found in range :s", scheduleRange)
	} else {
		for _, row := range resp.Values {
			sch := schedule.New(row[0], row[1], row[2], row[3], row[4], row[5],
				row[6], row[7], row[8], row[9], row[10], row[11])
			scheduleCurrent[sch.Date] = sch
		}
	}
	return scheduleCurrent
}

func StoreSchedule(schedule schedule.Schedule) {
	initiateSrv()

	var vr sheets.ValueRange
	myval := []interface{}{
		schedule.Date,
		schedule.Manager,
		schedule.E2eWatcherY0,
		schedule.E2eWatcherY1,
		schedule.E2eWatcherY2,
		schedule.E2eWatcherY3,
		schedule.E2eWatcherY4,
		schedule.UpgrWatcherY0,
		schedule.UpgrWatcherY1,
		schedule.UpgrWatcherY2,
		schedule.UpgrWatcherY3,
		schedule.UpgrWatcherY4,
	}
	vr.Values = append(vr.Values, myval)
	_, err := srv.Spreadsheets.Values.Append(spreadsheetId, "CI_Watch_Schedule!A1", &vr).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		logging.Error("Unable to store data in sheet. %v", err)
	}
}
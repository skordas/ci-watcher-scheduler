package spreadsheets

import (
	"context"
	"io/ioutil"
	"os"
	"time"

	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/engineer"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/holiday"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/schedule"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// inputOption is used by ValueInputOption in spreadsheets.
// Values saved in spreadsheet cells will be interpreted like entered by user
// string: 3/14/2022 after insert will be automatically format to dates
const inputOption string = "USER_ENTERED"
const layoutUS = "1/2/2006"

var credentialsJson string
var spreadsheetId string

// TODO - Not use hardcoded values. Need to decide how we want to pass this
// values. os ENV or some file.
var engineersRange string = "Engineers!A2:U"
var holidaysRange string = "Holidays!A2:C"
var scheduleRange string = "CI_Watch_Schedule!A2:N"
var insertScheduleRange string = "CI_Watch_Schedule!A1:N1"

var srv *sheets.Service
var initiate bool = true

// InitiateSrv is getting all informations needed to start connection
// with spreadsheets
func initiateSrv() {
	if initiate {
		log.Info("------ Staring sheets client ------")
		credentialsJson = os.Getenv("CREDENTIALS")
		log.WithField("path", credentialsJson).Info("Starting with credentials:")

		spreadsheetId = os.Getenv("SPREADSHEET_ID")
		log.WithField("SpreadsheetID", spreadsheetId).Info("Working with spreadsheet:")

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

// Get list of engineers from spreadsheet
func GetEngineers() map[string]engineer.Engineer {
	initiateSrv()

	log.Info("------ Getting list of engineers ------")
	engineers := make(map[string]engineer.Engineer)
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, engineersRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		log.Warnf("No data found in range: %s", engineersRange)
	} else {
		for _, row := range resp.Values {
			e := engineer.New(row[0], row[1], row[2], row[3], row[4], row[5],
				row[6], row[7], row[8], row[9], row[10], row[11], row[12],
				row[13], row[14], row[15], row[16], row[17], row[18], row[19],
				row[20])
			engineers[e.Kerberos] = e
		}
	}
	return engineers
}

// Get list of holidays from spreadsheet
func GetHolidays() []holiday.Holiday {
	initiateSrv()

	log.Info("------ Getting list of holidays ------")
	var holidays = []holiday.Holiday{}
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, holidaysRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrive data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		log.Warnf("No data found in range: %s", holidaysRange)
	} else {
		for _, row := range resp.Values {
			h := holiday.New(row[0], row[1], row[2])
			holidays = append(holidays, h)
		}
	}
	return holidays
}

// Get current schedule from spreadsheet
func GetCurrentSchedule() map[time.Time]schedule.Schedule {
	initiateSrv()

	log.Info("------ Getting current schedule ------")
	var scheduleCurrent = make(map[time.Time]schedule.Schedule)
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, scheduleRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrive data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		log.Warnf("No data found in range %s", scheduleRange)
	} else {
		for _, row := range resp.Values {
			sch := schedule.New(row[0], row[1], row[2], row[3], row[4], row[5],
				row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13])
			scheduleCurrent[sch.Date] = sch
		}
	}
	return scheduleCurrent
}

// store a new line is schedule spreadsheet
func StoreSchedule(schedule schedule.Schedule) {
	initiateSrv()
	insertEmptyRow()

	var vr sheets.ValueRange
	myval := []interface{}{
		schedule.Date.Format(layoutUS),
		schedule.Manager,
		schedule.E2eWatcherY0,
		schedule.E2eWatcherY1,
		schedule.E2eWatcherY2,
		schedule.E2eWatcherY3,
		schedule.E2eWatcherY4,
		schedule.E2eWatcherY5,
		schedule.UpgrWatcherY0,
		schedule.UpgrWatcherY1,
		schedule.UpgrWatcherY2,
		schedule.UpgrWatcherY3,
		schedule.UpgrWatcherY4,
		schedule.UpgrWatcherY5,
	}
	vr.Values = append(vr.Values, myval)
	_, err := srv.Spreadsheets.Values.Append(spreadsheetId, insertScheduleRange, &vr).ValueInputOption(inputOption).Do()
	if err != nil {
		log.Fatalf("Unable to store data in sheet. %v", err)
	}
}

func insertEmptyRow() {
	emptyLine := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			&sheets.Request{
				InsertDimension: &sheets.InsertDimensionRequest{
					Range: &sheets.DimensionRange{
						SheetId:    0,
						Dimension:  "ROWS",
						StartIndex: 1,
						EndIndex:   2,
					},
					InheritFromBefore: false,
				},
			},
		},
	}

	_, err := srv.Spreadsheets.BatchUpdate(spreadsheetId, emptyLine).Do()
	if err != nil {
		log.Fatalf("Unable insert empty line. %v", err)
	}
	log.Debug("Empty line inserted!")
}

package main

import (
	"context"
	"fmt"
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

func main() {
	// Get Environment Variables
	credentialsJson := os.Getenv("CREDENTIALS")
	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	// TODO move sheets ranges to some properties file
	engineersRange := "Engineers!A2:S"
	holidaysRange := "Holidays!A2:C"
	scheduleRange := "CI_Watch_Schedule!A2:L"

	ctx := context.Background()
	b, err := ioutil.ReadFile(credentialsJson)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
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
	engineers := make(map[string]engineer.Engineer)
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
	var holidays = []holiday.Holiday{}
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
	scheduleCurrent := make(map[string]schedule.Schedule)
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
	fmt.Print(scheduleCurrent)
}

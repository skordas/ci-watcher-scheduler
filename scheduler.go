package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/skordas/ci-watcher-scheduler/spreadsheets/engineer"
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
	// holidaysRange := "Holidays!A2:C"
	// scheduleRange := "CI_Watch_Schedule!A2:L"

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

	engineersMap := make(map[string]engineer)

	//Some first test if it's working
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, engineersRange).Do()
	if err != nil {
		log.Fatal("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found")
	} else {
		fmt.Println("Here are some Values:")
		for _, row := range resp.Values {
			logging.Debug("%s, %s, %s", row[0], row[1], row[2])
		}
	}
}

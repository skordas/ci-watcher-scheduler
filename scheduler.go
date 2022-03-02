package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/skordas/ci-watcher-scheduler/tools/debug"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	// Get Environment Variables
	credentialsJson := os.Getenv("CREDENTIALS")
	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	readRange := os.Getenv("ENGINEERS_SHEET")
	debug.Log("info", "First info")
	debug.Log("error", "Here some error")
	debug.Log("warning", "some warning")
	debug.Log("debug", "Here is some dubugging")

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

	//Some first test if it's working
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatal("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found")
	} else {
		fmt.Println("Here are some Values:")
		for _, row := range resp.Values {
			fmt.Printf("%s, %s, %s\n", row[0], row[1], row[2])
		}
	}
}

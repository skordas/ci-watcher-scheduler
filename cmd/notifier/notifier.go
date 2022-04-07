package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"os"
	// "time"

	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/schedule"
	"github.com/skordas/ci-watcher-scheduler/tools/logging"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const layoutUS = "1/2/2006"

//calendar properties
var calendarName string = "CI-Watchers Schedule"
var timeZone string = "America/New_York"
var calId string

var currentSchedule = make(map[string]schedule.Schedule)

var credentialsJson string
var srv *calendar.Service
var initiate bool = true

// InitiateSrv is getting all informations needed to start connection
// with carendar
func initiateSrv() {
	if initiate {
		logging.Info("------ Staring calendar client ------")
		credentialsJson = os.Getenv("CREDENTIALS")
		logging.Info("Credentials path: %s", credentialsJson)

		ctx := context.Background()
		credentials, err := ioutil.ReadFile(credentialsJson)
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		config, err := google.JWTConfigFromJSON(credentials, calendar.CalendarScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}

		client := config.Client(ctx)
		srv, err = calendar.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Unable to retrive Sheets client: %v", err)
		}
		initiate = false
	}
}

func main() {
	// dayToNotify := os.Getenv("DATE")
	currentSchedule = spreadsheets.GetCurrentSchedule()

	//setting connetion
	initiateSrv()

	// day, _ := time.Parse(layoutUS, dayToNotify)
	// t := time.Now().Format(time.RFC3339)

	// Checking if calendar exist - if not - create one
	scheduleCalendar, err := srv.CalendarList.List().Do()
	if err != nil {
		log.Fatalf("Unable to get calendar list: %v", err)
	}
	// TODO Move creating calendar outside
	if len(scheduleCalendar.Items) == 0 {
		logging.Info("No calendars. Creating new %s", calendarName)

		schCal := &calendar.Calendar{
			Summary:  calendarName,
			TimeZone: timeZone,
		}
		r, err := srv.Calendars.Insert(schCal).Do()
		if err != nil {
			log.Fatalf("Sorry - can't create a calendar %s, err: %v", calendarName, err)
		}
		calId = r.Id
		logging.Info("New calendar %s created. ID of new calendar: %s", calendarName, calId)
		// adding owners
		// TODO - get managers and add them here

		per := &calendar.AclRule{
			Role: "owner",
			Scope: &calendar.AclRuleScope{
				Type:  "user",
				Value: "skordas@redhat.com",
			},
		}

		srv.Acl.Insert(calId, per)
		// TODO Verification of that.
		logging.Info("Added permission")

	} else {
		for _, c := range scheduleCalendar.Items {
			if c.Summary == calendarName {
				calId = c.Id
				logging.Info("'%s' calendar founded. Calendar ID: %s", calendarName, calId)
				break
			}
		}
	}
	// TODO - maybe  here creaet correct calenda. This can happend when there are some calendars, bot non of what we need.
	if calId == "" {
		logging.Error("Can't find '%s' calendar!", calendarName)
		log.Fatalf("Can't find calendar!")
	}

	// Create event
	// TODO - this is test one

	event := &calendar.Event{
		Summary:     "First Schedule",
		Description: "First test of shedule",
		Start: &calendar.EventDateTime{
			DateTime: "2022-03-23T09:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: "2022-03-23T17:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		Attendees: []*calendar.EventAttendee{
			&calendar.EventAttendee{Email: "skordas@redhat.com"},
		},
	}

	event, err = srv.Events.Insert(calId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event %v\n", err)
	}
	fmt.Printf("event created %s\n", event.HtmlLink)
}

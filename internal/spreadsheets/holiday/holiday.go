package holiday

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const layoutUS = "1/2/2006"

type Holiday struct {
	Country string    // Region where in holiday
	Date    time.Time // Day of holiday
	Name    string    // Name of holiday
}

func New(country interface{}, date interface{}, name interface{}) Holiday {
	h := Holiday{
		country.(string),
		parseDate(date.(string)),
		name.(string),
	}

	log.WithFields(log.Fields{
		"Country": h.Country,
		"Date":    h.Date,
		"Name":    h.Name,
	}).Debug("Getting Holiday from spreadsheet")

	return h
}

func parseDate(dateAsString string) time.Time {
	day, err := time.Parse(layoutUS, dateAsString)
	if err != nil {
		log.WithFields(log.Fields{"Layout": layoutUS, "Date To Parse": dateAsString}).Fatal("Can't parse date!")
	}
	return day
}

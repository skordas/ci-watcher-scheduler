package holiday

import (
	log "github.com/sirupsen/logrus"
)

type Holiday struct {
	Country string // Region where in holiday
	Date    string // Day of holiday
	Name    string // Name of holiday
}

func New(country interface{}, date interface{}, name interface{}) Holiday {
	h := Holiday{
		country.(string),
		date.(string),
		name.(string),
	}

	log.WithFields(log.Fields{
		"Country": h.Country,
		"Date":    h.Date,
		"Name":    h.Name,
	}).Debug("Getting Holiday from spreadsheet")

	return h
}

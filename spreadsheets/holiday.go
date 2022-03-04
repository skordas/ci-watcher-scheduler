package spreadsheets

import "time"

type holiday struct {
	Country string    // Region where in holiday
	Date time.Date()  // Day of holiday
	Name string       // Name of holiday
}
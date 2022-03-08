package holiday

//import "time"
import "github.com/skordas/ci-watcher-scheduler/tools/logging"

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

	logging.Debug("------ Getting holiday from spreadsheet ------")
	logging.Debug("Country: %s", h.Country)
	logging.Debug("Date: %s", h.Date)
	logging.Debug("Name: %s", h.Name)

	return h
}

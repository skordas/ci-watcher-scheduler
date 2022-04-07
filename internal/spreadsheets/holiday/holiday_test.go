package holiday

import (
	"testing"
)

func TestNewHolidayCreateCorrectly(t *testing.T) {
	h := New("US", "01/01/2022", "Holiday")

	if h.Country != "US" {
		t.Errorf("Holiday Country should be: %s, but it's: %v", "US", h.Country)
	}
	if h.Date != "01/01/2022" {
		t.Errorf("Holiday Date should be: %s, but it's %v", "01/01/2022", h.Date)
	}
	if h.Name != "Holiday" {
		t.Errorf("Holida Name should be: %s, but it's: %v", "Holiday", h.Name)
	}
}

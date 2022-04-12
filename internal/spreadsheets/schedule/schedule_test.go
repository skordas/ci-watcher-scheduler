package schedule

import (
	"testing"
	"time"
)

func TestNewCreateCorrectSchedule(t *testing.T) {
	sch := New(
		"01/01/2022",
		"Men",
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
	)

	if sch.Date != time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Schedule Date should be: %s, but it's: %v", "01/01/2022", sch.Date)
	}
	if sch.Manager != "Men" {
		t.Errorf("Schedule Manager should be: %s, but it's: %v", "Men", sch.Manager)
	}
	if sch.E2eWatcherY0 != "A" {
		t.Errorf("Schedule E2eWatcherY0 should be: %s, but it's: %v", "A", sch.E2eWatcherY0)
	}
	if sch.E2eWatcherY1 != "B" {
		t.Errorf("Schedule E2eWatcherY1 should be: %s, but it's: %v", "B", sch.E2eWatcherY1)
	}
	if sch.E2eWatcherY2 != "C" {
		t.Errorf("Schedule E2eWatcherY2 should be: %s, but it's: %v", "C", sch.E2eWatcherY2)
	}
	if sch.E2eWatcherY3 != "D" {
		t.Errorf("Schedule E2eWatcherY3 should be: %s, but it's: %v", "D", sch.E2eWatcherY3)
	}
	if sch.E2eWatcherY4 != "E" {
		t.Errorf("Schedule E2eWatcherY4 should be: %s, but it's: %v", "E", sch.E2eWatcherY4)
	}
	if sch.E2eWatcherY5 != "F" {
		t.Errorf("Schedule E2eWatcherY5 should be: %s, but it's: %v", "F", sch.E2eWatcherY5)
	}
	if sch.UpgrWatcherY0 != "G" {
		t.Errorf("Schedule UpgrWatcherY0 should be: %s, but it's: %v", "G", sch.UpgrWatcherY0)
	}
	if sch.UpgrWatcherY1 != "H" {
		t.Errorf("Schedule UpgrWatcherY1 should be: %s, but it's: %v", "H", sch.UpgrWatcherY1)
	}
	if sch.UpgrWatcherY2 != "I" {
		t.Errorf("Schedule UpgrWatcherY2 should be: %s, but it's: %v", "I", sch.UpgrWatcherY2)
	}
	if sch.UpgrWatcherY3 != "J" {
		t.Errorf("Schedule UpgrWatcherY3 should be: %s, but it's: %v", "J", sch.UpgrWatcherY3)
	}
	if sch.UpgrWatcherY4 != "K" {
		t.Errorf("Schedule UpgrWatcherY4 should be: %s, but it's: %v", "K", sch.UpgrWatcherY4)
	}
	if sch.UpgrWatcherY5 != "L" {
		t.Errorf("Schedule UpgrWatcherY5 should be: %s, but it's: %v", "L", sch.UpgrWatcherY5)
	}
}

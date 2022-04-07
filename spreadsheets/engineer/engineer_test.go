package engineer

import (
	"testing"
)

func TestNewEngineerCreatedCoretly(t *testing.T) {
	e := New(
		"A",
		"B",
		"C",
		"2",
		"true",
		"true",
		"true",
		"true",
		"D",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
	)

	if e.Kerberos != "A" {
		t.Errorf("New Engineer Kerberos should be: %s, but it's: %v", "A", e.Kerberos)
	}
	if e.Team != "B" {
		t.Errorf("New Engineer Team should be: %s, but it's: %v", "B", e.Team)
	}
	if e.Country != "C" {
		t.Errorf("New Engineer Country should be: %s, but it's: %v", "B", e.Country)
	}
	if e.Week != 2 {
		t.Errorf("New Engineer week should be: %d, but it's: %v", 2, e.Country)
	}
	if !e.Manager {
		t.Errorf("New Engineer Manager should be %t, but it's: %v", true, e.Manager)
	}
	if !e.ReleaseLead {
		t.Errorf("New Engineer Release Lead should be %t, but it's: %v", true, e.ReleaseLead)
	}
	if !e.TeamLead {
		t.Errorf("New Engineer TeamLead should be %t, but it's: %v", true, e.TeamLead)
	}
	if !e.NewToCi {
		t.Errorf("New Engineer NewToCi should be %t, but it's: %v", true, e.NewToCi)
	}
	if e.NewToCiBuddy != "D" {
		t.Errorf("New Engineer Buddy should be: %s, but it's: %v", "D", e.NewToCiBuddy)
	}
	if !e.E2eY0 {
		t.Errorf("New Engineer E2eY0 should be %t, but it's: %v", true, e.E2eY0)
	}
	if !e.E2eY1 {
		t.Errorf("New Engineer E2eY1 should be %t, but it's: %v", true, e.E2eY1)
	}
	if !e.E2eY2 {
		t.Errorf("New Engineer E2eY2 should be %t, but it's: %v", true, e.E2eY2)
	}
	if !e.E2eY3 {
		t.Errorf("New Engineer E2eY3 should be %t, but it's: %v", true, e.E2eY3)
	}
	if !e.E2eY4 {
		t.Errorf("New Engineer E2eY4 should be %t, but it's: %v", true, e.E2eY4)
	}
	if !e.E2eY5 {
		t.Errorf("New Engineer E2eY5 should be %t, but it's: %v", true, e.E2eY5)
	}
	if !e.UpgrY0 {
		t.Errorf("New Engineer UpgrY0 should be %t, but it's: %v", true, e.UpgrY0)
	}
	if !e.UpgrY1 {
		t.Errorf("New Engineer UpgrY1 should be %t, but it's: %v", true, e.UpgrY1)
	}
	if !e.UpgrY2 {
		t.Errorf("New Engineer UpgrY2 should be %t, but it's: %v", true, e.UpgrY2)
	}
	if !e.UpgrY3 {
		t.Errorf("New Engineer UpgrY3 should be %t, but it's: %v", true, e.UpgrY3)
	}
	if !e.UpgrY4 {
		t.Errorf("New Engineer UpgrY4 should be %t, but it's: %v", true, e.UpgrY4)
	}
	if !e.UpgrY5 {
		t.Errorf("New Engineer UpgrY5 should be %t, but it's: %v", true, e.UpgrY5)
	}
	if e.Activity != 0 {
		t.Errorf("New Engineer Activity should be %d, but it's: %v", 0, e.Activity)
	}
}

func TestNewWithEmptyFields(t *testing.T) {
	e := New("A",
		"B",
		"C",
		"",
		"",
		"true",
		"true",
		"true",
		"D",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
		"true",
	)

	if e.Week != 0 {
		t.Errorf("New Engineer Week should be: %d, but it's: %v", 0, e.Week)
	}
	if e.Manager {
		t.Errorf("New Engineer Manager should be: %t, but it's: %v", false, e.Manager)
	}
}

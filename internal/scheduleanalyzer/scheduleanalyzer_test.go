package scheduleanalyzer

import (
	"fmt"
	"testing"
	"time"

	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/engineer"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/holiday"
	"github.com/skordas/ci-watcher-scheduler/internal/spreadsheets/schedule"
)

func TestAddActivity(t *testing.T) {
	t.Log("------ AddActivity test ------")

	//Given
	eng := loadEngineers()

	// When
	AddActivity(eng, "A")
	result := eng["A"].Activity
	correct := 1

	// Then
	if result != correct {
		t.Errorf("AddActivity should return %d. Instead got %d", correct, result)
	}

}

func TestCountEngineersActivityWithLoadedSchedule(t *testing.T) {
	t.Log("------ CountEngineersActivity with loaded schedule ------")

	// Given
	eng := loadEngineers()
	sch := loadSchedule()
	ea := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}

	// When
	CountEngineersActivity(eng, sch)
	correct := 6

	// Then
	for _, e := range ea {
		t.Run(fmt.Sprintf("Check activity for engineer: %s", e), func(t *testing.T) {
			result := eng[e].Activity
			if result != correct {
				t.Errorf("Expected Activity for engineer: %s, is %d, but got: %d", e, correct, result)
			}
		})
	}
}

func TestCountEngineersActivityWithEmptySchedule(t *testing.T) {
	t.Log("------ CountEngineersActivity with empty schedule ------")

	// Given
	eng := loadEngineers()
	sch := make(map[time.Time]schedule.Schedule)
	ea := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}

	// When
	CountEngineersActivity(eng, sch)
	correct := 0

	// Then
	for _, e := range ea {
		result := eng[e].Activity
		if result != correct {
			t.Errorf("Expected Activity for engineer: %s, is %d, but got: %d", e, correct, result)
		}
	}
}

func TestGetWatcherForRemovingManagersFromList(t *testing.T) {
	t.Log("------ GetWatcherFor will remove all managers from list ------")

	// Given
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.Manager = true
	eng["B"] = e
	eng["C"] = e
	eng["D"] = e
	eng["E"] = e
	eng["F"] = e
	eng["G"] = e
	eng["H"] = e
	eng["I"] = e
	eng["J"] = e
	eng["K"] = e
	eng["L"] = e

	// When
	result := GetWatcherFor("abc", eng, date, hol)
	correct := "A"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as NO manager. But got %s", correct, result)
	}
}

func TestGetWatcherForRemovingNewToCiFromList(t *testing.T) {
	t.Log("------ GetWatcherFor will remove all New To CI engineers from list ------")

	// Given
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.NewToCi = true
	eng["A"] = e
	eng["C"] = e
	eng["D"] = e
	eng["E"] = e
	eng["F"] = e
	eng["G"] = e
	eng["H"] = e
	eng["I"] = e
	eng["J"] = e
	eng["K"] = e
	eng["L"] = e

	// When
	result := GetWatcherFor("abc", eng, date, hol)
	correct := "B"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as NO new to CI. But got %s", correct, result)
	}
}

func TestGetWatcherForRemovingEngineerForOddWeek(t *testing.T) {
	t.Log("--- GetWatcherFor will remove all engineers with assigned odd week")

	// Given
	eng := loadEngineers()
	date := time.Date(2022, time.February, 9, 0, 0, 0, 0, time.UTC) // even week (6th)
	hol := loadHolidays()
	e := eng["A"]
	e.Week = 2
	eng["C"] = e

	// When
	result := GetWatcherFor("abc", eng, date, hol)
	correct := "C"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as only for even week. But got %s", correct, result)
	}
}

func TestGetWatcherForRemovingEngineerForEvenWeek(t *testing.T) {
	t.Log("--- GetWatcherFor will remove all engineers with assigned even week")

	// Given
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC) // odd week (5th)
	hol := loadHolidays()
	e := eng["A"]
	e.Week = 2
	eng["A"] = e
	eng["B"] = e
	eng["C"] = e
	eng["E"] = e
	eng["F"] = e
	eng["G"] = e
	eng["H"] = e
	eng["I"] = e
	eng["J"] = e
	eng["K"] = e
	eng["L"] = e

	// When
	result := GetWatcherFor("abc", eng, date, hol)
	correct := "D"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as only for even week. But got %s", correct, result)
	}
}

func TestGetWatcherForRemovingEngineerWithHolidays(t *testing.T) {
	t.Log("--- GetWatcherFor will remove all engineers celebrating holidays")

	// Given
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC) // odd week (5th)
	hol := loadHolidays()
	e := eng["A"]
	e.Country = "UK"
	eng["A"] = e
	eng["B"] = e
	eng["C"] = e
	eng["D"] = e
	eng["F"] = e
	eng["G"] = e
	eng["H"] = e
	eng["I"] = e
	eng["J"] = e
	eng["K"] = e
	eng["L"] = e

	// When
	result := GetWatcherFor("abc", eng, date, hol)
	correct := "E"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as only not celebrating holidays. But got %s", correct, result)
	}
}

func TestGetWatcherForE2ey0WillBeSelected(t *testing.T) {
	t.Log("------ Getting e2ey0 watcher ------")

	// Given
	watchType := "e2ey0"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.E2eY0 = true
	eng["F"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "F"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForE2ey1WillBeSelected(t *testing.T) {
	t.Log("------ Getting e2ey1 watcher ------")

	// Given
	watchType := "e2ey1"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.E2eY1 = true
	eng["G"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "G"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForE2ey2WillBeSelected(t *testing.T) {
	t.Log("------ Getting e2ey2 watcher ------")

	// Given
	watchType := "e2ey2"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.E2eY2 = true
	eng["I"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "I"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForE2ey3WillBeSelected(t *testing.T) {
	t.Log("------ Getting e2ey3 watcher ------")

	// Given
	watchType := "e2ey3"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.E2eY3 = true
	eng["J"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "J"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForE2ey4WillBeSelected(t *testing.T) {
	t.Log("------ Getting e2ey4 watcher ------")

	// Given
	watchType := "e2ey4"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.E2eY4 = true
	eng["K"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "K"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForE2ey5WillBeSelected(t *testing.T) {
	t.Log("------ Getting e2ey5 watcher ------")

	// Given
	watchType := "e2ey5"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.E2eY5 = true
	eng["L"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "L"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForUpgry0WillBeSelected(t *testing.T) {
	t.Log("------ Getting upgry0 watcher ------")

	// Given
	watchType := "upgry0"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.UpgrY0 = true
	eng["A"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "A"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForUpgry1WillBeSelected(t *testing.T) {
	t.Log("------ Getting upgry1 watcher ------")

	// Given
	watchType := "upgry1"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.UpgrY1 = true
	eng["B"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "B"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForUpgry2WillBeSelected(t *testing.T) {
	t.Log("------ Getting upgry1 watcher ------")

	// Given
	watchType := "upgry2"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.UpgrY2 = true
	eng["C"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "C"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForUpgry3WillBeSelected(t *testing.T) {
	t.Log("------ Getting upgry1 watcher ------")

	// Given
	watchType := "upgry3"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.UpgrY3 = true
	eng["D"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "D"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForUpgry4WillBeSelected(t *testing.T) {
	t.Log("------ Getting upgry4 watcher ------")

	// Given
	watchType := "upgry4"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.UpgrY4 = true
	eng["E"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "E"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForUpgry5WillBeSelected(t *testing.T) {
	t.Log("------ Getting upgry5 watcher ------")

	// Given
	watchType := "upgry5"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.UpgrY5 = true
	eng["F"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "F"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s as %s watcher. But got %s", correct, watchType, result)
	}
}

func TestGetWatcherForOneWithPickEngineerWithTheLowestActivity(t *testing.T) {
	t.Log("------ Getting watcher with the lowest activity")

	// Given
	watchType := "e2ey0"
	eng := loadEngineers()
	date := time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC)
	hol := loadHolidays()
	e := eng["A"]
	e.E2eY0 = true
	e.Activity = 14
	eng["A"] = e
	e.Activity = 14
	eng["B"] = e
	e.Activity = 13
	eng["C"] = e
	e.Activity = 12
	eng["D"] = e
	e.Activity = 11
	eng["E"] = e
	e.Activity = 10
	eng["F"] = e
	e.Activity = 9
	eng["G"] = e
	e.Activity = 8
	eng["H"] = e
	e.Activity = 7
	eng["I"] = e
	e.Activity = 6
	eng["J"] = e
	e.Activity = 5
	eng["K"] = e
	e.Activity = 4
	eng["L"] = e

	// When
	result := GetWatcherFor(watchType, eng, date, hol)
	correct := "L"

	// Then
	if result != correct {
		t.Errorf("Expected engineer %s with the lowest activity, but got %s", correct, result)
	}
}

func TestGetDayToSchedule(t *testing.T) {
	// Given
	sch := loadSchedule()

	// When
	result := GetDayToSchedule(sch)
	correct := time.Date(2022, time.January, 7, 0, 0, 0, 0, time.UTC)

	// Then
	if result != correct {
		t.Errorf("Expected date %v, but got %v", correct, result)
	}
}

// loadEngineers will return loaded map with sample engineers - later used
// for tests
func loadEngineers() map[string]engineer.Engineer {
	var em = make(map[string]engineer.Engineer)
	e := engineer.Engineer{
		"A",        // Kerberos
		"TeamName", // Team
		"US",       // Country
		1,          // Week
		false,      // Manager
		false,      // ReleaseLead
		false,      // TeamLead
		false,      // NewToCi
		"buddy",    // NewToCiBuddy
		false,      // E2eY0
		false,      // E2eY1
		false,      // E2eY2
		false,      // E2eY3
		false,      // E2eY4
		false,      // E2eY5
		false,      // UpgrY0
		false,      // UpgrY1
		false,      // UpgrY2
		false,      // UpgrY3
		false,      // UpgrY4
		false,      // UpgrY5
		0,          // Activity
	}
	em["A"] = e

	e.Kerberos = "B"
	em["B"] = e

	e.Kerberos = "C"
	em["C"] = e

	e.Kerberos = "D"
	em["D"] = e

	e.Kerberos = "E"
	em["E"] = e

	e.Kerberos = "F"
	em["F"] = e

	e.Kerberos = "G"
	em["G"] = e

	e.Kerberos = "H"
	em["H"] = e

	e.Kerberos = "I"
	em["I"] = e

	e.Kerberos = "J"
	em["J"] = e

	e.Kerberos = "K"
	em["K"] = e

	e.Kerberos = "L"
	em["L"] = e

	return em
}

// loadSchedule will return map with sample schedule - loaded with names of
// engineers used in loadEngineers func.
func loadSchedule() map[time.Time]schedule.Schedule {
	var sm = make(map[time.Time]schedule.Schedule)
	s := schedule.Schedule{
		time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC), // Date
		"",  // Manager
		"A", // E2eWatcherY0
		"B", // E2eWatcherY1
		"C", // E2eWatcherY2
		"D", // E2eWatcherY3
		"E", // E2eWatcherY4
		"F", // E2eWatcherY5
		"G", // UpgrWatcherY0
		"H", // UpgrWatcherY1
		"I", // UpgrWatcherY2
		"J", // UpgrWatcherY3
		"K", // UpgrWatcherY4
		"L", // UpgrWatcherY5
	}
	sm[time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)] = s

	s.Date = time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC)
	sm[time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC)] = s

	s.Date = time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC)
	sm[time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC)] = s

	s.Date = time.Date(2022, time.January, 4, 0, 0, 0, 0, time.UTC)
	sm[time.Date(2022, time.January, 4, 0, 0, 0, 0, time.UTC)] = s

	s.Date = time.Date(2022, time.January, 5, 0, 0, 0, 0, time.UTC)
	sm[time.Date(2022, time.January, 5, 0, 0, 0, 0, time.UTC)] = s

	s.Date = time.Date(2022, time.January, 6, 0, 0, 0, 0, time.UTC)
	sm[time.Date(2022, time.January, 7, 0, 0, 0, 0, time.UTC)] = s

	return sm
}

// loadHolidays will return array with test example
func loadHolidays() []holiday.Holiday {
	var ha = []holiday.Holiday{}
	h := holiday.Holiday{
		"UK", //Country
		time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC), // Date
		"Holiday", // Name
	}
	ha = append(ha, h)
	return ha
}

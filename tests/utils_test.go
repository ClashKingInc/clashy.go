package clashy_test

import (
	"testing"
	"time"

	clashy "github.com/clashkinginc/clashy.go"
)

func TestGetSeasonBeforeTransition(t *testing.T) {
	season := clashy.GetSeason(time.Date(2025, time.August, 24, 12, 0, 0, 0, time.UTC), true)

	if season.SeasonID != "2025-08" {
		t.Fatalf("unexpected season id: %s", season.SeasonID)
	}
	if !season.StartTime.Equal(time.Date(2025, time.July, 28, 5, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected season start: %s", season.StartTime)
	}
	if !season.EndTime.Equal(time.Date(2025, time.August, 25, 5, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected season end: %s", season.EndTime)
	}
}

func TestGetSeasonFixedSeptemberWindow(t *testing.T) {
	season := clashy.GetSeason(time.Date(2025, time.September, 15, 12, 0, 0, 0, time.UTC), true)

	if season.SeasonID != "2025-09" {
		t.Fatalf("unexpected season id: %s", season.SeasonID)
	}
	if !season.StartTime.Equal(time.Date(2025, time.August, 25, 5, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected season start: %s", season.StartTime)
	}
	if !season.EndTime.Equal(time.Date(2025, time.October, 6, 5, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected season end: %s", season.EndTime)
	}
}

func TestGetSeasonAfterTransition(t *testing.T) {
	season := clashy.GetSeason(time.Date(2025, time.November, 10, 12, 0, 0, 0, time.UTC), true)

	if season.SeasonID != "2025-11" {
		t.Fatalf("unexpected season id: %s", season.SeasonID)
	}
	if !season.StartTime.Equal(time.Date(2025, time.November, 3, 5, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected season start: %s", season.StartTime)
	}
	if !season.EndTime.Equal(time.Date(2025, time.December, 1, 5, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected season end: %s", season.EndTime)
	}
}

func TestGetSeasonByID(t *testing.T) {
	tests := []struct {
		id    string
		start time.Time
		end   time.Time
	}{
		{
			id:    "2025-08",
			start: time.Date(2025, time.July, 28, 5, 0, 0, 0, time.UTC),
			end:   time.Date(2025, time.August, 25, 5, 0, 0, 0, time.UTC),
		},
		{
			id:    "2025-09",
			start: time.Date(2025, time.August, 25, 5, 0, 0, 0, time.UTC),
			end:   time.Date(2025, time.October, 6, 5, 0, 0, 0, time.UTC),
		},
		{
			id:    "2025-10",
			start: time.Date(2025, time.October, 6, 5, 0, 0, 0, time.UTC),
			end:   time.Date(2025, time.November, 3, 5, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		season, err := clashy.GetSeasonByID(tt.id)
		if err != nil {
			t.Fatalf("get season by id %s: %v", tt.id, err)
		}
		if !season.StartTime.Equal(tt.start) {
			t.Fatalf("%s unexpected start: %s", tt.id, season.StartTime)
		}
		if !season.EndTime.Equal(tt.end) {
			t.Fatalf("%s unexpected end: %s", tt.id, season.EndTime)
		}
	}
}

func TestGetTournamentWindow(t *testing.T) {
	window := clashy.GetTournamentWindow(time.Date(2026, time.March, 31, 12, 0, 0, 0, time.UTC))

	if window.ID != "2026-03-30" {
		t.Fatalf("unexpected window id: %s", window.ID)
	}
	if !window.StartTime.Equal(time.Date(2026, time.March, 30, 5, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected start: %s", window.StartTime)
	}
	if !window.EndTime.Equal(time.Date(2026, time.April, 6, 5, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected end: %s", window.EndTime)
	}
}

func TestGetTournamentWindowBeforeMondayReset(t *testing.T) {
	window := clashy.GetTournamentWindow(time.Date(2026, time.March, 30, 4, 59, 0, 0, time.UTC))

	if window.ID != "2026-03-23" {
		t.Fatalf("unexpected window id: %s", window.ID)
	}
}

func TestGetClanGamesWindow(t *testing.T) {
	start := clashy.GetClanGamesStart(time.Date(2026, time.March, 29, 12, 0, 0, 0, time.UTC))
	end := clashy.GetClanGamesEnd(time.Date(2026, time.March, 29, 12, 0, 0, 0, time.UTC))

	if !start.Equal(time.Date(2026, time.April, 22, 8, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected clan games start: %s", start)
	}
	if !end.Equal(time.Date(2026, time.April, 28, 8, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected clan games end: %s", end)
	}
}

func TestGetRaidWeekendWindow(t *testing.T) {
	start := clashy.GetRaidWeekendStart(time.Date(2026, time.March, 31, 12, 0, 0, 0, time.UTC))
	end := clashy.GetRaidWeekendEnd(time.Date(2026, time.March, 31, 12, 0, 0, 0, time.UTC))

	if !start.Equal(time.Date(2026, time.April, 3, 7, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected raid weekend start: %s", start)
	}
	if !end.Equal(time.Date(2026, time.April, 6, 7, 0, 0, 0, time.UTC)) {
		t.Fatalf("unexpected raid weekend end: %s", end)
	}
}

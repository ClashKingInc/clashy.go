package clashy

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var invalidTagRegex = regexp.MustCompile(`[^A-Z0-9#]`)

var (
	seasonCutoffStart = time.Date(2025, time.August, 25, 5, 0, 0, 0, time.UTC)
	seasonCutoffEnd   = time.Date(2025, time.October, 6, 5, 0, 0, 0, time.UTC)
	seasonDuration    = 28 * 24 * time.Hour
)

type SeasonWindow struct {
	SeasonID  string
	StartTime time.Time
	EndTime   time.Time
}

type TournamentWindow struct {
	ID        string
	StartTime time.Time
	EndTime   time.Time
}

func CorrectTag(tag string) string {
	tag = strings.ToUpper(strings.TrimSpace(tag))
	tag = strings.ReplaceAll(tag, "O", "0")
	tag = invalidTagRegex.ReplaceAllString(tag, "")
	if tag == "" {
		return tag
	}
	if !strings.HasPrefix(tag, "#") {
		tag = "#" + tag
	}
	return tag
}

func encodeTag(tag string) string {
	return url.PathEscape(CorrectTag(tag))
}

func cacheExpiry(header string) time.Duration {
	for _, part := range strings.Split(header, ",") {
		part = strings.TrimSpace(part)
		part = strings.TrimPrefix(part, "public ")
		if strings.HasPrefix(part, "max-age=") {
			if secs, err := time.ParseDuration(strings.TrimPrefix(part, "max-age=") + "s"); err == nil {
				return secs
			}
		}
	}
	return 0
}

func jwtIP(token string) string {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return ""
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return ""
	}
	var decoded struct {
		Limits []struct {
			Cidrs []string `json:"cidrs"`
		} `json:"limits"`
	}
	if json.Unmarshal(payload, &decoded) != nil || len(decoded.Limits) < 2 || len(decoded.Limits[1].Cidrs) == 0 {
		return ""
	}
	return strings.Split(decoded.Limits[1].Cidrs[0], "/")[0]
}

func contextOrBackground(ctx context.Context) context.Context {
	if ctx != nil {
		return ctx
	}
	return context.Background()
}

func normalizeAPIBase(raw string) string {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimSuffix(raw, "/")
	return raw
}

func encodeRealtime(path string, realtime bool) string {
	if !realtime {
		return path
	}
	sep := "?"
	if strings.Contains(path, "?") {
		sep = "&"
	}
	return path + sep + url.Values{"realtime": []string{"true"}}.Encode()
}

func GetSeasonID() string {
	return GetSeason(time.Time{}, true).SeasonID
}

func GetSeason(timestamp time.Time, forward bool) SeasonWindow {
	target := utcNowIfZero(timestamp)

	if !target.After(seasonCutoffStart) {
		endTime := oldSeasonEnd(target, forward)
		startTime := oldSeasonStartFromEnd(endTime)
		return SeasonWindow{
			SeasonID:  endTime.Format("2006-01"),
			StartTime: startTime,
			EndTime:   endTime,
		}
	}

	if !target.After(seasonCutoffEnd) {
		return SeasonWindow{
			SeasonID:  "2025-09",
			StartTime: seasonCutoffStart,
			EndTime:   seasonCutoffEnd,
		}
	}

	timeDifference := target.Sub(seasonCutoffEnd)
	seasonsPassed := int(timeDifference / seasonDuration)

	startTime := seasonCutoffEnd.Add(time.Duration(seasonsPassed) * seasonDuration)
	endTime := startTime.Add(seasonDuration)

	refYear, refMonthIndex := seasonCutoffEnd.UTC().Year(), int(seasonCutoffEnd.UTC().Month())-1
	totalMonths := refYear*12 + refMonthIndex + seasonsPassed
	year := totalMonths / 12
	month := totalMonths - year*12 + 1

	return SeasonWindow{
		SeasonID:  fmt.Sprintf("%04d-%02d", year, month),
		StartTime: startTime,
		EndTime:   endTime,
	}
}

func GetSeasonByID(seasonID string) (SeasonWindow, error) {
	if seasonID == "2025-09" {
		return SeasonWindow{
			SeasonID:  seasonID,
			StartTime: seasonCutoffStart,
			EndTime:   seasonCutoffEnd,
		}, nil
	}

	year, month, err := parseSeasonID(seasonID)
	if err != nil {
		return SeasonWindow{}, err
	}

	refYear, refMonthIndex := seasonCutoffEnd.UTC().Year(), int(seasonCutoffEnd.UTC().Month())-1
	totalMonthsTarget := year*12 + (month - 1)
	totalMonthsRef := refYear*12 + refMonthIndex
	seasonsPassed := totalMonthsTarget - totalMonthsRef

	if seasonsPassed < 0 {
		endTime := lastMondayAtFiveUTC(year, time.Month(month))
		return SeasonWindow{
			SeasonID:  seasonID,
			StartTime: oldSeasonStartFromEnd(endTime),
			EndTime:   endTime,
		}, nil
	}

	startTime := seasonCutoffEnd.Add(time.Duration(seasonsPassed) * seasonDuration)
	return SeasonWindow{
		SeasonID:  seasonID,
		StartTime: startTime,
		EndTime:   startTime.Add(seasonDuration),
	}, nil
}

func GetTournamentWindow(timestamp time.Time) TournamentWindow {
	now := utcNowIfZero(timestamp)

	day := int(now.Weekday())
	hour := now.Hour()
	daysSinceMonday := (day + 6) % 7
	if day == int(time.Monday) && hour < 5 {
		daysSinceMonday = 7
	}

	lastMonday := time.Date(now.Year(), now.Month(), now.Day()-daysSinceMonday, 5, 0, 0, 0, time.UTC)
	nextMonday := lastMonday.Add(7 * 24 * time.Hour)

	return TournamentWindow{
		ID:        lastMonday.Format("2006-01-02"),
		StartTime: lastMonday,
		EndTime:   nextMonday,
	}
}

func GetTournamentWindowByID(id string) (TournamentWindow, error) {
	timestamp, err := time.Parse(time.RFC3339, id+"T05:00:00.000Z")
	if err != nil {
		return TournamentWindow{}, err
	}
	return GetTournamentWindow(timestamp), nil
}

func GetClanGamesStart(timestamp time.Time) time.Time {
	t := utcNowIfZero(timestamp)
	month := t.Month()
	year := t.Year()
	thisMonthEnd := time.Date(year, month, 28, 8, 0, 0, 0, time.UTC)
	if t.After(thisMonthEnd) {
		if month == time.December {
			month = time.January
			year++
		} else {
			month++
		}
	}
	return time.Date(year, month, 22, 8, 0, 0, 0, time.UTC)
}

func GetClanGamesEnd(timestamp time.Time) time.Time {
	t := utcNowIfZero(timestamp)
	month := t.Month()
	year := t.Year()
	thisMonthEnd := time.Date(year, month, 28, 8, 0, 0, 0, time.UTC)
	if t.After(thisMonthEnd) {
		if month == time.December {
			month = time.January
			year++
		} else {
			month++
		}
	}
	return time.Date(year, month, 28, 8, 0, 0, 0, time.UTC)
}

func GetRaidWeekendStart(timestamp time.Time) time.Time {
	return GetRaidWeekendEnd(timestamp).Add(-72 * time.Hour)
}

func GetRaidWeekendEnd(timestamp time.Time) time.Time {
	t := utcNowIfZero(timestamp).Add(-7*time.Hour - time.Microsecond)
	daysUntilNextMonday := 7 - int(t.Weekday()-time.Monday)
	if t.Weekday() == time.Sunday {
		daysUntilNextMonday = 1
	}
	if t.Weekday() == time.Monday {
		daysUntilNextMonday = 7
	}
	nextMonday := t.AddDate(0, 0, daysUntilNextMonday)
	return time.Date(nextMonday.Year(), nextMonday.Month(), nextMonday.Day(), 7, 0, 0, 0, time.UTC)
}

func utcNowIfZero(timestamp time.Time) time.Time {
	if timestamp.IsZero() {
		return time.Now().UTC()
	}
	return timestamp.UTC()
}

func parseSeasonID(seasonID string) (int, int, error) {
	parts := strings.Split(seasonID, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid season id %q", seasonID)
	}
	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid season id %q", seasonID)
	}
	month, err := strconv.Atoi(parts[1])
	if err != nil || month < 1 || month > 12 {
		return 0, 0, fmt.Errorf("invalid season id %q", seasonID)
	}
	return year, month, nil
}

func oldSeasonEnd(timestamp time.Time, forward bool) time.Time {
	end := lastMondayAtFiveUTC(timestamp.Year(), timestamp.Month())
	if forward && !timestamp.Before(end) {
		year, month := nextMonth(timestamp.Year(), timestamp.Month())
		end = lastMondayAtFiveUTC(year, month)
	}
	return end
}

func oldSeasonStartFromEnd(end time.Time) time.Time {
	year, month := previousMonth(end.Year(), end.Month())
	return lastMondayAtFiveUTC(year, month)
}

func lastMondayAtFiveUTC(year int, month time.Month) time.Time {
	lastDay := time.Date(year, month+1, 0, 5, 0, 0, 0, time.UTC)
	offset := (int(lastDay.Weekday()) + 6) % 7
	return lastDay.AddDate(0, 0, -offset)
}

func nextMonth(year int, month time.Month) (int, time.Month) {
	if month == time.December {
		return year + 1, time.January
	}
	return year, month + 1
}

func previousMonth(year int, month time.Month) (int, time.Month) {
	if month == time.January {
		return year - 1, time.December
	}
	return year, month - 1
}

package clashy

import (
	"encoding/json"
	"time"
)

// Timestamp stores both the raw Clash API timestamp string and its parsed time.
type Timestamp struct {
	// RawTime is the original API timestamp.
	//
	//	20060102T150405.000Z
	RawTime string
	// Time is the parsed UTC time.
	Time time.Time
}

// UnmarshalJSON parses Clash API timestamp strings into Timestamp values.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	ts, err := FromTimestamp(raw)
	if err != nil {
		return err
	}
	t.RawTime = raw
	t.Time = ts
	return nil
}

// SecondsUntil returns the number of whole seconds from now until the timestamp.
func (t Timestamp) SecondsUntil() int {
	if t.Time.IsZero() {
		return 0
	}
	return int(time.Until(t.Time).Seconds())
}

// Before reports whether this timestamp occurs before another timestamp.
func (t Timestamp) Before(other Timestamp) bool { return t.Time.Before(other.Time) }

// After reports whether this timestamp occurs after another timestamp.
func (t Timestamp) After(other Timestamp) bool { return t.Time.After(other.Time) }

// TimeDelta represents an elapsed duration.
type TimeDelta struct {
	// Duration is the underlying Go duration.
	time.Duration
}

// Badge contains the common small, medium, and large image URLs for clan badges.
type Badge struct {
	// Small is the small badge image URL.
	Small string `json:"small"`
	// Medium is the medium badge image URL.
	Medium string `json:"medium"`
	// Large is the large badge image URL.
	Large string `json:"large"`
}

// URL returns the preferred badge URL, choosing medium, then large, then small.
func (b Badge) URL() string {
	if b.Medium != "" {
		return b.Medium
	}
	if b.Large != "" {
		return b.Large
	}
	return b.Small
}

// Icon contains small icon URLs returned for leagues and labels.
type Icon struct {
	// Small is the small icon URL.
	Small string `json:"small"`
	// Medium is the medium icon URL.
	Medium string `json:"medium"`
	// Tiny is the tiny icon URL.
	Tiny string `json:"tiny"`
}

// Achievement describes one player achievement and its current progress.
type Achievement struct {
	// Name is the achievement display name.
	Name string `json:"name"`
	// Stars is the number of achievement stars earned.
	Stars int `json:"stars"`
	// Value is the current progress value.
	Value int `json:"value"`
	// Target is the value needed to complete the achievement.
	Target int `json:"target"`
	// Info describes the achievement goal.
	Info string `json:"info"`
	// CompletionInfo describes the completed achievement state.
	CompletionInfo string `json:"completionInfo"`
	// Village identifies the village or game area for the achievement.
	Village string `json:"village"`
}

// Location is a country or global location used by ranking endpoints.
type Location struct {
	// ID is the numeric location identifier.
	ID int `json:"id"`
	// Name is the English location name.
	Name string `json:"name"`
	// IsCountry reports whether the location is a country.
	IsCountry bool `json:"isCountry"`
	// CountryCode is the ISO-style country code when the location is a country.
	CountryCode string `json:"countryCode"`
	// Localised is the API-provided localized display name.
	Localised string `json:"localizedName"`
	responseMeta
}

// League is a league, war league, builder-base league, or capital league.
type League struct {
	// ID is the numeric league identifier.
	ID int `json:"id"`
	// Name is the league display name.
	Name string `json:"name"`
	// Icon contains league icon URLs when the endpoint provides them.
	Icon *Icon `json:"iconUrls"`
	responseMeta
}

// Season describes one ranked season placement.
type Season struct {
	// ID is the season identifier, usually YYYY-MM.
	ID string `json:"id"`
	// Rank is the player's season rank.
	Rank int `json:"rank"`
	// Trophies is the player's trophy count for the season.
	Trophies int `json:"trophies"`
}

// LegendStatistics contains a player's legend trophies and season snapshots.
type LegendStatistics struct {
	// LegendTrophies is the player's lifetime legend trophy count.
	LegendTrophies int `json:"legendTrophies"`
	// BestSeason is the player's best legend season.
	BestSeason *Season `json:"bestSeason"`
	// PreviousSeason is the player's previous legend season.
	PreviousSeason *Season `json:"previousSeason"`
	// BestVersusSeason is the player's legacy best Builder Base season.
	BestVersusSeason *Season `json:"bestVersusSeason"`
	// CurrentSeason is the player's current legend season progress.
	CurrentSeason *Season `json:"currentSeason"`
}

// Label is a player or clan label.
type Label struct {
	// ID is the label identifier.
	ID int `json:"id"`
	// Name is the label display name.
	Name string `json:"name"`
	// Icon contains label icon URLs.
	Icon *Icon `json:"iconUrls"`
	responseMeta
}

// CapitalDistrict describes a clan capital district from clan and raid data.
type CapitalDistrict struct {
	// ID is the district identifier.
	ID int `json:"id"`
	// Name is the district display name.
	Name string `json:"name"`
	// DistrictHallLevel is the district hall level.
	DistrictHallLevel int `json:"districtHallLevel"`
	// DestructionPercent is the destruction percentage in raid contexts.
	DestructionPercent float64 `json:"destructionPercent"`
	// AttackCount is the number of attacks used against the district.
	AttackCount int `json:"attackCount"`
	// Looted is the total capital gold looted from the district.
	Looted int `json:"totalLooted"`
}

// ChatLanguage describes the preferred language configured for a clan.
type ChatLanguage struct {
	// ID is the language identifier.
	ID int `json:"id"`
	// Name is the language display name.
	Name string `json:"name"`
	// LanguageCode is the language code returned by the API.
	LanguageCode string `json:"languageCode"`
}

// GoldPassSeason describes the current Gold Pass season.
type GoldPassSeason struct {
	// StartTime is when the Gold Pass season starts.
	StartTime *Timestamp `json:"startTime"`
	// EndTime is when the Gold Pass season ends.
	EndTime *Timestamp `json:"endTime"`
	responseMeta
}

// PlayerHouseElement is one cosmetic element of a player's house.
type PlayerHouseElement struct {
	// ID is the cosmetic element identifier.
	ID int `json:"id"`
	// Type is the cosmetic element type.
	Type string `json:"type"`
}

// Translation contains one static-data translation entry.
type Translation struct {
	// ID is the translation identifier.
	ID string `json:"id"`
	// English is the EN translation value.
	English string `json:"EN"`
	// Languages maps language codes to translated strings.
	Languages map[string]string `json:"-"`
}

// UnmarshalJSON stores all language entries and promotes EN into English.
func (t *Translation) UnmarshalJSON(data []byte) error {
	var payload map[string]string
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}
	t.Languages = payload
	t.English = payload["EN"]
	return nil
}

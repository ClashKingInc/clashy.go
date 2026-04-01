package clashy

import (
	"encoding/json"
	"time"
)

type Timestamp struct {
	RawTime string
	Time    time.Time
}

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

func (t Timestamp) SecondsUntil() int {
	if t.Time.IsZero() {
		return 0
	}
	return int(time.Until(t.Time).Seconds())
}

func (t Timestamp) Before(other Timestamp) bool { return t.Time.Before(other.Time) }
func (t Timestamp) After(other Timestamp) bool  { return t.Time.After(other.Time) }

type TimeDelta struct {
	time.Duration
}

type Badge struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
	URL    string `json:"-"`
}

func (b *Badge) finalize() {
	if b == nil {
		return
	}
	if b.URL == "" {
		b.URL = b.Medium
	}
}

func (b *Badge) postDecode(*Client) {
	b.finalize()
}

type Icon struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Tiny   string `json:"tiny"`
}

type Achievement struct {
	Name           string `json:"name"`
	Stars          int    `json:"stars"`
	Value          int    `json:"value"`
	Target         int    `json:"target"`
	Info           string `json:"info"`
	CompletionInfo string `json:"completionInfo"`
	Village        string `json:"village"`
}

type Location struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsCountry   bool   `json:"isCountry"`
	CountryCode string `json:"countryCode"`
	Localised   string `json:"localizedName"`
}

type BaseLeague struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon *Icon  `json:"iconUrls"`
}

type League = BaseLeague

type Season struct {
	ID       string `json:"id"`
	Rank     int    `json:"rank"`
	Trophies int    `json:"trophies"`
}

type LegendStatistics struct {
	LegendTrophies   int     `json:"legendTrophies"`
	BestSeason       *Season `json:"bestSeason"`
	PreviousSeason   *Season `json:"previousSeason"`
	BestVersusSeason *Season `json:"bestVersusSeason"`
	CurrentSeason    *Season `json:"currentSeason"`
}

type Label struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon *Icon  `json:"iconUrls"`
}

type CapitalDistrict struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	DistrictHallLevel  int     `json:"districtHallLevel"`
	DestructionPercent float64 `json:"destructionPercent"`
	AttackCount        int     `json:"attackCount"`
	Looted             int     `json:"totalLooted"`
}

type ChatLanguage struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	LanguageCode string `json:"languageCode"`
}

type GoldPassSeason struct {
	StartTime *Timestamp `json:"startTime"`
	EndTime   *Timestamp `json:"endTime"`
}

type PlayerHouseElement struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type Translation struct {
	ID        string            `json:"id"`
	English   string            `json:"EN"`
	Languages map[string]string `json:"-"`
}

func (t *Translation) UnmarshalJSON(data []byte) error {
	type alias Translation
	var payload map[string]string
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}
	t.Languages = payload
	t.English = payload["EN"]
	return nil
}

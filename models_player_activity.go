package clashy

import (
	"encoding/json"
	"fmt"
	"strings"
)

type FlexibleID string

func (id *FlexibleID) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*id = ""
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*id = FlexibleID(str)
		return nil
	}

	var num json.Number
	if err := json.Unmarshal(data, &num); err == nil {
		*id = FlexibleID(num.String())
		return nil
	}

	return fmt.Errorf("invalid flexible id: %s", strings.TrimSpace(string(data)))
}

func (id FlexibleID) String() string { return string(id) }

type Resource struct {
	Name   string `json:"name,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

type BattleLogEntry struct {
	BattleType            string     `json:"battleType,omitempty"`
	Attack                bool       `json:"attack,omitempty"`
	ArmyShareCode         string     `json:"armyShareCode,omitempty"`
	OpponentPlayerTag     string     `json:"opponentPlayerTag,omitempty"`
	Stars                 int        `json:"stars,omitempty"`
	DestructionPercentage int        `json:"destructionPercentage,omitempty"`
	LootedResources       []Resource `json:"lootedResources,omitempty"`
	ExtraLootedResources  []Resource `json:"extraLootedResources,omitempty"`
	AvailableLoot         []Resource `json:"availableLoot,omitempty"`
	responseMeta
}

type LeagueHistoryEntry struct {
	LeagueSeasonID FlexibleID `json:"leagueSeasonId,omitempty"`
	LeagueTrophies int        `json:"leagueTrophies,omitempty"`
	LeagueTierID   int        `json:"leagueTierId,omitempty"`
	Placement      int        `json:"placement,omitempty"`
	AttackWins     int        `json:"attackWins,omitempty"`
	AttackLosses   int        `json:"attackLosses,omitempty"`
	AttackStars    int        `json:"attackStars,omitempty"`
	DefenseWins    int        `json:"defenseWins,omitempty"`
	DefenseLosses  int        `json:"defenseLosses,omitempty"`
	DefenseStars   int        `json:"defenseStars,omitempty"`
	MaxBattles     int        `json:"maxBattles,omitempty"`
	responseMeta
}

type LeagueTierGroupBattleLogEntry struct {
	OpponentPlayerTag     string `json:"opponentPlayerTag,omitempty"`
	OpponentName          string `json:"opponentName,omitempty"`
	Stars                 int    `json:"stars,omitempty"`
	DestructionPercentage int    `json:"destructionPercentage,omitempty"`
	Trophies              int    `json:"trophies,omitempty"`
	CreationTime          string `json:"creationTime,omitempty"`
}

type LeagueTierGroupMember struct {
	PlayerTag        string `json:"playerTag,omitempty"`
	PlayerName       string `json:"playerName,omitempty"`
	ClanTag          string `json:"clanTag,omitempty"`
	ClanName         string `json:"clanName,omitempty"`
	LeagueTrophies   int    `json:"leagueTrophies,omitempty"`
	AttackWinCount   int    `json:"attackWinCount,omitempty"`
	AttackLoseCount  int    `json:"attackLoseCount,omitempty"`
	DefenseWinCount  int    `json:"defenseWinCount,omitempty"`
	DefenseLoseCount int    `json:"defenseLoseCount,omitempty"`
}

type LeagueTierGroup struct {
	Members     []LeagueTierGroupMember         `json:"members,omitempty"`
	AttackLogs  []LeagueTierGroupBattleLogEntry `json:"attackLogs,omitempty"`
	DefenseLogs []LeagueTierGroupBattleLogEntry `json:"defenseLogs,omitempty"`
	responseMeta
}

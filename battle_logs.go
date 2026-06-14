package clashy

// Resource is a named resource amount in a player battle log entry.
type Resource struct {
	// Name is the resource name, such as gold, elixir, or dark elixir.
	Name string `json:"name,omitempty"`
	// Amount is the resource quantity.
	Amount int `json:"amount,omitempty"`
}

// BattleLogEntry is one player battle log entry.
type BattleLogEntry struct {
	// BattleType describes the game mode for the battle.
	BattleType string `json:"battleType,omitempty"`
	// Attack reports whether the entry is an attack made by the requested
	// player. False entries are defenses.
	Attack bool `json:"attack,omitempty"`
	// ArmyShareCode is the in-game army share payload when available.
	ArmyShareCode string `json:"armyShareCode,omitempty"`
	// OpponentPlayerTag is the opponent's player tag.
	OpponentPlayerTag string `json:"opponentPlayerTag,omitempty"`
	// Stars is the number of stars earned by the attacker.
	Stars int `json:"stars,omitempty"`
	// DestructionPercentage is the destruction percentage earned by the attacker.
	DestructionPercentage int `json:"destructionPercentage,omitempty"`
	// LootedResources contains resources actually looted.
	LootedResources []Resource `json:"lootedResources,omitempty"`
	// ExtraLootedResources contains bonus resources awarded by the battle.
	ExtraLootedResources []Resource `json:"extraLootedResources,omitempty"`
	// AvailableLoot contains resources that were available before the battle.
	AvailableLoot []Resource `json:"availableLoot,omitempty"`
	responseMeta
}

// LeagueHistoryEntry is one historical legend-league season result.
type LeagueHistoryEntry struct {
	// LeagueSeasonID is the numeric legend season identifier.
	LeagueSeasonID int `json:"leagueSeasonId,omitempty"`
	// LeagueTrophies is the player's ending legend trophy count.
	LeagueTrophies int `json:"leagueTrophies,omitempty"`
	// LeagueTierID is the league tier identifier for the season.
	LeagueTierID int `json:"leagueTierId,omitempty"`
	// Placement is the player's final placement.
	Placement int `json:"placement,omitempty"`
	// AttackWins is the number of attack wins.
	AttackWins int `json:"attackWins,omitempty"`
	// AttackLosses is the number of attack losses.
	AttackLosses int `json:"attackLosses,omitempty"`
	// AttackStars is the total stars earned on attack.
	AttackStars int `json:"attackStars,omitempty"`
	// DefenseWins is the number of defense wins.
	DefenseWins int `json:"defenseWins,omitempty"`
	// DefenseLosses is the number of defense losses.
	DefenseLosses int `json:"defenseLosses,omitempty"`
	// DefenseStars is the total stars allowed on defense.
	DefenseStars int `json:"defenseStars,omitempty"`
	// MaxBattles is the maximum battle count for the season.
	MaxBattles int `json:"maxBattles,omitempty"`
	responseMeta
}

// LeagueTierGroupBattleLogEntry is one attack or defense inside a legend group.
type LeagueTierGroupBattleLogEntry struct {
	// OpponentPlayerTag is the opponent's player tag.
	OpponentPlayerTag string `json:"opponentPlayerTag,omitempty"`
	// OpponentName is the opponent's player name.
	OpponentName string `json:"opponentName,omitempty"`
	// Stars is the number of stars earned by the attacker.
	Stars int `json:"stars,omitempty"`
	// DestructionPercentage is the destruction percentage earned by the attacker.
	DestructionPercentage int `json:"destructionPercentage,omitempty"`
	// Trophies is the trophy delta for the battle.
	Trophies int `json:"trophies,omitempty"`
	// CreationTime is the API timestamp for the battle.
	CreationTime string `json:"creationTime,omitempty"`
}

// LeagueTierGroupMember is one player in a legend league group.
type LeagueTierGroupMember struct {
	// PlayerTag is the player's tag.
	PlayerTag string `json:"playerTag,omitempty"`
	// PlayerName is the player's display name.
	PlayerName string `json:"playerName,omitempty"`
	// ClanTag is the player's clan tag when present.
	ClanTag string `json:"clanTag,omitempty"`
	// ClanName is the player's clan name when present.
	ClanName string `json:"clanName,omitempty"`
	// LeagueTrophies is the player's current legend trophy count.
	LeagueTrophies int `json:"leagueTrophies,omitempty"`
	// AttackWinCount is the player's attack win count in the group.
	AttackWinCount int `json:"attackWinCount,omitempty"`
	// AttackLoseCount is the player's attack loss count in the group.
	AttackLoseCount int `json:"attackLoseCount,omitempty"`
	// DefenseWinCount is the player's defense win count in the group.
	DefenseWinCount int `json:"defenseWinCount,omitempty"`
	// DefenseLoseCount is the player's defense loss count in the group.
	DefenseLoseCount int `json:"defenseLoseCount,omitempty"`
}

// LeagueTierGroup contains members and battle logs for a legend league group.
type LeagueTierGroup struct {
	// Members contains the players in the legend group.
	Members []LeagueTierGroupMember `json:"members,omitempty"`
	// AttackLogs contains attack entries for the requested player.
	AttackLogs []LeagueTierGroupBattleLogEntry `json:"attackLogs,omitempty"`
	// DefenseLogs contains defense entries for the requested player.
	DefenseLogs []LeagueTierGroupBattleLogEntry `json:"defenseLogs,omitempty"`
	responseMeta
}

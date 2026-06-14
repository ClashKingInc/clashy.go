package clashy

// Player is the full player profile returned by GetPlayer.
type Player struct {
	// Tag is the player's Clash tag.
	Tag string `json:"tag,omitempty"`
	// Name is the player's current display name.
	Name string `json:"name,omitempty"`
	// ExpLevel is the player's experience level.
	ExpLevel int `json:"expLevel,omitempty"`
	// Trophies is the player's current home village trophy count.
	Trophies int `json:"trophies,omitempty"`
	// BestTrophies is the player's all-time best home village trophy count.
	BestTrophies int `json:"bestTrophies,omitempty"`
	// WarStars is the player's lifetime war star count.
	WarStars int `json:"warStars,omitempty"`
	// TownHall is the player's home village Town Hall level.
	TownHall int `json:"townHallLevel,omitempty"`
	// TownHallWeapon is the weapon level for Town Hall levels that have one.
	TownHallWeapon int `json:"townHallWeaponLevel,omitempty"`
	// BuilderHall is the player's Builder Hall level.
	BuilderHall int `json:"builderHallLevel,omitempty"`
	// BestBuilderBaseTrophies is the all-time best Builder Base trophy count.
	BestBuilderBaseTrophies int `json:"bestBuilderBaseTrophies,omitempty"`
	// VersusAttackWins is the legacy Builder Base attack-win field.
	VersusAttackWins int `json:"versusBattleWins,omitempty"`
	// Donations is the number of troops donated this season.
	Donations int `json:"donations,omitempty"`
	// Received is the number of donated troops received this season.
	Received int `json:"donationsReceived,omitempty"`
	// ClanCapitalContributions is the lifetime Clan Capital contribution count.
	ClanCapitalContributions int `json:"clanCapitalContributions,omitempty"`
	// ClanRank is the player's current rank inside their clan.
	ClanRank int `json:"clanRank,omitempty"`
	// ClanPreviousRank is the player's previous rank inside their clan.
	ClanPreviousRank int `json:"previousClanRank,omitempty"`
	// VersusTrophies is the legacy Builder Base trophy field.
	VersusTrophies int `json:"versusTrophies,omitempty"`
	// BuilderBaseTrophies is the player's current Builder Base trophy count.
	BuilderBaseTrophies int `json:"builderBaseTrophies,omitempty"`
	// LeagueTier is the player's home village league.
	LeagueTier League `json:"leagueTier,omitempty"`
	// BuilderBaseLeague is the player's Builder Base league.
	BuilderBaseLeague *League `json:"builderBaseLeague,omitempty"`
	// Role is the player's role in their current clan.
	Role Role `json:"role,omitempty"`
	// Clan is the compact clan object for the player's current clan.
	Clan *PlayerClan `json:"clan,omitempty"`
	// CurrentLeagueGroupTag is the active legend league group tag when present.
	CurrentLeagueGroupTag string `json:"currentLeagueGroupTag,omitempty"`
	// CurrentLeagueSeasonID is the active legend league season ID when present.
	CurrentLeagueSeasonID int `json:"currentLeagueSeasonId,omitempty"`
	// PreviousLeagueGroupTag is the previous legend league group tag when
	// present.
	PreviousLeagueGroupTag string `json:"previousLeagueGroupTag,omitempty"`
	// PreviousLeagueSeasonID is the previous legend league season ID when
	// present.
	PreviousLeagueSeasonID int `json:"previousLeagueSeasonId,omitempty"`
	// LegendStatistics contains legend trophies and seasonal legend finishes.
	LegendStatistics *LegendStatistics `json:"legendStatistics,omitempty"`
	// Labels are public player labels.
	Labels []Label `json:"labels,omitempty"`
	// Achievements contains achievement progress for both villages and Clan
	// Capital.
	Achievements []Achievement `json:"achievements,omitempty"`
	// Troops contains unlocked troops with current and max levels.
	Troops []Troop `json:"troops,omitempty"`
	// Heroes contains unlocked heroes with current and max levels.
	Heroes []Hero `json:"heroes,omitempty"`
	// Spells contains unlocked spells with current and max levels.
	Spells []Spell `json:"spells,omitempty"`
	// HeroEquipment contains unlocked hero equipment with current and max levels.
	HeroEquipment []Equipment `json:"heroEquipment,omitempty"`
	responseMeta
}

// GetAchievement returns the achievement with the provided display name.
func (p *Player) GetAchievement(name string) *Achievement {
	for i := range p.Achievements {
		if p.Achievements[i].Name == name {
			return &p.Achievements[i]
		}
	}
	return nil
}

// GetHero returns the hero with the provided display name.
func (p *Player) GetHero(name string) *Hero {
	for i := range p.Heroes {
		if p.Heroes[i].Name == name {
			return &p.Heroes[i]
		}
	}
	return nil
}

// GetTroop returns the troop with the provided display name.
func (p *Player) GetTroop(name string) *Troop {
	for i := range p.Troops {
		if p.Troops[i].Name == name {
			return &p.Troops[i]
		}
	}
	return nil
}

// GetSpell returns the spell with the provided display name.
func (p *Player) GetSpell(name string) *Spell {
	for i := range p.Spells {
		if p.Spells[i].Name == name {
			return &p.Spells[i]
		}
	}
	return nil
}

// HomeTroops returns troops that belong to the home village.
//
// Older API responses may omit Village for home-village troops, so an empty
// village is treated as home.
func (p *Player) HomeTroops() []Troop {
	var out []Troop
	for _, troop := range p.Troops {
		if troop.Village == string(VillageHome) || troop.Village == "" {
			out = append(out, troop)
		}
	}
	return out
}

// BuilderTroops returns troops that belong to Builder Base.
func (p *Player) BuilderTroops() []Troop {
	var out []Troop
	for _, troop := range p.Troops {
		if troop.Village == string(VillageBuilderBase) {
			out = append(out, troop)
		}
	}
	return out
}

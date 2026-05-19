package clashy

type Player struct {
	Tag                     string            `json:"tag,omitempty"`
	Name                    string            `json:"name,omitempty"`
	ExpLevel                int               `json:"expLevel,omitempty"`
	Trophies                int               `json:"trophies,omitempty"`
	BestTrophies            int               `json:"bestTrophies,omitempty"`
	WarStars                int               `json:"warStars,omitempty"`
	TownHall                int               `json:"townHallLevel,omitempty"`
	TownHallWeapon          int               `json:"townHallWeaponLevel,omitempty"`
	BuilderHall             int               `json:"builderHallLevel,omitempty"`
	BestBuilderBaseTrophies int               `json:"bestBuilderBaseTrophies,omitempty"`
	VersusAttackWins        int               `json:"versusBattleWins,omitempty"`
	Donations               int               `json:"donations,omitempty"`
	Received                int               `json:"donationsReceived,omitempty"`
	ClanRank                int               `json:"clanRank,omitempty"`
	ClanPreviousRank        int               `json:"previousClanRank,omitempty"`
	VersusTrophies          int               `json:"versusTrophies,omitempty"`
	BuilderBaseTrophies     int               `json:"builderBaseTrophies,omitempty"`
	LeagueTier              League            `json:"leagueTier,omitempty"`
	BuilderBaseLeague       *League           `json:"builderBaseLeague,omitempty"`
	Role                    Role              `json:"role,omitempty"`
	Clan                    *PlayerClan       `json:"clan,omitempty"`
	CurrentLeagueGroupTag   string            `json:"currentLeagueGroupTag,omitempty"`
	CurrentLeagueSeasonID   int               `json:"currentLeagueSeasonId,omitempty"`
	PreviousLeagueGroupTag  string            `json:"previousLeagueGroupTag,omitempty"`
	PreviousLeagueSeasonID  int               `json:"previousLeagueSeasonId,omitempty"`
	LegendStatistics        *LegendStatistics `json:"legendStatistics,omitempty"`
	Labels                  []Label           `json:"labels,omitempty"`
	Achievements            []Achievement     `json:"achievements,omitempty"`
	Troops                  []Troop           `json:"troops,omitempty"`
	Heroes                  []Hero            `json:"heroes,omitempty"`
	Spells                  []Spell           `json:"spells,omitempty"`
	HeroEquipment           []Equipment       `json:"heroEquipment,omitempty"`
	responseMeta
}

func (p *Player) GetAchievement(name string) *Achievement {
	for i := range p.Achievements {
		if p.Achievements[i].Name == name {
			return &p.Achievements[i]
		}
	}
	return nil
}

func (p *Player) GetHero(name string) *Hero {
	for i := range p.Heroes {
		if p.Heroes[i].Name == name {
			return &p.Heroes[i]
		}
	}
	return nil
}

func (p *Player) GetTroop(name string) *Troop {
	for i := range p.Troops {
		if p.Troops[i].Name == name {
			return &p.Troops[i]
		}
	}
	return nil
}

func (p *Player) GetSpell(name string) *Spell {
	for i := range p.Spells {
		if p.Spells[i].Name == name {
			return &p.Spells[i]
		}
	}
	return nil
}

func (p *Player) HomeTroops() []Troop {
	var out []Troop
	for _, troop := range p.Troops {
		if troop.Village == string(VillageHome) || troop.Village == "" {
			out = append(out, troop)
		}
	}
	return out
}

func (p *Player) BuilderTroops() []Troop {
	var out []Troop
	for _, troop := range p.Troops {
		if troop.Village == string(VillageBuilderBase) {
			out = append(out, troop)
		}
	}
	return out
}

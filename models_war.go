package clashy

type WarAttack struct {
	Order       int     `json:"order,omitempty"`
	AttackerTag string  `json:"attackerTag,omitempty"`
	DefenderTag string  `json:"defenderTag,omitempty"`
	Stars       int     `json:"stars,omitempty"`
	Destruction float64 `json:"destructionPercentage,omitempty"`
	Duration    int     `json:"duration,omitempty"`
	Attacker    *ClanWarMember
	Defender    *ClanWarMember
}

type ClanWarMember struct {
	Tag                string      `json:"tag,omitempty"`
	Name               string      `json:"name,omitempty"`
	MapPosition        int         `json:"mapPosition,omitempty"`
	Townhall           int         `json:"townhallLevel,omitempty"`
	OpponentAttacks    int         `json:"opponentAttacks,omitempty"`
	Attacks            []WarAttack `json:"attacks,omitempty"`
	BestOpponentAttack *WarAttack  `json:"bestOpponentAttack,omitempty"`
	Clan               *WarClan
}

type WarClan struct {
	Tag         string          `json:"tag,omitempty"`
	Name        string          `json:"name,omitempty"`
	Badge       Badge           `json:"badgeUrls,omitempty"`
	Level       int             `json:"clanLevel,omitempty"`
	Attacks     int             `json:"attacks,omitempty"`
	Stars       int             `json:"stars,omitempty"`
	Destruction float64         `json:"destructionPercentage,omitempty"`
	ExpEarned   int             `json:"expEarned,omitempty"`
	Members     []ClanWarMember `json:"members,omitempty"`
}

type ClanWar struct {
	State                WarState   `json:"state,omitempty"`
	TeamSize             int        `json:"teamSize,omitempty"`
	PreparationStartTime *Timestamp `json:"preparationStartTime,omitempty"`
	StartTime            *Timestamp `json:"startTime,omitempty"`
	EndTime              *Timestamp `json:"endTime,omitempty"`
	Clan                 *WarClan   `json:"clan,omitempty"`
	Opponent             *WarClan   `json:"opponent,omitempty"`
	BattleModifier       string     `json:"battleModifier,omitempty"`
	WarTag               string     `json:"tag,omitempty"`
	ClanTag              string     `json:"-"`
	responseMeta
}

func (w *ClanWar) Type() string {
	if w == nil || w.WarTag == "" {
		return "random"
	}
	return "cwl"
}

func (w *ClanWar) Attacks() []WarAttack {
	if w == nil {
		return nil
	}
	var out []WarAttack
	for i := range w.Clan.Members {
		for _, attack := range w.Clan.Members[i].Attacks {
			out = append(out, attack)
		}
	}
	for i := range w.Opponent.Members {
		for _, attack := range w.Opponent.Members[i].Attacks {
			out = append(out, attack)
		}
	}
	return out
}

type ClanWarLogEntry struct {
	Result   WarResult  `json:"result,omitempty"`
	EndTime  *Timestamp `json:"endTime,omitempty"`
	TeamSize int        `json:"teamSize,omitempty"`
	Clan     *WarClan   `json:"clan,omitempty"`
	Opponent *WarClan   `json:"opponent,omitempty"`
}

type ClanWarLeagueClan struct {
	Tag   string `json:"tag,omitempty"`
	Name  string `json:"name,omitempty"`
	Badge Badge  `json:"badgeUrls,omitempty"`
	Level int    `json:"clanLevel,omitempty"`
}

type ClanWarLeagueGroup struct {
	State  string              `json:"state,omitempty"`
	Season string              `json:"season,omitempty"`
	Clans  []ClanWarLeagueClan `json:"clans,omitempty"`
	Rounds []struct {
		WarTags []string `json:"warTags,omitempty"`
	} `json:"rounds,omitempty"`
}

type ExtendedCWLGroup struct {
	Name              string `json:"name,omitempty"`
	FirstPlaceMedals  int    `json:"first_place_medals,omitempty"`
	SecondPlaceMedals int    `json:"second_place_medals,omitempty"`
}

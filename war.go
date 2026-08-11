package clashy

// WarAttack is one attack inside a classic war or Clan War League war.
type WarAttack struct {
	// Order is the attack order assigned by the API.
	Order int `json:"order,omitempty"`
	// AttackerTag is the player tag of the attacker.
	AttackerTag string `json:"attackerTag,omitempty"`
	// DefenderTag is the player tag of the defender.
	DefenderTag string `json:"defenderTag,omitempty"`
	// Stars is the number of stars earned by the attack.
	Stars int `json:"stars,omitempty"`
	// Destruction is the destruction percentage earned by the attack.
	Destruction float64 `json:"destructionPercentage,omitempty"`
	// Duration is the attack duration in seconds.
	Duration int `json:"duration,omitempty"`
}

// ClanWarMember is a player entry on one side of a war.
type ClanWarMember struct {
	// Tag is the player's tag.
	Tag string `json:"tag,omitempty"`
	// Name is the player's display name at the time of the war.
	Name string `json:"name,omitempty"`
	// MapPosition is the player's position on the war map.
	MapPosition int `json:"mapPosition,omitempty"`
	// Townhall is the player's Town Hall level in the war response.
	Townhall int `json:"townhallLevel,omitempty"`
	// OpponentAttacks is the number of attacks used against this base.
	OpponentAttacks int `json:"opponentAttacks,omitempty"`
	// Attacks contains attacks made by this member.
	Attacks []WarAttack `json:"attacks,omitempty"`
	// BestOpponentAttack is the best attack received by this member.
	BestOpponentAttack *WarAttack `json:"bestOpponentAttack,omitempty"`
}

// WarClan is one clan side of a classic war or CWL war.
type WarClan struct {
	// Tag is the clan tag.
	Tag string `json:"tag,omitempty"`
	// Name is the clan name.
	Name string `json:"name,omitempty"`
	// Badge contains clan badge image URLs.
	Badge Badge `json:"badgeUrls,omitempty"`
	// Level is the clan level.
	Level int `json:"clanLevel,omitempty"`
	// Attacks is the number of attacks used by this clan.
	Attacks int `json:"attacks,omitempty"`
	// Stars is the total stars earned by this clan.
	Stars int `json:"stars,omitempty"`
	// Destruction is the total destruction percentage earned by this clan.
	Destruction float64 `json:"destructionPercentage,omitempty"`
	// ExpEarned is clan XP earned by this war when the endpoint includes it.
	ExpEarned int `json:"expEarned,omitempty"`
	// Members is the war roster for this side.
	Members []ClanWarMember `json:"members,omitempty"`
}

// ClanWar is the current, historical, or league war response.
//
// For league wars found through GetCurrentWar or GetLeagueWar, the client
// orients the result so Clan is the requested clan and Opponent is the opposing
// side, even if the API returned the requested clan under opponent.
type ClanWar struct {
	// State is the current state of the war.
	State WarState `json:"state,omitempty"`
	// TeamSize is the roster size for each side.
	TeamSize int `json:"teamSize,omitempty"`
	// PreparationStartTime is when preparation day began.
	PreparationStartTime *Timestamp `json:"preparationStartTime,omitempty"`
	// StartTime is when battle day starts.
	StartTime *Timestamp `json:"startTime,omitempty"`
	// EndTime is when the war ends.
	EndTime *Timestamp `json:"endTime,omitempty"`
	// Clan is the requested clan side for oriented responses.
	Clan *WarClan `json:"clan,omitempty"`
	// Opponent is the opposing clan side for oriented responses.
	Opponent *WarClan `json:"opponent,omitempty"`
	// BattleModifier describes event-specific modifiers when the API includes
	// one.
	BattleModifier BattleModifier `json:"battleModifier,omitempty"`
	// WarTag is the CWL war tag. It is empty for normal classic wars.
	WarTag string `json:"tag,omitempty"`
	// ClanTag is the requested clan tag associated with this response.
	ClanTag string `json:"-"`
	// LeagueGroup is the CWL group used to find this war when available.
	LeagueGroup *ClanWarLeagueGroup `json:"-"`
	responseMeta
}

// Type returns "cwl" when the war has a CWL war tag and "random" otherwise.
func (w *ClanWar) Type() string {
	if w == nil || w.WarTag == "" {
		return "random"
	}
	return "cwl"
}

// Attacks returns all attacks made by both sides of the war.
func (w *ClanWar) Attacks() []WarAttack {
	if w == nil {
		return nil
	}
	total := 0
	for _, side := range []*WarClan{w.Clan, w.Opponent} {
		if side == nil {
			continue
		}
		for i := range side.Members {
			total += len(side.Members[i].Attacks)
		}
	}
	out := make([]WarAttack, 0, total)
	if w.Clan != nil {
		for i := range w.Clan.Members {
			for _, attack := range w.Clan.Members[i].Attacks {
				out = append(out, attack)
			}
		}
	}
	if w.Opponent != nil {
		for i := range w.Opponent.Members {
			for _, attack := range w.Opponent.Members[i].Attacks {
				out = append(out, attack)
			}
		}
	}
	return out
}

// Member returns the war member with tag from either side of the war.
func (w *ClanWar) Member(tag string) *ClanWarMember {
	if w == nil {
		return nil
	}
	tag = CorrectTag(tag)
	for _, side := range []*WarClan{w.Clan, w.Opponent} {
		if side == nil {
			continue
		}
		for i := range side.Members {
			if CorrectTag(side.Members[i].Tag) == tag {
				return &side.Members[i]
			}
		}
	}
	return nil
}

// ResolveAttack resolves an attack's tags to the matching war members.
func (w *ClanWar) ResolveAttack(attack WarAttack) (attacker, defender *ClanWarMember) {
	return w.Member(attack.AttackerTag), w.Member(attack.DefenderTag)
}

// ClanWarLogEntry is one item from a clan's public war log.
type ClanWarLogEntry struct {
	// Result is the requested clan's result for this war.
	Result WarResult `json:"result,omitempty"`
	// EndTime is when the war ended.
	EndTime *Timestamp `json:"endTime,omitempty"`
	// TeamSize is the roster size for each side.
	TeamSize int `json:"teamSize,omitempty"`
	// Clan is the requested clan side.
	Clan *WarClan `json:"clan,omitempty"`
	// Opponent is the opposing clan side.
	Opponent *WarClan `json:"opponent,omitempty"`
	// BattleModifier describes event-specific modifiers when the API includes
	// one.
	BattleModifier BattleModifier `json:"battleModifier,omitempty"`
	responseMeta
}

// ClanWarLeagueClan is a clan entry inside a CWL group.
type ClanWarLeagueClan struct {
	// Tag is the clan tag.
	Tag string `json:"tag,omitempty"`
	// Name is the clan name.
	Name string `json:"name,omitempty"`
	// Badge contains clan badge image URLs.
	Badge Badge `json:"badgeUrls,omitempty"`
	// Level is the clan level.
	Level int `json:"clanLevel,omitempty"`
	// Members is the frozen master roster registered for this CWL clan.
	Members []ClanWarLeagueClanMember `json:"members,omitempty"`
}

// ClanWarLeagueClanMember is one player in a CWL clan's master roster.
type ClanWarLeagueClanMember struct {
	// Tag is the player's tag.
	Tag string `json:"tag,omitempty"`
	// Name is the player's name at registration time.
	Name string `json:"name,omitempty"`
	// TownHallLevel is the player's Town Hall level at registration time.
	TownHallLevel int `json:"townHallLevel,omitempty"`
}

// ClanWarLeagueGroup is the current CWL group for a clan.
type ClanWarLeagueGroup struct {
	// State is the group state returned by the API.
	State string `json:"state,omitempty"`
	// Season is the CWL season identifier returned by the API.
	Season string `json:"season,omitempty"`
	// Clans contains the clans participating in the group.
	Clans []ClanWarLeagueClan `json:"clans,omitempty"`
	// Rounds contains CWL war tags grouped by round. Future rounds may contain
	// placeholder "#0" tags.
	Rounds []struct {
		WarTags []string `json:"warTags,omitempty"`
	} `json:"rounds,omitempty"`
	responseMeta
}

// ExtendedCWLGroup contains static medal information for a CWL league.
type ExtendedCWLGroup struct {
	// Name is the league display name.
	Name string `json:"name,omitempty"`
	// FirstPlaceMedals is the medal reward for first place.
	FirstPlaceMedals int `json:"first_place_medals,omitempty"`
	// SecondPlaceMedals is the medal reward for second place.
	SecondPlaceMedals int `json:"second_place_medals,omitempty"`
}

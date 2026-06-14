package clashy

// RaidAttack is one attack against a Clan Capital district.
type RaidAttack struct {
	// AttackerTag can be populated by callers that join attacks to raid members.
	AttackerTag string `json:"-"`
	// AttackerName can be populated by callers that join attacks to raid
	// members.
	AttackerName string `json:"-"`
	// Stars is the star count earned by this district attack.
	Stars int `json:"stars,omitempty"`
	// Destruction is the destruction percentage earned by this district attack.
	Destruction float64 `json:"destructionPercent,omitempty"`
}

// RaidDistrict describes one district in a raid attack or defense log.
type RaidDistrict struct {
	// ID is the district identifier.
	ID int `json:"id,omitempty"`
	// Name is the district display name.
	Name string `json:"name,omitempty"`
	// HallLevel is the district hall level.
	HallLevel int `json:"districtHallLevel,omitempty"`
	// Destruction is the final destruction percentage for the district.
	Destruction float64 `json:"destructionPercent,omitempty"`
	// AttackCount is the number of attacks used against the district.
	AttackCount int `json:"attackCount,omitempty"`
	// Looted is the total capital gold looted from the district.
	Looted int `json:"totalLooted,omitempty"`
	// Attacks contains individual attacks against the district.
	Attacks []RaidAttack `json:"attacks,omitempty"`
}

// RaidClan describes one opposing clan entry in a raid attack or defense log.
//
// Attack log entries use Attacker for the clan being attacked by the requested
// clan. Defense log entries use Defender for the clan that attacked the
// requested clan.
type RaidClan struct {
	// AttackCount is the number of attacks used in this raid.
	AttackCount int `json:"attackCount,omitempty"`
	// DistrictCount is the number of districts available.
	DistrictCount int `json:"districtCount,omitempty"`
	// DestroyedDistrictCount is the number of districts destroyed.
	DestroyedDistrictCount int `json:"districtsDestroyed,omitempty"`
	// Districts contains district-level attack details.
	Districts []RaidDistrict `json:"districts,omitempty"`
	// Attacker is set on attack-log entries.
	Attacker *struct {
		Tag   string `json:"tag,omitempty"`
		Name  string `json:"name,omitempty"`
		Level int    `json:"level,omitempty"`
	} `json:"attacker,omitempty"`
	// Defender is set on defense-log entries.
	Defender *struct {
		Tag   string `json:"tag,omitempty"`
		Name  string `json:"name,omitempty"`
		Level int    `json:"level,omitempty"`
	} `json:"defender,omitempty"`
}

// Tag returns the attacker or defender tag for this raid clan entry.
func (r RaidClan) Tag() string {
	if r.Attacker != nil {
		return r.Attacker.Tag
	}
	if r.Defender != nil {
		return r.Defender.Tag
	}
	return ""
}

// Name returns the attacker or defender name for this raid clan entry.
func (r RaidClan) Name() string {
	if r.Attacker != nil {
		return r.Attacker.Name
	}
	if r.Defender != nil {
		return r.Defender.Name
	}
	return ""
}

// Level returns the attacker or defender clan level for this raid clan entry.
func (r RaidClan) Level() int {
	if r.Attacker != nil {
		return r.Attacker.Level
	}
	if r.Defender != nil {
		return r.Defender.Level
	}
	return 0
}

// RaidMember is one clan member's contribution in a raid weekend.
type RaidMember struct {
	// Tag is the member's player tag.
	Tag string `json:"tag,omitempty"`
	// Name is the member's display name.
	Name string `json:"name,omitempty"`
	// AttackCount is the number of attacks used.
	AttackCount int `json:"attacks,omitempty"`
	// AttackLimit is the normal attack limit.
	AttackLimit int `json:"attackLimit,omitempty"`
	// BonusAttackLimit is the number of bonus attacks available.
	BonusAttackLimit int `json:"bonusAttackLimit,omitempty"`
	// CapitalResourcesLooted is the capital gold looted by the member.
	CapitalResourcesLooted int `json:"capitalResourcesLooted,omitempty"`
}

// RaidLogEntry is one Clan Capital raid weekend log entry.
type RaidLogEntry struct {
	// State is the raid weekend state.
	State string `json:"state,omitempty"`
	// TotalLoot is the clan's total capital gold looted.
	TotalLoot int `json:"capitalTotalLoot,omitempty"`
	// CompletedRaidCount is the number of completed raids.
	CompletedRaidCount int `json:"raidsCompleted,omitempty"`
	// AttackCount is the total number of attacks used by the clan.
	AttackCount int `json:"totalAttacks,omitempty"`
	// DestroyedDistrictCount is the number of enemy districts destroyed.
	DestroyedDistrictCount int `json:"enemyDistrictsDestroyed,omitempty"`
	// OffensiveReward is the offensive raid medal reward.
	OffensiveReward int `json:"offensiveReward,omitempty"`
	// DefensiveReward is the defensive raid medal reward.
	DefensiveReward int `json:"defensiveReward,omitempty"`
	// StartTime is when the raid weekend started.
	StartTime *Timestamp `json:"startTime,omitempty"`
	// EndTime is when the raid weekend ended.
	EndTime *Timestamp `json:"endTime,omitempty"`
	// AttackLog contains raids made by the requested clan.
	AttackLog []RaidClan `json:"attackLog,omitempty"`
	// DefenseLog contains raids made against the requested clan.
	DefenseLog []RaidClan `json:"defenseLog,omitempty"`
	// Members contains member-level attack and loot totals.
	Members []RaidMember `json:"members,omitempty"`
	responseMeta
}

// GetMember returns the raid member with the provided tag.
func (r *RaidLogEntry) GetMember(tag string) *RaidMember {
	tag = CorrectTag(tag)
	for i := range r.Members {
		if r.Members[i].Tag == tag {
			return &r.Members[i]
		}
	}
	return nil
}

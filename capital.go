package clashy

type RaidAttack struct {
	AttackerTag  string  `json:"-"`
	AttackerName string  `json:"-"`
	Stars        int     `json:"stars,omitempty"`
	Destruction  float64 `json:"destructionPercent,omitempty"`
}

type RaidDistrict struct {
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	HallLevel   int          `json:"districtHallLevel,omitempty"`
	Destruction float64      `json:"destructionPercent,omitempty"`
	AttackCount int          `json:"attackCount,omitempty"`
	Looted      int          `json:"totalLooted,omitempty"`
	Attacks     []RaidAttack `json:"attacks,omitempty"`
}

type RaidClan struct {
	AttackCount            int            `json:"attackCount,omitempty"`
	DistrictCount          int            `json:"districtCount,omitempty"`
	DestroyedDistrictCount int            `json:"districtsDestroyed,omitempty"`
	Districts              []RaidDistrict `json:"districts,omitempty"`
	Attacker               *struct {
		Tag   string `json:"tag,omitempty"`
		Name  string `json:"name,omitempty"`
		Level int    `json:"level,omitempty"`
	} `json:"attacker,omitempty"`
	Defender *struct {
		Tag   string `json:"tag,omitempty"`
		Name  string `json:"name,omitempty"`
		Level int    `json:"level,omitempty"`
	} `json:"defender,omitempty"`
}

func (r RaidClan) Tag() string {
	if r.Attacker != nil {
		return r.Attacker.Tag
	}
	if r.Defender != nil {
		return r.Defender.Tag
	}
	return ""
}

func (r RaidClan) Name() string {
	if r.Attacker != nil {
		return r.Attacker.Name
	}
	if r.Defender != nil {
		return r.Defender.Name
	}
	return ""
}

func (r RaidClan) Level() int {
	if r.Attacker != nil {
		return r.Attacker.Level
	}
	if r.Defender != nil {
		return r.Defender.Level
	}
	return 0
}

type RaidMember struct {
	Tag                    string `json:"tag,omitempty"`
	Name                   string `json:"name,omitempty"`
	AttackCount            int    `json:"attacks,omitempty"`
	AttackLimit            int    `json:"attackLimit,omitempty"`
	BonusAttackLimit       int    `json:"bonusAttackLimit,omitempty"`
	CapitalResourcesLooted int    `json:"capitalResourcesLooted,omitempty"`
}

type RaidLogEntry struct {
	State                  string       `json:"state,omitempty"`
	TotalLoot              int          `json:"capitalTotalLoot,omitempty"`
	CompletedRaidCount     int          `json:"raidsCompleted,omitempty"`
	AttackCount            int          `json:"totalAttacks,omitempty"`
	DestroyedDistrictCount int          `json:"enemyDistrictsDestroyed,omitempty"`
	OffensiveReward        int          `json:"offensiveReward,omitempty"`
	DefensiveReward        int          `json:"defensiveReward,omitempty"`
	StartTime              *Timestamp   `json:"startTime,omitempty"`
	EndTime                *Timestamp   `json:"endTime,omitempty"`
	AttackLog              []RaidClan   `json:"attackLog,omitempty"`
	DefenseLog             []RaidClan   `json:"defenseLog,omitempty"`
	Members                []RaidMember `json:"members,omitempty"`
	responseMeta
}

func (r *RaidLogEntry) GetMember(tag string) *RaidMember {
	tag = CorrectTag(tag)
	for i := range r.Members {
		if r.Members[i].Tag == tag {
			return &r.Members[i]
		}
	}
	return nil
}

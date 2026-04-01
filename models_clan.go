package clashy

type PlayerClan struct {
	Tag   string `json:"tag,omitempty"`
	Name  string `json:"name,omitempty"`
	Level int    `json:"clanLevel,omitempty"`
	Badge Badge  `json:"badgeUrls,omitempty"`
}

type ClanMember struct {
	Tag                 string      `json:"tag,omitempty"`
	Name                string      `json:"name,omitempty"`
	Role                Role        `json:"role,omitempty"`
	ExpLevel            int         `json:"expLevel,omitempty"`
	Trophies            int         `json:"trophies,omitempty"`
	ClanRank            int         `json:"clanRank,omitempty"`
	ClanPreviousRank    int         `json:"previousClanRank,omitempty"`
	Donations           int         `json:"donations,omitempty"`
	Received            int         `json:"donationsReceived,omitempty"`
	VersusTrophies      int         `json:"versusTrophies,omitempty"`
	BuilderBaseTrophies int         `json:"builderBaseTrophies,omitempty"`
	VersusRank          int         `json:"versusRank,omitempty"`
	BuilderBaseRank     int         `json:"builderBaseRank,omitempty"`
	League              *League     `json:"league,omitempty"`
	Clan                *PlayerClan `json:"clan,omitempty"`
	responseMeta
}

type Clan struct {
	Tag                         string            `json:"tag,omitempty"`
	Name                        string            `json:"name,omitempty"`
	Type                        ClanType          `json:"type,omitempty"`
	Description                 string            `json:"description,omitempty"`
	FamilyFriendly              bool              `json:"isFamilyFriendly,omitempty"`
	Level                       int               `json:"clanLevel,omitempty"`
	Points                      int               `json:"clanPoints,omitempty"`
	BuilderBasePoints           int               `json:"clanBuilderBasePoints,omitempty"`
	CapitalPoints               int               `json:"clanCapitalPoints,omitempty"`
	RequiredTrophies            int               `json:"requiredTrophies,omitempty"`
	WarFrequency                string            `json:"warFrequency,omitempty"`
	WarWinStreak                int               `json:"warWinStreak,omitempty"`
	WarWins                     int               `json:"warWins,omitempty"`
	WarTies                     int               `json:"warTies,omitempty"`
	WarLosses                   int               `json:"warLosses,omitempty"`
	PublicWarLog                bool              `json:"isWarLogPublic,omitempty"`
	MemberCount                 int               `json:"members,omitempty"`
	RequiredBuilderBaseTrophies int               `json:"requiredBuilderBaseTrophies,omitempty"`
	RequiredTownhall            int               `json:"requiredTownhallLevel,omitempty"`
	Location                    *Location         `json:"location,omitempty"`
	Badge                       Badge             `json:"badgeUrls,omitempty"`
	Labels                      []Label           `json:"labels,omitempty"`
	Members                     []ClanMember      `json:"memberList,omitempty"`
	WarLeague                   *BaseLeague       `json:"warLeague,omitempty"`
	CapitalLeague               *BaseLeague       `json:"capitalLeague,omitempty"`
	ChatLanguage                *ChatLanguage     `json:"chatLanguage,omitempty"`
	CapitalDistricts            []CapitalDistrict `json:"-"`
	responseMeta
	RawClanCapital struct {
		Districts []CapitalDistrict `json:"districts,omitempty"`
	} `json:"clanCapital,omitempty"`
}

func (c *Clan) Finalize() {
	c.Badge.finalize()
	c.CapitalDistricts = c.RawClanCapital.Districts
	if c.MemberCount == 0 && len(c.Members) > 0 {
		c.MemberCount = len(c.Members)
	}
	for i := range c.Members {
		if c.Members[i].Clan == nil {
			c.Members[i].Clan = &PlayerClan{Tag: c.Tag, Name: c.Name, Level: c.Level, Badge: c.Badge}
		}
	}
}

func (c *Clan) GetMember(tag string) *ClanMember {
	tag = CorrectTag(tag)
	for i := range c.Members {
		if c.Members[i].Tag == tag {
			return &c.Members[i]
		}
	}
	return nil
}

func (c *Clan) GetMemberBy(name string, trophies int) *ClanMember {
	for i := range c.Members {
		if (name == "" || c.Members[i].Name == name) && (trophies == 0 || c.Members[i].Trophies == trophies) {
			return &c.Members[i]
		}
	}
	return nil
}

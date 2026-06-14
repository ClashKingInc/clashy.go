package clashy

// PlayerClan is the compact clan object embedded in player responses.
type PlayerClan struct {
	// Tag is the clan tag, including the leading # when returned by the API.
	Tag string `json:"tag,omitempty"`
	// Name is the clan display name.
	Name string `json:"name,omitempty"`
	// Level is the clan level.
	Level int `json:"clanLevel,omitempty"`
	// Badge contains the clan badge image URLs.
	Badge Badge `json:"badgeUrls,omitempty"`
}

// ClanMember describes a member entry from a clan profile or member list.
type ClanMember struct {
	// Tag is the member's player tag.
	Tag string `json:"tag,omitempty"`
	// Name is the member's player name.
	Name string `json:"name,omitempty"`
	// Role is the member's clan role.
	Role Role `json:"role,omitempty"`
	// ExpLevel is the player's experience level.
	ExpLevel int `json:"expLevel,omitempty"`
	// TownHall is the player's home village Town Hall level.
	TownHall int `json:"townHallLevel,omitempty"`
	// Trophies is the player's home village trophy count.
	Trophies int `json:"trophies,omitempty"`
	// ClanRank is the player's current position in the clan trophy ranking.
	ClanRank int `json:"clanRank,omitempty"`
	// ClanPreviousRank is the player's previous position in the clan trophy
	// ranking.
	ClanPreviousRank int `json:"previousClanRank,omitempty"`
	// Donations is the number of troops donated this season.
	Donations int `json:"donations,omitempty"`
	// Received is the number of donated troops received this season.
	Received int `json:"donationsReceived,omitempty"`
	// VersusTrophies is the legacy Builder Base trophy field used by older API
	// responses.
	VersusTrophies int `json:"versusTrophies,omitempty"`
	// BuilderBaseTrophies is the player's Builder Base trophy count.
	BuilderBaseTrophies int `json:"builderBaseTrophies,omitempty"`
	// VersusRank is the legacy Builder Base rank field used by older API
	// responses.
	VersusRank int `json:"versusRank,omitempty"`
	// BuilderBaseRank is the player's current Builder Base rank in the clan.
	BuilderBaseRank int `json:"builderBaseRank,omitempty"`
	// LeagueTier is the player's home village league.
	LeagueTier League `json:"leagueTier,omitempty"`
	responseMeta
}

// ClanCapital describes a clan's capital districts from the clan profile.
type ClanCapital struct {
	// Districts contains the visible Clan Capital districts and their hall
	// levels.
	Districts []CapitalDistrict `json:"districts,omitempty"`
}

// Clan is the full clan profile returned by GetClan and search endpoints.
type Clan struct {
	// Tag is the clan tag.
	Tag string `json:"tag,omitempty"`
	// Name is the clan display name.
	Name string `json:"name,omitempty"`
	// Type describes whether the clan is open, closed, or invite-only.
	Type ClanType `json:"type,omitempty"`
	// Description is the public clan description.
	Description string `json:"description,omitempty"`
	// FamilyFriendly reports whether the clan is marked family friendly.
	FamilyFriendly bool `json:"isFamilyFriendly,omitempty"`
	// Level is the clan level.
	Level int `json:"clanLevel,omitempty"`
	// Points is the clan's home village trophy score.
	Points int `json:"clanPoints,omitempty"`
	// BuilderBasePoints is the clan's Builder Base trophy score.
	BuilderBasePoints int `json:"clanBuilderBasePoints,omitempty"`
	// CapitalPoints is the clan's Clan Capital score.
	CapitalPoints int `json:"clanCapitalPoints,omitempty"`
	// RequiredTrophies is the home village trophy requirement to join.
	RequiredTrophies int `json:"requiredTrophies,omitempty"`
	// WarFrequency is the clan's declared war frequency.
	WarFrequency string `json:"warFrequency,omitempty"`
	// WarWinStreak is the current classic-war win streak.
	WarWinStreak int `json:"warWinStreak,omitempty"`
	// WarWins is the number of classic-war wins.
	WarWins int `json:"warWins,omitempty"`
	// WarTies is the number of classic-war ties when the API includes it.
	WarTies int `json:"warTies,omitempty"`
	// WarLosses is the number of classic-war losses when the API includes it.
	WarLosses int `json:"warLosses,omitempty"`
	// PublicWarLog reports whether the clan war log is public.
	PublicWarLog bool `json:"isWarLogPublic,omitempty"`
	// MemberCount is the number of current clan members.
	MemberCount int `json:"members,omitempty"`
	// RequiredBuilderBaseTrophies is the Builder Base trophy requirement to
	// join.
	RequiredBuilderBaseTrophies int `json:"requiredBuilderBaseTrophies,omitempty"`
	// RequiredTownhall is the minimum Town Hall level required to join.
	RequiredTownhall int `json:"requiredTownhallLevel,omitempty"`
	// Location is the clan's declared location.
	Location *Location `json:"location,omitempty"`
	// Badge contains the clan badge image URLs.
	Badge Badge `json:"badgeUrls,omitempty"`
	// Labels are public labels assigned to the clan.
	Labels []Label `json:"labels,omitempty"`
	// Members is the member list embedded in full clan responses.
	Members []ClanMember `json:"memberList,omitempty"`
	// WarLeague is the clan's current Clan War League tier.
	WarLeague League `json:"warLeague,omitempty"`
	// CapitalLeague is the clan's Clan Capital league.
	CapitalLeague *League `json:"capitalLeague,omitempty"`
	// ChatLanguage is the clan's preferred chat language.
	ChatLanguage *ChatLanguage `json:"chatLanguage,omitempty"`
	// ClanCapital contains Clan Capital district information.
	ClanCapital *ClanCapital `json:"clanCapital,omitempty"`
	responseMeta
}

// GetMember returns the member with the provided tag, or nil when the clan
// member list does not contain that tag.
func (c *Clan) GetMember(tag string) *ClanMember {
	tag = CorrectTag(tag)
	for i := range c.Members {
		if c.Members[i].Tag == tag {
			return &c.Members[i]
		}
	}
	return nil
}

// GetMemberBy returns the first member matching the provided name and trophy
// filters.
//
// Empty name and zero trophies are treated as wildcards, which is useful when a
// caller only has one of the two values from an external event.
func (c *Clan) GetMemberBy(name string, trophies int) *ClanMember {
	for i := range c.Members {
		if (name == "" || c.Members[i].Name == name) && (trophies == 0 || c.Members[i].Trophies == trophies) {
			return &c.Members[i]
		}
	}
	return nil
}

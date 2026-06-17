package clashy

// Role is a member's role inside a clan.
type Role string

const (
	// RoleMember is a regular clan member.
	RoleMember Role = "member"
	// RoleElder is a clan elder. The Clash API value is "admin".
	RoleElder Role = "admin"
	// RoleCoLeader is a clan co-leader.
	RoleCoLeader Role = "coLeader"
	// RoleLeader is the clan leader.
	RoleLeader Role = "leader"
)

// WarRound identifies the logical CWL round requested from GetCurrentWar or
// GetLeagueWar.
type WarRound int

const (
	// PreviousWar selects the previous completed or in-war CWL round.
	PreviousWar WarRound = iota
	// CurrentWar selects the active CWL war, or the latest completed/in-war
	// round when the latest real round is only preparation.
	CurrentWar
	// CurrentPreparation selects the upcoming CWL preparation round when one is
	// available.
	CurrentPreparation
)

// WarState is the lifecycle state of a classic war or CWL war.
type WarState string

const (
	// WarStateNotInWar means the clan is not in a regular war.
	WarStateNotInWar WarState = "notInWar"
	// WarStatePreparation means the war is in preparation day.
	WarStatePreparation WarState = "preparation"
	// WarStateInWar means battle day is active.
	WarStateInWar WarState = "inWar"
	// WarStateEnded means the war has ended.
	WarStateEnded WarState = "warEnded"
)

// WarResult is the requested clan's result in a war log entry.
type WarResult string

const (
	// WarResultWin means the requested clan won.
	WarResultWin WarResult = "win"
	// WarResultLose means the requested clan lost.
	WarResultLose WarResult = "lose"
	// WarResultTie means the war ended in a tie.
	WarResultTie WarResult = "tie"
)

// BattleType describes the game mode for a player battle log entry.
type BattleType string

const (
	// BattleTypeHomeVillage is a home-village battle log entry.
	BattleTypeHomeVillage BattleType = "HOME_VILLAGE"
	// BattleTypeRanked is a ranked battle log entry.
	BattleTypeRanked BattleType = "RANKED"
	// BattleTypeLegend is a Legend League battle log entry.
	BattleTypeLegend BattleType = "LEGEND"
)

// BattleModifier describes the modifier applied to a war battle.
type BattleModifier string

const (
	// BattleModifierNone means the war has no battle modifier.
	BattleModifierNone BattleModifier = "none"
	// BattleModifierHardMode is the esports hard mode modifier.
	BattleModifierHardMode BattleModifier = "hardMode"
	// BattleModifierMinusOne is the Legend I battle modifier.
	BattleModifierMinusOne BattleModifier = "minusOne"
	// BattleModifierMinusTwo is the Legend II battle modifier.
	BattleModifierMinusTwo BattleModifier = "minusTwo"
	// BattleModifierMinusThree is the Legend III battle modifier.
	BattleModifierMinusThree BattleModifier = "minusThree"
)

// InGameName returns a client-facing display name for the battle modifier.
func (m BattleModifier) InGameName() string {
	switch m {
	case "", BattleModifierNone:
		return "None"
	case BattleModifierHardMode:
		return "Hard Mode"
	case BattleModifierMinusOne:
		return "Minus One"
	case BattleModifierMinusTwo:
		return "Minus Two"
	case BattleModifierMinusThree:
		return "Minus Three"
	default:
		return string(m)
	}
}

// ClanType describes a clan's join policy.
type ClanType string

const (
	// ClanTypeOpen means players can join directly when requirements are met.
	ClanTypeOpen ClanType = "open"
	// ClanTypeClosed means the clan is closed to new members.
	ClanTypeClosed ClanType = "closed"
	// ClanTypeInviteOnly means players must request or be invited to join.
	ClanTypeInviteOnly ClanType = "inviteOnly"
)

// VillageType identifies the village or game area for static data and units.
type VillageType string

const (
	// VillageHome is the home village.
	VillageHome VillageType = "home"
	// VillageBuilderBase is Builder Base.
	VillageBuilderBase VillageType = "builderBase"
	// VillageClanCapital is Clan Capital.
	VillageClanCapital VillageType = "clanCapital"
)

// LoadGameData describes when static game data should be loaded.
type LoadGameData struct {
	// Default uses the package's normal embedded static-data behavior.
	Default bool
	// StartupOnly indicates static data should be loaded during client
	// construction only.
	StartupOnly bool
	// Always indicates static data should be refreshed whenever supported by the
	// caller's workflow.
	Always bool
	// Never indicates static data should not be loaded.
	Never bool
}

// DefaultLoadGameData returns the default static-data loading policy.
func DefaultLoadGameData() LoadGameData {
	return LoadGameData{Default: true}
}

package clashy

type Role string

const (
	RoleMember   Role = "member"
	RoleElder    Role = "admin"
	RoleCoLeader Role = "coLeader"
	RoleLeader   Role = "leader"
)

type WarRound int

const (
	PreviousWar WarRound = iota
	CurrentWar
	CurrentPreparation
)

type WarState string

const (
	WarStateNotInWar    WarState = "notInWar"
	WarStatePreparation WarState = "preparation"
	WarStateInWar       WarState = "inWar"
	WarStateEnded       WarState = "warEnded"
)

type WarResult string

const (
	WarResultWin  WarResult = "win"
	WarResultLose WarResult = "lose"
	WarResultTie  WarResult = "tie"
)

type ClanType string

const (
	ClanTypeOpen       ClanType = "open"
	ClanTypeClosed     ClanType = "closed"
	ClanTypeInviteOnly ClanType = "inviteOnly"
)

type VillageType string

const (
	VillageHome        VillageType = "home"
	VillageBuilderBase VillageType = "builderBase"
	VillageClanCapital VillageType = "clanCapital"
)

type LoadGameData struct {
	Default     bool
	StartupOnly bool
	Always      bool
	Never       bool
}

func DefaultLoadGameData() LoadGameData {
	return LoadGameData{Default: true}
}

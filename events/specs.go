package events

import (
	"context"
	"reflect"

	clashy "github.com/clashkinginc/clashy.go"
)

func ClanMemberJoined(fn Handler[clashy.Clan]) Spec[clashy.Clan] {
	return Custom("clan_member_joined", func(oldValue, newValue *clashy.Clan) bool {
		if oldValue == nil || newValue == nil {
			return false
		}
		return len(newValue.Members) > len(oldValue.Members)
	}, fn)
}

func ClanMemberLeft(fn Handler[clashy.Clan]) Spec[clashy.Clan] {
	return Custom("clan_member_left", func(oldValue, newValue *clashy.Clan) bool {
		if oldValue == nil || newValue == nil {
			return false
		}
		return len(newValue.Members) < len(oldValue.Members)
	}, fn)
}

func PlayerAchievementChanged(fn Handler[clashy.Player]) Spec[clashy.Player] {
	return Custom("player_achievement_changed", func(oldValue, newValue *clashy.Player) bool {
		if oldValue == nil || newValue == nil {
			return false
		}
		return !reflect.DeepEqual(oldValue.Achievements, newValue.Achievements)
	}, fn)
}

func WarStateChanged(fn func(context.Context, FieldChange[clashy.ClanWar, clashy.WarState]) error) Spec[clashy.ClanWar] {
	return FieldChanged("war_state_changed", func(war clashy.ClanWar) clashy.WarState {
		return war.State
	}, fn)
}

func WarAttackAdded(fn Handler[clashy.ClanWar]) Spec[clashy.ClanWar] {
	return Custom("war_attack_added", func(oldValue, newValue *clashy.ClanWar) bool {
		if oldValue == nil || newValue == nil {
			return false
		}
		return len(newValue.Attacks()) > len(oldValue.Attacks())
	}, fn)
}

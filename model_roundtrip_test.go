package clashy

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestTimestampJSONRoundTripUsesClashWireFormat(t *testing.T) {
	const raw = `"20260616T120102.000Z"`
	var timestamp Timestamp
	if err := json.Unmarshal([]byte(raw), &timestamp); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	encoded, err := json.Marshal(timestamp)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(encoded) != raw {
		t.Fatalf("encoded timestamp = %s, want %s", encoded, raw)
	}

	fromTime := Timestamp{Time: time.Date(2026, 6, 16, 12, 1, 2, 0, time.UTC)}
	encoded, err = json.Marshal(fromTime)
	if err != nil {
		t.Fatalf("Marshal from Time: %v", err)
	}
	if string(encoded) != raw {
		t.Fatalf("encoded Time = %s, want %s", encoded, raw)
	}
}

func TestPlayerUnitsDoNotExposeStaticOnlyFields(t *testing.T) {
	var troop Troop
	if err := json.Unmarshal([]byte(`{"name":"Barbarian","level":12,"maxLevel":13,"village":"home"}`), &troop); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	encoded, err := json.Marshal(troop)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	for _, field := range [][]byte{[]byte("upgradeCost"), []byte("upgradeTime")} {
		if bytes.Contains(encoded, field) {
			t.Fatalf("player troop unexpectedly contains static field %q: %s", field, encoded)
		}
	}
}

func TestResolveAttackUsesWarMemberTags(t *testing.T) {
	war := &ClanWar{
		Clan:     &WarClan{Members: []ClanWarMember{{Tag: "#ATT", Name: "Attacker"}}},
		Opponent: &WarClan{Members: []ClanWarMember{{Tag: "#DEF", Name: "Defender"}}},
	}
	attacker, defender := war.ResolveAttack(WarAttack{AttackerTag: "att", DefenderTag: "def"})
	if attacker == nil || attacker.Name != "Attacker" {
		t.Fatalf("attacker = %#v", attacker)
	}
	if defender == nil || defender.Name != "Defender" {
		t.Fatalf("defender = %#v", defender)
	}
}

func TestClanWarLeagueClanDecodesMasterRoster(t *testing.T) {
	var clan ClanWarLeagueClan
	if err := json.Unmarshal([]byte(`{"tag":"#AAA","members":[{"tag":"#P1","name":"One","townHallLevel":17}]}`), &clan); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(clan.Members) != 1 || clan.Members[0].Tag != "#P1" || clan.Members[0].TownHallLevel != 17 {
		t.Fatalf("members = %#v", clan.Members)
	}
}

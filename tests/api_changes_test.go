package clashy_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	clashy "github.com/clashkinginc/clashy.go"
)

func TestBattleLogEntryDecodesSwaggerFields(t *testing.T) {
	t.Parallel()

	client := newHTTPTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.EscapedPath() != "/players/%232PP/battlelog" {
			t.Fatalf("unexpected battlelog path: %s", r.URL.EscapedPath())
		}
		writeJSON(t, w, `{"items":[{"battleType":"LEGEND","attack":true,"armyShareCode":"u1x1","opponentPlayerTag":"#ABC","opponentName":"Defender","opponentTownHallLevel":17,"stars":3,"destructionPercentage":100,"lootedResources":[{"name":"gold","amount":1234}],"extraLootedResources":[{"name":"elixir","amount":50}],"availableLoot":[{"name":"darkElixir","amount":9}],"battleTime":171,"battleTimestamp":"20260616T120102.000Z"}]}`)
	}))

	entries, err := client.GetBattleLog(context.Background(), "#2PP")
	if err != nil {
		t.Fatalf("get battle log: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 battle log entry, got %d", len(entries))
	}
	entry := entries[0]
	if entry.BattleType != clashy.BattleTypeLegend {
		t.Fatalf("battle type = %q, want %q", entry.BattleType, clashy.BattleTypeLegend)
	}
	if !entry.Attack || entry.OpponentPlayerTag != "#ABC" || entry.OpponentName != "Defender" || entry.OpponentTownHallLevel != 17 {
		t.Fatalf("unexpected opponent fields: %#v", entry)
	}
	if entry.Duration != 171 || entry.Timestamp != "20260616T120102.000Z" {
		t.Fatalf("unexpected battle timing fields: %#v", entry)
	}
	if len(entry.LootedResources) != 1 || entry.LootedResources[0].Amount != 1234 {
		t.Fatalf("unexpected looted resources: %#v", entry.LootedResources)
	}
}

func TestLeagueGroupUsesStringSeasonIDAndEscapedTags(t *testing.T) {
	t.Parallel()

	client := newHTTPTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.EscapedPath() != "/leaguegroup/%23LG123/2026-06-02" {
			t.Fatalf("unexpected league group path: %s", r.URL.EscapedPath())
		}
		if got := r.URL.Query().Get("playerTag"); got != "#2PP" {
			t.Fatalf("playerTag query = %q, want #2PP", got)
		}
		writeJSON(t, w, `{"members":[{"playerTag":"#2PP","playerName":"Player","clanTag":"#AAA","clanName":"Clan","leagueTrophies":6000,"attackWinCount":1,"attackLoseCount":2,"defenseWinCount":3,"defenseLoseCount":4}],"attackLogs":[{"opponentPlayerTag":"#DEF","opponentName":"Defender","stars":2,"destructionPercentage":80,"trophies":20,"creationTime":"20260616T120102.000Z"}],"defenseLogs":[{"opponentPlayerTag":"#ATK","opponentName":"Attacker","stars":1,"destructionPercentage":50,"trophies":-10,"creationTime":"20260616T130102.000Z"}]}`)
	}))

	group, err := client.GetPlayerLeagueGroup(context.Background(), "#2PP", "#LG123", "2026-06-02")
	if err != nil {
		t.Fatalf("get player league group: %v", err)
	}
	if len(group.Members) != 1 || group.Members[0].DefenseLoseCount != 4 {
		t.Fatalf("unexpected group members: %#v", group.Members)
	}
	if len(group.AttackLogs) != 1 || group.AttackLogs[0].OpponentName != "Defender" {
		t.Fatalf("unexpected attack logs: %#v", group.AttackLogs)
	}
}

func TestRankingAndWarPatchFieldsDecode(t *testing.T) {
	t.Parallel()

	var ranking clashy.RankedPlayer
	if err := json.Unmarshal([]byte(`{"tag":"#2PP","name":"Player","league":{"id":29000022,"name":"Legend League"},"leagueTier":{"id":29000023,"name":"Legend I"},"attackWins":8,"defenseWins":5,"rank":1,"previousRank":2,"trophies":6123}`), &ranking); err != nil {
		t.Fatalf("unmarshal ranked player: %v", err)
	}
	if ranking.League.Name != "Legend League" || ranking.LeagueTier.Name != "Legend I" {
		t.Fatalf("unexpected ranking leagues: %#v", ranking)
	}
	if ranking.AttackWins != 8 || ranking.DefenseWins != 5 {
		t.Fatalf("unexpected ranking wins: %#v", ranking)
	}

	var war clashy.ClanWar
	if err := json.Unmarshal([]byte(`{"clan":{"members":[{"tag":"#AAA","attacks":[{"attackerTag":"#AAA","defenderTag":"#BBB","stars":3,"destructionPercentage":100,"duration":179}]}]}}`), &war); err != nil {
		t.Fatalf("unmarshal clan war: %v", err)
	}
	if war.Clan == nil || len(war.Clan.Members) != 1 || len(war.Clan.Members[0].Attacks) != 1 {
		t.Fatalf("unexpected war payload: %#v", war)
	}
	if got := war.Clan.Members[0].Attacks[0].Duration; got != 179 {
		t.Fatalf("duration = %d, want 179", got)
	}
}

func TestBattleModifierDecodesAndFormats(t *testing.T) {
	t.Parallel()

	cases := map[clashy.BattleModifier]string{
		clashy.BattleModifierNone:       "None",
		clashy.BattleModifierHardMode:   "Hard Mode",
		clashy.BattleModifierMinusOne:   "Minus One",
		clashy.BattleModifierMinusTwo:   "Minus Two",
		clashy.BattleModifierMinusThree: "Minus Three",
		"":                              "None",
	}
	for modifier, want := range cases {
		if got := modifier.InGameName(); got != want {
			t.Fatalf("%q InGameName = %q, want %q", modifier, got, want)
		}
	}

	var war clashy.ClanWar
	if err := json.Unmarshal([]byte(`{"battleModifier":"hardMode"}`), &war); err != nil {
		t.Fatalf("unmarshal clan war battle modifier: %v", err)
	}
	if war.BattleModifier != clashy.BattleModifierHardMode {
		t.Fatalf("war battle modifier = %q, want %q", war.BattleModifier, clashy.BattleModifierHardMode)
	}

	var entry clashy.ClanWarLogEntry
	if err := json.Unmarshal([]byte(`{"battleModifier":"minusThree"}`), &entry); err != nil {
		t.Fatalf("unmarshal war log battle modifier: %v", err)
	}
	if entry.BattleModifier != clashy.BattleModifierMinusThree {
		t.Fatalf("war log battle modifier = %q, want %q", entry.BattleModifier, clashy.BattleModifierMinusThree)
	}
}

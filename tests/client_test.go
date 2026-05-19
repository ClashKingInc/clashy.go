package clashy_test

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	clashy "github.com/clashkinginc/clashy.go"
)

const defaultMockAPIBaseURL = "https://api.clashapi.dev"

func newMockAPIClient(t *testing.T) *clashy.Client {
	t.Helper()

	cfg := clashy.DefaultClientConfig()
	cfg.BaseURL = mockAPIBaseURL()
	cfg.Timeout = 15 * time.Second
	cfg.LookupCache = false
	cfg.UpdateCache = false

	client, err := clashy.NewClient(cfg)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	if err := client.LoginWithTokens(context.Background(), "mock-token"); err != nil {
		t.Fatalf("login with tokens: %v", err)
	}
	return client
}

func mockAPIBaseURL() string {
	if baseURL := strings.TrimSpace(os.Getenv("CLASHY_TEST_BASE_URL")); baseURL != "" {
		return baseURL
	}
	return defaultMockAPIBaseURL
}

func testContext(t *testing.T) (context.Context, context.CancelFunc) {
	t.Helper()
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func TestMockAPIClanEndpoints(t *testing.T) {
	client := newMockAPIClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	clan, err := client.GetClan(ctx, "2PP")
	if err != nil {
		t.Fatalf("get clan: %v", err)
	}
	if clan.Tag != "#2PP" || clan.Name == "" {
		t.Fatalf("unexpected clan: %#v", clan)
	}
	if clan.Badge.URL() == "" {
		t.Fatalf("expected finalized badge url")
	}
	if clan.MemberCount != len(clan.Members) {
		t.Fatalf("member count mismatch: %d vs %d", clan.MemberCount, len(clan.Members))
	}
	if len(clan.Members) == 0 {
		t.Fatalf("expected clan members in clan payload")
	}
	if member := clan.GetMember(clan.Members[0].Tag); member == nil || member.Name == "" {
		t.Fatalf("expected member lookup to work: %#v", member)
	}

	found, err := client.SearchClans(ctx, clashy.SearchClansRequest{Name: "Order", Limit: 3})
	if err != nil {
		t.Fatalf("search clans: %v", err)
	}
	if len(found) == 0 || found[0].Tag == "" || found[0].Name == "" {
		t.Fatalf("unexpected clan search results: %#v", found)
	}

	members, err := client.GetMembers(ctx, "#2PP", 5, "", "")
	if err != nil {
		t.Fatalf("get members: %v", err)
	}
	if len(members) != 5 {
		t.Fatalf("expected 5 members, got %d", len(members))
	}
	if members[0].Tag == "" || members[0].Name == "" {
		t.Fatalf("unexpected member payload: %#v", members[0])
	}

	warLog, err := client.GetWarLog(ctx, "#2PP", 2, "", "")
	if err != nil {
		t.Fatalf("get war log: %v", err)
	}
	if len(warLog) != 2 {
		t.Fatalf("expected 2 war log entries, got %d", len(warLog))
	}
	if warLog[0].Clan == nil || warLog[0].Opponent == nil {
		t.Fatalf("expected war log entries to contain both sides: %#v", warLog[0])
	}

	var privateWarLog *clashy.PrivateWarLog
	var forbidden *clashy.Forbidden
	if _, err := client.GetWarLog(ctx, "#2PPP", 0, "", ""); !errors.As(err, &privateWarLog) && !errors.As(err, &forbidden) {
		t.Fatalf("expected private war log error, got %v", err)
	}
}

func TestMockAPIPlayerEndpoints(t *testing.T) {
	client := newMockAPIClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	player, err := client.GetPlayer(ctx, "#2PP")
	if err != nil {
		t.Fatalf("get player: %v", err)
	}
	if player.Tag != "#2PP" || player.Name == "" {
		t.Fatalf("unexpected player: %#v", player)
	}
	if player.Clan == nil || player.Clan.Tag == "" {
		t.Fatalf("expected player clan data: %#v", player.Clan)
	}
	if player.LeagueTier.Name == "" {
		t.Fatalf("expected league tier data: %#v", player.LeagueTier)
	}
	if achievement := player.GetAchievement("Bigger Coffers"); achievement == nil {
		t.Fatalf("expected achievement lookup")
	}
	if troop := player.GetTroop("Barbarian"); troop == nil {
		t.Fatalf("expected troop lookup")
	} else {
		if staticTroop := troop.Static(client); staticTroop == nil || staticTroop.MaxLevel == 0 {
			t.Fatalf("expected explicit static troop lookup: %#v", staticTroop)
		}
	}
	if player.CurrentLeagueGroupTag == "" || player.CurrentLeagueSeasonID == 0 {
		t.Fatalf("expected current league group metadata on player: %#v", player)
	}
	if player.PreviousLeagueGroupTag == "" || player.PreviousLeagueSeasonID == 0 {
		t.Fatalf("expected previous league group metadata on player: %#v", player)
	}

	battleLog, err := client.GetBattleLog(ctx, "#2PP")
	if err != nil {
		t.Fatalf("get battle log: %v", err)
	}
	if len(battleLog) == 0 || battleLog[0].OpponentPlayerTag == "" || len(battleLog[0].LootedResources) == 0 {
		t.Fatalf("unexpected battle log: %#v", battleLog)
	}

	history, err := client.GetPlayerLeagueHistory(ctx, "#2PP")
	if err != nil {
		t.Fatalf("get player league history: %v", err)
	}
	if len(history) == 0 || history[0].LeagueSeasonID == 0 || history[0].LeagueTierID == 0 {
		t.Fatalf("unexpected player league history: %#v", history)
	}

	currentGroup, err := client.GetPlayerLeagueGroup(ctx, player.Tag, player.CurrentLeagueGroupTag, player.CurrentLeagueSeasonID)
	if err != nil {
		t.Fatalf("get current player league group: %v", err)
	}
	if len(currentGroup.Members) == 0 || len(currentGroup.AttackLogs) == 0 || len(currentGroup.DefenseLogs) == 0 {
		t.Fatalf("unexpected current player league group: %#v", currentGroup)
	}

	previousGroup, err := client.GetPlayerLeagueGroup(ctx, player.Tag, player.PreviousLeagueGroupTag, player.PreviousLeagueSeasonID)
	if err != nil {
		t.Fatalf("get previous player league group: %v", err)
	}
	if len(previousGroup.Members) == 0 || len(previousGroup.AttackLogs) == 0 || len(previousGroup.DefenseLogs) == 0 {
		t.Fatalf("unexpected previous player league group: %#v", previousGroup)
	}
	foundSelf := false
	for _, member := range previousGroup.Members {
		if member.PlayerTag == player.Tag {
			foundSelf = true
			break
		}
	}
	if !foundSelf {
		t.Fatalf("expected player %q in previous league group", player.Tag)
	}

	valid, err := client.VerifyPlayerToken(ctx, "#2PP", "TOKEN")
	if err != nil {
		t.Fatalf("verify player token: %v", err)
	}
	if !valid {
		t.Fatalf("expected valid token")
	}

	invalid, err := client.VerifyPlayerToken(ctx, "#2PP", "WRONG")
	if err != nil {
		t.Fatalf("verify invalid token: %v", err)
	}
	if invalid {
		t.Fatalf("expected invalid token to return false")
	}

	var notFound *clashy.NotFound
	if _, err := client.VerifyPlayerToken(ctx, "#2PPP", "TOKEN"); !errors.As(err, &notFound) {
		t.Fatalf("expected not found from verify token variant, got %v", err)
	}
}

func TestMockAPIWarAndRaidEndpoints(t *testing.T) {
	client := newMockAPIClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	war, err := client.GetCurrentWar(ctx, "#2PP")
	if err != nil {
		t.Fatalf("get current war: %v", err)
	}
	if war.State == "" || war.TeamSize == 0 || war.Clan == nil || war.Opponent == nil {
		t.Fatalf("unexpected war payload: %#v", war)
	}
	if len(war.Attacks()) == 0 {
		t.Fatalf("expected war attacks")
	}

	group, err := client.GetLeagueGroup(ctx, "#2PP")
	if err != nil {
		t.Fatalf("get league group: %v", err)
	}
	if group.State == "" || len(group.Rounds) == 0 || len(group.Clans) == 0 {
		t.Fatalf("unexpected league group payload: %#v", group)
	}

	raids, err := client.GetRaidLog(ctx, "#2PP", 2, "", "")
	if err != nil {
		t.Fatalf("get raid log: %v", err)
	}
	if len(raids) != 2 {
		t.Fatalf("expected 2 raid log entries, got %d", len(raids))
	}
	if len(raids[0].Members) == 0 || len(raids[0].AttackLog) == 0 {
		t.Fatalf("expected raid members and attack log: %#v", raids[0])
	}
	if member := raids[0].GetMember(raids[0].Members[0].Tag); member == nil {
		t.Fatalf("expected raid member lookup")
	}
}

func TestMockAPIMetadataAndLeagueEndpoints(t *testing.T) {
	client := newMockAPIClient(t)
	ctx, cancel := testContext(t)
	defer cancel()

	location, err := client.GetLocationNamed(ctx, "International")
	if err != nil {
		t.Fatalf("get location named: %v", err)
	}
	if location == nil || location.Name != "International" {
		t.Fatalf("unexpected location: %#v", location)
	}

	clanLabels, err := client.GetClanLabels(ctx, 3, "", "")
	if err != nil {
		t.Fatalf("get clan labels: %v", err)
	}
	if len(clanLabels) != 3 || clanLabels[0].Name == "" {
		t.Fatalf("unexpected clan labels: %#v", clanLabels)
	}

	playerLabels, err := client.GetPlayerLabels(ctx, 3, "", "")
	if err != nil {
		t.Fatalf("get player labels: %v", err)
	}
	if len(playerLabels) != 3 || playerLabels[0].Name == "" {
		t.Fatalf("unexpected player labels: %#v", playerLabels)
	}

	rankings, err := client.GetLocationPlayers(ctx, 32000249, 3, "", "")
	if err != nil {
		t.Fatalf("get location players: %v", err)
	}
	if len(rankings) != 3 || rankings[0].Tag == "" {
		t.Fatalf("unexpected player rankings: %#v", rankings)
	}

	leagues, err := client.SearchLeagues(ctx, 3, "", "")
	if err != nil {
		t.Fatalf("search leagues: %v", err)
	}
	if len(leagues) != 3 || leagues[0].ID == 0 || leagues[0].Name == "" {
		t.Fatalf("unexpected leagues: %#v", leagues)
	}

	league, err := client.GetLeague(ctx, leagues[0].ID)
	if err != nil {
		t.Fatalf("get league: %v", err)
	}
	if league.ID != leagues[0].ID || league.Name == "" {
		t.Fatalf("unexpected league: %#v", league)
	}

	seasons, err := client.GetSeasons(ctx, 29000022)
	if err != nil {
		t.Fatalf("get seasons: %v", err)
	}
	if len(seasons) == 0 {
		t.Fatalf("expected league seasons")
	}

	seasonRankings, err := client.GetSeasonRankings(ctx, 29000022, seasons[0])
	if err != nil {
		t.Fatalf("get season rankings: %v", err)
	}
	if len(seasonRankings) == 0 || seasonRankings[0].Tag == "" {
		t.Fatalf("unexpected season rankings: %#v", seasonRankings)
	}

	goldPass, err := client.GetCurrentGoldPassSeason(ctx)
	if err != nil {
		t.Fatalf("get current gold pass season: %v", err)
	}
	if goldPass.StartTime == nil || goldPass.EndTime == nil {
		t.Fatalf("expected gold pass season window: %#v", goldPass)
	}
}

func TestStaticHelpers(t *testing.T) {
	client, err := clashy.NewClient(clashy.DefaultClientConfig())
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	troop := client.GetTroop("Barbarian", true, 12)
	if troop == nil || troop.Name != "Barbarian" || troop.Level != 12 || troop.MaxLevel == 0 {
		t.Fatalf("unexpected troop: %#v", troop)
	}

	spell := client.GetSpell("Lightning Spell", 3)
	if spell == nil || spell.Name != "Lightning Spell" {
		t.Fatalf("unexpected spell: %#v", spell)
	}

	translation := client.GetTranslation("TID_BARBARIAN")
	if translation == nil || !strings.Contains(strings.ToLower(translation.English), "barbarian") {
		t.Fatalf("unexpected translation: %#v", translation)
	}

	if got := client.GetExtendedCWLGroupData("Champion League I"); got == nil || got.Name == "" {
		t.Fatalf("expected extended cwl group data")
	}
}

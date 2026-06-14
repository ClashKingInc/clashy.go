package clashy_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	clashy "github.com/clashkinginc/clashy.go"
)

func newHTTPTestClient(t *testing.T, handler http.Handler) *clashy.Client {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	cfg := clashy.DefaultClientConfig()
	cfg.BaseURL = server.URL
	cfg.LookupCache = false
	cfg.UpdateCache = false
	client, err := clashy.NewClient(cfg)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	if err := client.LoginWithTokens(context.Background(), "test-token"); err != nil {
		t.Fatalf("login with token: %v", err)
	}
	return client
}

func writeJSON(t *testing.T, w http.ResponseWriter, body string) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(body))
}

func TestGetCurrentWarReturnsRegularWarBeforeSearchingCWL(t *testing.T) {
	client := newHTTPTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/leaguegroup") || strings.Contains(r.URL.Path, "/clanwarleagues/wars/") {
			t.Fatalf("GetCurrentWar should not search CWL when regular war is active: %s", r.URL.Path)
		}
		if strings.HasSuffix(r.URL.Path, "/currentwar") {
			writeJSON(t, w, `{"state":"inWar","teamSize":15,"clan":{"tag":"#AAA"},"opponent":{"tag":"#BBB"}}`)
			return
		}
		http.NotFound(w, r)
	}))

	war, err := client.GetCurrentWar(context.Background(), "#AAA")
	if err != nil {
		t.Fatalf("get current war: %v", err)
	}
	if war == nil || war.Type() != "random" || war.Clan == nil || war.Clan.Tag != "#AAA" {
		t.Fatalf("unexpected regular war: %#v", war)
	}
}

func TestGetCurrentWarFallsBackToCWLAndOrientsClan(t *testing.T) {
	client := newHTTPTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/currentwar"):
			writeJSON(t, w, `{"state":"notInWar"}`)
		case strings.HasSuffix(r.URL.Path, "/currentwar/leaguegroup"):
			writeJSON(t, w, `{"state":"ended","season":"2026-05","clans":[{"tag":"#AAA"},{"tag":"#BBB"}],"rounds":[{"warTags":["#WAR1"]}]}`)
		case strings.Contains(r.URL.Path, "/clanwarleagues/wars/"):
			writeJSON(t, w, `{"state":"warEnded","teamSize":15,"clan":{"tag":"#BBB"},"opponent":{"tag":"#AAA"}}`)
		default:
			http.NotFound(w, r)
		}
	}))

	war, err := client.GetCurrentWar(context.Background(), "#AAA")
	if err != nil {
		t.Fatalf("get current war: %v", err)
	}
	if war == nil {
		t.Fatal("expected cwl war")
	}
	if war.Type() != "cwl" || war.WarTag != "#WAR1" || war.LeagueGroup == nil {
		t.Fatalf("expected cwl metadata on war: %#v", war)
	}
	if war.Clan == nil || war.Clan.Tag != "#AAA" {
		t.Fatalf("expected target clan to be oriented as clan side: %#v", war)
	}
}

func TestGetCurrentWarReturnsFirstRoundPreparationWar(t *testing.T) {
	client := newHTTPTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/currentwar"):
			writeJSON(t, w, `{"state":"notInWar"}`)
		case strings.HasSuffix(r.URL.Path, "/currentwar/leaguegroup"):
			writeJSON(t, w, `{"state":"preparation","season":"2026-05","rounds":[{"warTags":["#WAR1"]},{"warTags":["#0"]}]}`)
		case strings.Contains(r.URL.Path, "/clanwarleagues/wars/"):
			writeJSON(t, w, `{"state":"preparation","teamSize":15,"clan":{"tag":"#AAA"},"opponent":{"tag":"#BBB"}}`)
		default:
			http.NotFound(w, r)
		}
	}))

	war, err := client.GetCurrentWar(context.Background(), "#AAA")
	if err != nil {
		t.Fatalf("get current war: %v", err)
	}
	if war == nil || war.WarTag != "#WAR1" || war.State != clashy.WarStatePreparation {
		t.Fatalf("expected first round preparation war, got %#v", war)
	}
}

func TestGetCurrentWarUsesPreviousRoundWhenLatestRoundStillPreparing(t *testing.T) {
	client := newHTTPTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/currentwar"):
			writeJSON(t, w, `{"state":"notInWar"}`)
		case strings.HasSuffix(r.URL.Path, "/currentwar/leaguegroup"):
			writeJSON(t, w, `{"state":"inWar","season":"2026-05","rounds":[{"warTags":["#WAR1"]},{"warTags":["#WAR2"]}]}`)
		case strings.Contains(r.URL.Path, "/clanwarleagues/wars/") && strings.Contains(r.URL.Path, "WAR2"):
			writeJSON(t, w, `{"state":"preparation","teamSize":15,"clan":{"tag":"#AAA"},"opponent":{"tag":"#CCC"}}`)
		case strings.Contains(r.URL.Path, "/clanwarleagues/wars/") && strings.Contains(r.URL.Path, "WAR1"):
			writeJSON(t, w, `{"state":"warEnded","teamSize":15,"clan":{"tag":"#AAA"},"opponent":{"tag":"#BBB"}}`)
		default:
			http.NotFound(w, r)
		}
	}))

	war, err := client.GetCurrentWar(context.Background(), "#AAA")
	if err != nil {
		t.Fatalf("get current war: %v", err)
	}
	if war == nil || war.WarTag != "#WAR1" {
		t.Fatalf("expected previous round war while latest round is preparing, got %#v", war)
	}
}

func TestGetCurrentWarCurrentPreparationReturnsLatestPreparingRound(t *testing.T) {
	client := newHTTPTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/currentwar"):
			writeJSON(t, w, `{"state":"notInWar"}`)
		case strings.HasSuffix(r.URL.Path, "/currentwar/leaguegroup"):
			writeJSON(t, w, `{"state":"inWar","season":"2026-05","rounds":[{"warTags":["#WAR1"]},{"warTags":["#WAR2"]}]}`)
		case strings.Contains(r.URL.Path, "/clanwarleagues/wars/") && strings.Contains(r.URL.Path, "WAR2"):
			writeJSON(t, w, `{"state":"preparation","teamSize":15,"clan":{"tag":"#AAA"},"opponent":{"tag":"#CCC"}}`)
		case strings.Contains(r.URL.Path, "/clanwarleagues/wars/") && strings.Contains(r.URL.Path, "WAR1"):
			writeJSON(t, w, `{"state":"warEnded","teamSize":15,"clan":{"tag":"#AAA"},"opponent":{"tag":"#BBB"}}`)
		default:
			http.NotFound(w, r)
		}
	}))

	war, err := client.GetCurrentWar(context.Background(), "#AAA", clashy.CurrentPreparation)
	if err != nil {
		t.Fatalf("get current preparation: %v", err)
	}
	if war == nil || war.WarTag != "#WAR2" || war.State != clashy.WarStatePreparation {
		t.Fatalf("expected latest preparation round, got %#v", war)
	}
}

func TestGetCurrentWarFinalDayReturnsLastRound(t *testing.T) {
	client := newHTTPTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/currentwar"):
			writeJSON(t, w, `{"state":"notInWar"}`)
		case strings.HasSuffix(r.URL.Path, "/currentwar/leaguegroup"):
			writeJSON(t, w, `{"state":"ended","season":"2026-05","rounds":[{"warTags":["#WAR1"]},{"warTags":["#WAR2"]},{"warTags":["#WAR3"]}]}`)
		case strings.Contains(r.URL.Path, "/clanwarleagues/wars/") && strings.Contains(r.URL.Path, "WAR3"):
			writeJSON(t, w, `{"state":"warEnded","teamSize":15,"clan":{"tag":"#AAA"},"opponent":{"tag":"#DDD"}}`)
		case strings.Contains(r.URL.Path, "/clanwarleagues/wars/") && strings.Contains(r.URL.Path, "WAR2"):
			writeJSON(t, w, `{"state":"warEnded","teamSize":15,"clan":{"tag":"#AAA"},"opponent":{"tag":"#CCC"}}`)
		default:
			http.NotFound(w, r)
		}
	}))

	war, err := client.GetCurrentWar(context.Background(), "#AAA")
	if err != nil {
		t.Fatalf("get current war: %v", err)
	}
	if war == nil || war.WarTag != "#WAR3" {
		t.Fatalf("expected final round as current war, got %#v", war)
	}
	previous, err := client.GetCurrentWar(context.Background(), "#AAA", clashy.PreviousWar)
	if err != nil {
		t.Fatalf("get previous war: %v", err)
	}
	if previous == nil || previous.WarTag != "#WAR2" {
		t.Fatalf("expected previous round before final round, got %#v", previous)
	}
}

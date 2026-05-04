package clashy_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	clashy "github.com/clashkinginc/clashy.go"
)

func TestDirectReturnedModelsCarryResponseMeta(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "public, max-age=17")

		switch r.URL.Path {
		case "/locations":
			_, _ = w.Write([]byte(`{"items":[{"id":32000007,"name":"International","isCountry":false,"countryCode":"","localizedName":"International"}]}`))
		case "/locations/32000007":
			_, _ = w.Write([]byte(`{"id":32000007,"name":"International","isCountry":false,"countryCode":"","localizedName":"International"}`))
		default:
			if strings.Contains(r.URL.Path, "/clans/") && strings.HasSuffix(r.URL.Path, "/warlog") {
				_, _ = w.Write([]byte(`{"items":[{"result":"win","endTime":"20260101T000000.000Z","teamSize":15}]}`))
				return
			}
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	cfg := clashy.DefaultClientConfig()
	cfg.BaseURL = server.URL
	cfg.LookupCache = false
	cfg.UpdateCache = false
	cfg.Timeout = 5 * time.Second

	client, err := clashy.NewClient(cfg)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	if err := client.LoginWithTokens(context.Background(), "token"); err != nil {
		t.Fatalf("login with tokens: %v", err)
	}

	locations, err := client.SearchLocations(context.Background(), 0, "", "")
	if err != nil {
		t.Fatalf("search locations: %v", err)
	}
	if len(locations) != 1 {
		t.Fatalf("expected one location, got %d", len(locations))
	}
	if locations[0].RetryAfter() != 0 {
		t.Fatalf("expected search location item retry-after to stay unset, got %d", locations[0].RetryAfter())
	}

	location, err := client.GetLocation(context.Background(), 32000007)
	if err != nil {
		t.Fatalf("get location: %v", err)
	}
	if location.RetryAfter() != 17 {
		t.Fatalf("expected single location retry-after 17, got %d", location.RetryAfter())
	}

	warLog, err := client.GetWarLog(context.Background(), "#2PP", 0, "", "")
	if err != nil {
		t.Fatalf("get war log: %v", err)
	}
	if len(warLog) != 1 {
		t.Fatalf("expected one war log entry, got %d", len(warLog))
	}
	if warLog[0].RetryAfter() != 0 {
		t.Fatalf("expected war log item retry-after to stay unset, got %d", warLog[0].RetryAfter())
	}
}

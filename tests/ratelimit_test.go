package clashy_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	clashy "github.com/clashkinginc/clashy.go"
)

func TestClientThrottleLimitIsEnforced(t *testing.T) {
	t.Parallel()

	var started atomic.Int32
	var inFlight atomic.Int32
	var maxInFlight atomic.Int32
	release := make(chan struct{})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started.Add(1)
		current := inFlight.Add(1)
		setMax(&maxInFlight, current)
		defer inFlight.Add(-1)

		<-release
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"tag":"#2PP","name":"Test","members":0,"memberList":[]}`))
	}))
	defer server.Close()

	cfg := clashy.DefaultClientConfig()
	cfg.BaseURL = server.URL
	cfg.LookupCache = false
	cfg.UpdateCache = false
	cfg.ThrottleLimit = 2

	client, err := clashy.NewClient(cfg)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	if err := client.LoginWithTokens(context.Background(), "token"); err != nil {
		t.Fatalf("login with tokens: %v", err)
	}

	var wg sync.WaitGroup
	for range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := client.GetClan(context.Background(), "#2PP"); err != nil {
				t.Errorf("get clan: %v", err)
			}
		}()
	}

	waitForCount(t, &started, 2, 500*time.Millisecond)
	if got := started.Load(); got != 2 {
		t.Fatalf("expected exactly 2 requests to start before release, got %d", got)
	}
	if got := maxInFlight.Load(); got != 2 {
		t.Fatalf("expected shared client limiter to cap in-flight requests at 2, got %d", got)
	}

	close(release)
	wg.Wait()
}

func TestWithoutRateLimitBypassesClientThrottle(t *testing.T) {
	t.Parallel()

	var started atomic.Int32
	var inFlight atomic.Int32
	var maxInFlight atomic.Int32
	release := make(chan struct{})

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started.Add(1)
		current := inFlight.Add(1)
		setMax(&maxInFlight, current)
		defer inFlight.Add(-1)

		<-release
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"tag":"#2PP","name":"Test","members":0,"memberList":[]}`))
	}))
	defer server.Close()

	cfg := clashy.DefaultClientConfig()
	cfg.BaseURL = server.URL
	cfg.LookupCache = false
	cfg.UpdateCache = false
	cfg.ThrottleLimit = 1

	client, err := clashy.NewClient(cfg)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	if err := client.LoginWithTokens(context.Background(), "token"); err != nil {
		t.Fatalf("login with tokens: %v", err)
	}

	ctx := clashy.WithoutRateLimit(context.Background())
	var wg sync.WaitGroup
	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := client.GetClan(ctx, "#2PP"); err != nil {
				t.Errorf("get clan: %v", err)
			}
		}()
	}

	waitForCount(t, &started, 2, 500*time.Millisecond)
	if got := maxInFlight.Load(); got != 2 {
		t.Fatalf("expected rate-limit bypass context to ignore client concurrency cap, got %d in-flight", got)
	}

	close(release)
	wg.Wait()
}

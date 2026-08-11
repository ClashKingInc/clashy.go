package clashy

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) { return f(request) }

func TestCacheExpiryHonorsDirectivesAndAge(t *testing.T) {
	header := make(http.Header)
	header.Set("Cache-Control", `Public, MAX-AGE="120"`)
	header.Set("Age", "20")
	if got := cacheExpiry(header); got != 100*time.Second {
		t.Fatalf("cache expiry = %s, want 100s", got)
	}
	header.Set("Cache-Control", "max-age=120, no-store")
	if got := cacheExpiry(header); got != 0 {
		t.Fatalf("no-store cache expiry = %s, want 0", got)
	}
}

func TestCIDRMatchingDoesNotUseStringPrefixes(t *testing.T) {
	if !cidrContainsIP("10.0.0.0/24", "10.0.0.12") {
		t.Fatal("expected IP to be inside CIDR")
	}
	if cidrContainsIP("10.0.0.1/32", "10.0.0.10") {
		t.Fatal("string-prefix match incorrectly accepted a different IP")
	}
}

func TestHTTPClientReturnsNamedResultAndSkipsOversizedCacheEntry(t *testing.T) {
	cfg := DefaultClientConfig()
	cfg.ThrottleLimit = 0
	cfg.CacheMaxEntryBytes = 3
	client := NewHTTPClient(cfg)
	var calls atomic.Int32
	client.client.Transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
		calls.Add(1)
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Cache-Control": []string{"max-age=60"}},
			Body:       io.NopCloser(strings.NewReader("four")),
		}, nil
	})

	options := RequestOptions{LookupCache: true, UpdateCache: true}
	for range 2 {
		result, err := client.Do(context.Background(), http.MethodGet, "https://example.test/value", nil, options)
		if err != nil {
			t.Fatalf("Do: %v", err)
		}
		if result.StatusCode != http.StatusOK || string(result.Body) != "four" || result.RetryAfter != time.Minute {
			t.Fatalf("result = %#v", result)
		}
	}
	if got := calls.Load(); got != 2 {
		t.Fatalf("transport calls = %d, want 2 because response exceeds cache entry limit", got)
	}
}

func TestHTTPClientEvictsOldestCacheEntryWithoutShiftingQueue(t *testing.T) {
	cfg := DefaultClientConfig()
	cfg.ThrottleLimit = 0
	cfg.CacheMaxSize = 1
	client := NewHTTPClient(cfg)
	var calls atomic.Int32
	client.client.Transport = roundTripFunc(func(request *http.Request) (*http.Response, error) {
		calls.Add(1)
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Cache-Control": []string{"max-age=60"}},
			Body:       io.NopCloser(strings.NewReader(request.URL.Path)),
		}, nil
	})
	options := RequestOptions{LookupCache: true, UpdateCache: true}
	for _, path := range []string{"/first", "/second", "/first"} {
		if _, err := client.Do(context.Background(), http.MethodGet, "https://example.test"+path, nil, options); err != nil {
			t.Fatalf("Do %s: %v", path, err)
		}
	}
	if got := calls.Load(); got != 3 {
		t.Fatalf("transport calls = %d, want 3 after FIFO eviction", got)
	}
}

func TestNewHTTPClientAcceptsCustomDefaultTransport(t *testing.T) {
	original := http.DefaultTransport
	custom := roundTripFunc(func(*http.Request) (*http.Response, error) { return nil, nil })
	http.DefaultTransport = custom
	t.Cleanup(func() { http.DefaultTransport = original })

	client := NewHTTPClient(DefaultClientConfig())
	if _, ok := client.client.Transport.(roundTripFunc); !ok {
		t.Fatalf("transport = %T, want custom default transport", client.client.Transport)
	}
}

func TestPointerReturningMethodReturnsNilOnHTTPError(t *testing.T) {
	client, err := NewClient(DefaultClientConfig())
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	client.http.client.Transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(`{"reason":"notFound","message":"missing"}`)),
		}, nil
	})

	location, err := client.GetLocation(context.Background(), 1)
	if err == nil {
		t.Fatal("GetLocation error = nil, want typed HTTP error")
	}
	if location != nil {
		t.Fatalf("GetLocation result = %#v, want nil on error", location)
	}
}

func TestPageOptionsNamesCursorsExplicitly(t *testing.T) {
	client, err := NewClient(DefaultClientConfig())
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	var query string
	client.http.client.Transport = roundTripFunc(func(request *http.Request) (*http.Response, error) {
		query = request.URL.RawQuery
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(`{"items":[]}`)),
		}, nil
	})

	_, err = client.GetMembers(context.Background(), "#ABC", PageOptions{
		Limit:  25,
		Before: "older",
		After:  "newer",
	})
	if err != nil {
		t.Fatalf("GetMembers: %v", err)
	}
	if query != "after=newer&before=older&limit=25" {
		t.Fatalf("query = %q", query)
	}
}

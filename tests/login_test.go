package clashy_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	clashy "github.com/clashkinginc/clashy.go"
)

func TestLoginDeveloperUsesLoginCookieForKeyList(t *testing.T) {
	t.Parallel()

	const temporaryToken = "temporary-token"
	const sessionCookie = "developer-session=abc123"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/login":
			if r.Method != http.MethodPost {
				t.Fatalf("login method = %s, want POST", r.Method)
			}
			w.Header().Add("Set-Cookie", sessionCookie+"; Path=/; HttpOnly")
			writeJSON(t, w, `{"temporaryAPIToken":"`+temporaryToken+`"}`)
		case "/api/apikey/list":
			if r.Method != http.MethodPost {
				t.Fatalf("key list method = %s, want POST", r.Method)
			}
			if got := r.Header.Get("Cookie"); got != sessionCookie {
				t.Fatalf("Cookie = %q, want login session cookie", got)
			}
			writeJSON(t, w, `{"keys":[{"id":"existing-key-id","name":"Other key","cidrRanges":["127.0.0.1"],"key":"secret"}]}`)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)

	cfg := clashy.DefaultClientConfig()
	cfg.DeveloperBaseURL = server.URL
	cfg.KeyCount = 0
	cfg.IP = "127.0.0.1"

	client, err := clashy.NewClient(cfg)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	if err := client.Login(context.Background(), "user@example.com", "password"); err != nil {
		t.Fatalf("login: %v", err)
	}
}

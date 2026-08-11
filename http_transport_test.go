package clashy

import (
	"net/http"
	"testing"
	"time"
)

func TestNewHTTPClientAppliesTransportConfig(t *testing.T) {
	cfg := DefaultClientConfig()
	cfg.MaxBaseURLConns = 303
	cfg.IdleConnTimeout = 45 * time.Second

	client := NewHTTPClient(cfg)
	transport, ok := client.client.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("transport type = %T, want *http.Transport", client.client.Transport)
	}
	if transport.MaxIdleConns != cfg.MaxBaseURLConns {
		t.Fatalf("MaxIdleConns = %d, want %d", transport.MaxIdleConns, cfg.MaxBaseURLConns)
	}
	if transport.MaxIdleConnsPerHost != cfg.MaxBaseURLConns {
		t.Fatalf("MaxIdleConnsPerHost = %d, want %d", transport.MaxIdleConnsPerHost, cfg.MaxBaseURLConns)
	}
	if transport.MaxConnsPerHost != cfg.MaxBaseURLConns {
		t.Fatalf("MaxConnsPerHost = %d, want %d", transport.MaxConnsPerHost, cfg.MaxBaseURLConns)
	}
	if transport.IdleConnTimeout != cfg.IdleConnTimeout {
		t.Fatalf("IdleConnTimeout = %s, want %s", transport.IdleConnTimeout, cfg.IdleConnTimeout)
	}
}

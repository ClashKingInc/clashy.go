package clashy

import "time"

// ClientConfig controls API endpoints, authentication behavior, request
// throttling, response caching, and static-data loading.
//
// Most callers should start from DefaultClientConfig and override only the
// fields that differ for their service. The zero value is not the recommended
// production configuration because it has no API base URL or timeout.
type ClientConfig struct {
	// KeyCount is the number of developer-site API keys Login should make
	// available for token rotation.
	KeyCount int
	// KeyNames is the developer-site key name used when reusing or creating API
	// keys during Login.
	KeyNames string
	// ThrottleLimit is the maximum number of concurrent HTTP requests allowed by
	// the client. A value less than or equal to zero disables the limiter.
	ThrottleLimit int
	// Timeout is applied to the underlying http.Client.
	Timeout time.Duration
	// BaseURL is the Clash API or compatible proxy base URL, usually ending in
	// /v1 without a trailing slash.
	BaseURL string
	// DeveloperBaseURL is the developer-site base URL used only by Login.
	DeveloperBaseURL string
	// IP overrides the IP address used when Login creates developer-site API
	// keys. When empty, the IP is inferred from the temporary developer token.
	IP string
	// Realtime adds realtime=true to current-war requests that support it.
	Realtime bool
	// CorrectTags enables Clash tag normalization before tags are placed in API
	// paths or query strings.
	CorrectTags bool
	// CacheMaxSize bounds the number of GET responses retained in memory.
	CacheMaxSize int
	// LookupCache allows GET requests to return fresh cached responses.
	LookupCache bool
	// UpdateCache allows successful GET responses to refresh the in-memory cache.
	UpdateCache bool
	// IgnoreCachedErrors is reserved for compatibility with callers that model
	// cache behavior after coc.py; current request handling does not use it.
	IgnoreCachedErrors []int
	// RawJSON is reserved for callers that need raw response capture; current
	// high-level methods unmarshal into typed models.
	RawJSON bool
	// LoadGameData describes when static game data should be loaded.
	LoadGameData LoadGameData
	// UserAgent is sent with Clash API or proxy requests.
	UserAgent string
	// DeveloperUserAgent is sent with developer-site login and key-management
	// requests.
	DeveloperUserAgent string
}

// DefaultClientConfig returns the recommended baseline configuration for the
// official Clash of Clans API.
//
// The defaults enable tag correction, GET response caching, embedded static
// data, a 30 second timeout, and a conservative request throttle. Callers using
// a ClashKing proxy typically override BaseURL and may enable Realtime.
func DefaultClientConfig() ClientConfig {
	return ClientConfig{
		KeyCount:           1,
		KeyNames:           "Created with clashy.go Client",
		ThrottleLimit:      30,
		Timeout:            30 * time.Second,
		BaseURL:            "https://api.clashofclans.com/v1",
		DeveloperBaseURL:   "https://developer.clashofclans.com",
		CorrectTags:        true,
		CacheMaxSize:       10000,
		LookupCache:        true,
		UpdateCache:        true,
		LoadGameData:       DefaultLoadGameData(),
		UserAgent:          "clashy.go",
		DeveloperUserAgent: "clashy.go/devsite",
	}
}

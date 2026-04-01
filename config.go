package clashy

import "time"

type ClientConfig struct {
	KeyCount           int
	KeyNames           string
	ThrottleLimit      int
	Timeout            time.Duration
	BaseURL            string
	DeveloperBaseURL   string
	IP                 string
	Realtime           bool
	CorrectTags        bool
	CacheMaxSize       int
	LookupCache        bool
	UpdateCache        bool
	IgnoreCachedErrors []int
	RawJSON            bool
	LoadGameData       LoadGameData
	UserAgent          string
	DeveloperUserAgent string
}

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

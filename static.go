package clashy

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

//go:embed static/static_data.json
var staticDataBytes []byte

//go:embed static/translations.json
var translationsBytes []byte

// StaticData is the parsed and indexed ClashKing static data embedded in the
// package.
type StaticData struct {
	// Raw preserves static-data sections exactly as parsed from static_data.json.
	Raw map[string][]map[string]any
	// ByID indexes static-data entries by their numeric _id value.
	ByID map[int]map[string]any
	// ByName indexes static-data entries by normalized name, section, and
	// village.
	ByName map[string]map[string]any
	// Translations maps translation IDs to language-code/value maps.
	Translations map[string]map[string]string
}

var (
	staticOnce sync.Once
	staticSet  *StaticData
	staticErr  error
)

const (
	staticDataURL    = "https://assets.clashk.ing/static_data.json"
	translationsURL  = "https://assets.clashk.ing/translations.json"
	staticDataPath   = "static/static_data.json"
	translationsPath = "static/translations.json"
)

// LoadStaticData parses the embedded static-data files once and returns the
// shared indexed result.
func LoadStaticData() (*StaticData, error) {
	staticOnce.Do(func() {
		staticSet, staticErr = parseStaticData(staticDataBytes, translationsBytes)
	})
	return staticSet, staticErr
}

// UpdateStatic downloads the latest ClashKing static-data and translation JSON,
// writes the embedded source files, and refreshes this client's in-memory
// StaticData.
func (c *Client) UpdateStatic(ctx context.Context) error {
	if err := downloadJSON(ctx, staticDataURL, staticDataPath); err != nil {
		return err
	}
	if err := downloadJSON(ctx, translationsURL, translationsPath); err != nil {
		return err
	}
	staticBytes, err := os.ReadFile(staticDataPath)
	if err != nil {
		return err
	}
	translationBytes, err := os.ReadFile(translationsPath)
	if err != nil {
		return err
	}
	updated, err := parseStaticData(staticBytes, translationBytes)
	if err != nil {
		return err
	}
	staticOnce = sync.Once{}
	staticOnce.Do(func() {
		staticSet = updated
		staticErr = nil
	})
	c.staticData = updated
	return nil
}

func downloadJSON(ctx context.Context, url, path string) error {
	if ctx == nil {
		ctx = context.Background()
	}
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("download %s: unexpected status %s", url, resp.Status)
	}
	body, err := readLimited(resp.Body, 64<<20)
	if err != nil {
		return err
	}
	var payload any
	if err := json.Unmarshal(body, &payload); err != nil {
		return fmt.Errorf("download %s: invalid json: %w", url, err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, body, 0o644)
}

func parseStaticData(staticBytes, translationBytes []byte) (*StaticData, error) {
	var raw map[string][]map[string]any
	if err := json.Unmarshal(staticBytes, &raw); err != nil {
		return nil, err
	}
	var translations map[string]map[string]string
	if err := json.Unmarshal(translationBytes, &translations); err != nil {
		return nil, err
	}

	s := &StaticData{
		Raw:          raw,
		ByID:         make(map[int]map[string]any),
		ByName:       make(map[string]map[string]any),
		Translations: translations,
	}
	for section, items := range raw {
		for _, item := range items {
			id, ok := asInt(item["_id"])
			if ok {
				s.ByID[id] = item
			}
			name, _ := item["name"].(string)
			village, _ := item["village"].(string)
			if name != "" {
				key := staticLookupKey(name, section, village)
				s.ByName[key] = item
			}
		}
	}
	return s, nil
}

func staticLookupKey(name, section, village string) string {
	return strings.ToLower(fmt.Sprintf("%s|%s|%s", name, section, village))
}

// LookupByName returns a static-data entry by display name, section, and
// village.
//
// The lookup is case-insensitive. The section should match a top-level static
// data section such as "troops", "spells", "heroes", "pets", or "equipment".
func (s *StaticData) LookupByName(name, section, village string) map[string]any {
	if s == nil {
		return nil
	}
	return s.ByName[staticLookupKey(name, section, village)]
}

// LookupByID returns a static-data entry by numeric static ID.
func (s *StaticData) LookupByID(id int) map[string]any {
	if s == nil {
		return nil
	}
	return s.ByID[id]
}

func asInt(v any) (int, bool) {
	switch n := v.(type) {
	case float64:
		return int(n), true
	case int:
		return n, true
	case int32:
		return int(n), true
	case int64:
		return int(n), true
	default:
		return 0, false
	}
}

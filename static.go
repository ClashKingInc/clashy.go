package clashy

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

//go:embed static/static_data.json
var staticDataBytes []byte

//go:embed static/translations.json
var translationsBytes []byte

// StaticData is the parsed and indexed ClashKing static data embedded in the
// package.
type StaticData struct {
	raw          map[string][]map[string]any
	byID         map[int]map[string]any
	byName       map[string]map[string]any
	translations map[string]map[string]string
}

var (
	staticOnce sync.Once
	staticSet  *StaticData
	staticErr  error
)

// LoadStaticData parses the embedded static-data files once and returns the
// shared indexed result.
func LoadStaticData() (*StaticData, error) {
	staticOnce.Do(func() {
		staticSet, staticErr = parseStaticData(staticDataBytes, translationsBytes)
	})
	return staticSet, staticErr
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
		raw:          raw,
		byID:         make(map[int]map[string]any),
		byName:       make(map[string]map[string]any),
		translations: translations,
	}
	for section, items := range raw {
		for _, item := range items {
			id, ok := asInt(item["_id"])
			if ok {
				s.byID[id] = item
			}
			name, _ := item["name"].(string)
			village, _ := item["village"].(string)
			if name != "" {
				key := staticLookupKey(name, section, village)
				s.byName[key] = item
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
	return cloneStaticMap(s.lookupByName(name, section, village))
}

// LookupByID returns a static-data entry by numeric static ID.
func (s *StaticData) LookupByID(id int) map[string]any {
	return cloneStaticMap(s.lookupByID(id))
}

// Section returns an isolated copy of one top-level static-data section.
func (s *StaticData) Section(name string) []map[string]any {
	if s == nil {
		return nil
	}
	items := s.raw[name]
	out := make([]map[string]any, len(items))
	for i := range items {
		out[i] = cloneStaticMap(items[i])
	}
	return out
}

// Sections returns an isolated copy of all top-level static-data sections.
func (s *StaticData) Sections() map[string][]map[string]any {
	if s == nil {
		return nil
	}
	out := make(map[string][]map[string]any, len(s.raw))
	for name := range s.raw {
		out[name] = s.Section(name)
	}
	return out
}

// Translation returns an isolated language map for one translation ID.
func (s *StaticData) Translation(id string) map[string]string {
	if s == nil || len(s.translations[id]) == 0 {
		return nil
	}
	out := make(map[string]string, len(s.translations[id]))
	for language, value := range s.translations[id] {
		out[language] = value
	}
	return out
}

// Translations returns an isolated copy of all translations.
func (s *StaticData) Translations() map[string]map[string]string {
	if s == nil {
		return nil
	}
	out := make(map[string]map[string]string, len(s.translations))
	for id := range s.translations {
		out[id] = s.Translation(id)
	}
	return out
}

func (s *StaticData) lookupByName(name, section, village string) map[string]any {
	if s == nil {
		return nil
	}
	return s.byName[staticLookupKey(name, section, village)]
}

func (s *StaticData) lookupByID(id int) map[string]any {
	if s == nil {
		return nil
	}
	return s.byID[id]
}

func cloneStaticMap(source map[string]any) map[string]any {
	if source == nil {
		return nil
	}
	out := make(map[string]any, len(source))
	for key, value := range source {
		out[key] = cloneStaticValue(value)
	}
	return out
}

func cloneStaticValue(value any) any {
	switch value := value.(type) {
	case map[string]any:
		return cloneStaticMap(value)
	case []any:
		out := make([]any, len(value))
		for i := range value {
			out[i] = cloneStaticValue(value[i])
		}
		return out
	case []map[string]any:
		out := make([]map[string]any, len(value))
		for i := range value {
			out[i] = cloneStaticMap(value[i])
		}
		return out
	default:
		return value
	}
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

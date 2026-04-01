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

type StaticData struct {
	Raw          map[string][]map[string]any
	ByID         map[int]map[string]any
	ByName       map[string]map[string]any
	Translations map[string]map[string]string
}

var (
	staticOnce sync.Once
	staticSet  *StaticData
	staticErr  error
)

func LoadStaticData() (*StaticData, error) {
	staticOnce.Do(func() {
		var raw map[string][]map[string]any
		if err := json.Unmarshal(staticDataBytes, &raw); err != nil {
			staticErr = err
			return
		}
		var translations map[string]map[string]string
		if err := json.Unmarshal(translationsBytes, &translations); err != nil {
			staticErr = err
			return
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
		staticSet = s
	})
	return staticSet, staticErr
}

func staticLookupKey(name, section, village string) string {
	return strings.ToLower(fmt.Sprintf("%s|%s|%s", name, section, village))
}

func (s *StaticData) LookupByName(name, section, village string) map[string]any {
	if s == nil {
		return nil
	}
	return s.ByName[staticLookupKey(name, section, village)]
}

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

package clashy

import (
	"encoding/json"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type StaticUnit struct {
	Name        string
	Level       int
	MaxLevel    int
	Village     string
	UpgradeCost int
	UpgradeTime time.Duration
	Raw         map[string]any
}

func fillStaticUnit(unit *StaticUnit, data map[string]any, level int) {
	if data == nil {
		return
	}
	unit.Raw = data
	if unit.Name == "" {
		unit.Name, _ = data["name"].(string)
	}
	if unit.Village == "" {
		unit.Village, _ = data["village"].(string)
	}
	maxLevel := 0
	if levels, ok := data["levels"].([]any); ok {
		for _, rawLevel := range levels {
			block, ok := rawLevel.(map[string]any)
			if !ok {
				continue
			}
			current, ok := asInt(block["level"])
			if !ok {
				continue
			}
			if current > maxLevel {
				maxLevel = current
			}
			if level != 0 && current == level {
				unit.Level = level
				unit.UpgradeCost, _ = asInt(block["upgrade_cost"])
				if secs, ok := asInt(block["upgrade_time"]); ok {
					unit.UpgradeTime = time.Duration(secs) * time.Second
				}
			}
		}
	} else {
		for _, v := range data {
			block, ok := v.(map[string]any)
			if !ok {
				continue
			}
			visual, ok := asInt(block["VisualLevel"])
			if !ok {
				continue
			}
			if visual > maxLevel {
				maxLevel = visual
			}
			if level != 0 && visual == level {
				unit.Level = level
				unit.UpgradeCost, _ = asInt(block["UpgradeCost"])
				secs := secondsFromParts(block["UpgradeTimeD"], block["UpgradeTimeH"], block["UpgradeTimeM"], block["UpgradeTimeS"])
				unit.UpgradeTime = time.Duration(secs) * time.Second
			}
		}
	}
	if unit.Level == 0 {
		unit.Level = max(level, 1)
	}
	unit.MaxLevel = maxLevel
}

func secondsFromParts(parts ...any) int {
	factors := []int{86400, 3600, 60, 1}
	total := 0
	for i, p := range parts {
		v, _ := asInt(p)
		total += v * factors[i]
	}
	return total
}

type Troop struct {
	Name               string `json:"name"`
	Level              int    `json:"level"`
	MaxLevel           int    `json:"maxLevel"`
	Village            string `json:"village"`
	SuperTroopIsActive bool   `json:"superTroopIsActive"`
	StaticUnit
}

func (t *Troop) postDecode(c *Client) {
	if c == nil || c.staticData == nil {
		return
	}
	fillStaticUnit(&t.StaticUnit, c.staticData.LookupByName(t.Name, "troops", t.Village), t.Level)
	t.Name = firstNonEmpty(t.Name, t.StaticUnit.Name)
	t.Village = firstNonEmpty(t.Village, t.StaticUnit.Village)
	t.Level = max(t.Level, t.StaticUnit.Level)
	t.MaxLevel = max(t.MaxLevel, t.StaticUnit.MaxLevel)
}

func (t Troop) IsHomeBase() bool    { return t.Village == "home" || t.Village == "" }
func (t Troop) IsBuilderBase() bool { return t.Village == "builderBase" }

type Spell struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
	Village  string `json:"village"`
	StaticUnit
}

func (s *Spell) postDecode(c *Client) {
	if c == nil || c.staticData == nil {
		return
	}
	fillStaticUnit(&s.StaticUnit, c.staticData.LookupByName(s.Name, "spells", s.Village), s.Level)
	s.Name = firstNonEmpty(s.Name, s.StaticUnit.Name)
	s.Level = max(s.Level, s.StaticUnit.Level)
	s.MaxLevel = max(s.MaxLevel, s.StaticUnit.MaxLevel)
}

type Hero struct {
	Name      string      `json:"name"`
	Level     int         `json:"level"`
	MaxLevel  int         `json:"maxLevel"`
	Village   string      `json:"village"`
	Equipment []Equipment `json:"equipment"`
	StaticUnit
}

func (h *Hero) postDecode(c *Client) {
	if c == nil || c.staticData == nil {
		return
	}
	fillStaticUnit(&h.StaticUnit, c.staticData.LookupByName(h.Name, "heroes", h.Village), h.Level)
	h.Name = firstNonEmpty(h.Name, h.StaticUnit.Name)
	h.Level = max(h.Level, h.StaticUnit.Level)
	h.MaxLevel = max(h.MaxLevel, h.StaticUnit.MaxLevel)
	for i := range h.Equipment {
		h.Equipment[i].postDecode(c)
	}
}

type Pet struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
	Village  string `json:"village"`
	StaticUnit
}

func (p *Pet) postDecode(c *Client) {
	if c == nil || c.staticData == nil {
		return
	}
	fillStaticUnit(&p.StaticUnit, c.staticData.LookupByName(p.Name, "pets", p.Village), p.Level)
	p.Name = firstNonEmpty(p.Name, p.StaticUnit.Name)
	p.Level = max(p.Level, p.StaticUnit.Level)
	p.MaxLevel = max(p.MaxLevel, p.StaticUnit.MaxLevel)
}

type Equipment struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
	Village  string `json:"village"`
	Rarity   string `json:"rarity"`
	StaticUnit
}

func (e *Equipment) postDecode(c *Client) {
	if c == nil || c.staticData == nil {
		return
	}
	fillStaticUnit(&e.StaticUnit, c.staticData.LookupByName(e.Name, "equipment", e.Village), e.Level)
	e.Name = firstNonEmpty(e.Name, e.StaticUnit.Name)
	e.Level = max(e.Level, e.StaticUnit.Level)
	e.MaxLevel = max(e.MaxLevel, e.StaticUnit.MaxLevel)
}

type HeroLoadout struct {
	Hero      Hero
	Pet       *Pet
	Equipment []Equipment
}

type ArmyRecipe struct {
	Link             string
	HeroesLoadout    []HeroLoadout
	Troops           []TroopCount
	Spells           []SpellCount
	ClanCastleTroops []TroopCount
	ClanCastleSpells []SpellCount
}

type TroopCount struct {
	Troop    Troop
	Quantity int
}

type SpellCount struct {
	Spell    Spell
	Quantity int
}

func ParseArmyRecipe(static *StaticData, link string) ArmyRecipe {
	recipe := ArmyRecipe{Link: link}
	if static == nil {
		return recipe
	}
	armyLinkSeparator := regexp.MustCompile(`h(?P<heroes>[^idus]+)|i(?P<castle_troops>[\d+x-]+)|d(?P<castle_spells>[\d+x-]+)|u(?P<units>[\d+x-]+)|s(?P<spells>[\d+x-]+)`)
	heroPattern := regexp.MustCompile(`(?P<hero_id>\d+)(?:m\d+)?(?:p(?P<pet_id>\d+))?(?:e(?P<eq1>\d+)(?:_(?P<eq2>\d+))?)?`)
	for _, match := range armyLinkSeparator.FindAllString(link, -1) {
		switch {
		case strings.HasPrefix(match, "h"):
			for _, entry := range strings.Split(strings.TrimPrefix(match, "h"), "-") {
				g := heroPattern.FindStringSubmatch(entry)
				if g == nil {
					continue
				}
				heroID, _ := strconv.Atoi(g[1])
				loadout := HeroLoadout{
					Hero: buildHeroFromStatic(static.LookupByID(HeroBaseID+heroID), 1),
				}
				if len(g) > 2 && g[2] != "" {
					petID, _ := strconv.Atoi(g[2])
					pet := buildPetFromStatic(static.LookupByID(PetBaseID+petID), 1)
					loadout.Pet = &pet
				}
				if len(g) > 3 && g[3] != "" {
					eq1, _ := strconv.Atoi(g[3])
					loadout.Equipment = append(loadout.Equipment, buildEquipmentFromStatic(static.LookupByID(EquipmentBaseID+eq1), 1))
				}
				if len(g) > 4 && g[4] != "" {
					eq2, _ := strconv.Atoi(g[4])
					loadout.Equipment = append(loadout.Equipment, buildEquipmentFromStatic(static.LookupByID(EquipmentBaseID+eq2), 1))
				}
				recipe.HeroesLoadout = append(recipe.HeroesLoadout, loadout)
			}
		case strings.HasPrefix(match, "i"):
			parseArmyItems(static, strings.TrimPrefix(match, "i"), TroopBaseID, true, &recipe)
		case strings.HasPrefix(match, "d"):
			parseArmyItems(static, strings.TrimPrefix(match, "d"), SpellBaseID, false, &recipe)
		case strings.HasPrefix(match, "u"):
			parseArmyItems(static, strings.TrimPrefix(match, "u"), TroopBaseID, true, &recipe)
		case strings.HasPrefix(match, "s"):
			parseArmyItems(static, strings.TrimPrefix(match, "s"), SpellBaseID, false, &recipe)
		}
	}
	return recipe
}

func parseArmyItems(static *StaticData, payload string, baseID int, troops bool, recipe *ArmyRecipe) {
	parts := strings.Split(payload, "-")
	for _, part := range parts {
		split := strings.Split(part, "x")
		if len(split) != 2 {
			continue
		}
		qty, _ := strconv.Atoi(split[0])
		id, _ := strconv.Atoi(split[1])
		if troops {
			tc := TroopCount{Troop: buildTroopFromStatic(static.LookupByID(baseID+id), 1), Quantity: qty}
			if strings.Contains(recipe.Link, "i"+payload) {
				recipe.ClanCastleTroops = append(recipe.ClanCastleTroops, tc)
			} else {
				recipe.Troops = append(recipe.Troops, tc)
			}
		} else {
			sc := SpellCount{Spell: buildSpellFromStatic(static.LookupByID(baseID+id), 1), Quantity: qty}
			if strings.Contains(recipe.Link, "d"+payload) {
				recipe.ClanCastleSpells = append(recipe.ClanCastleSpells, sc)
			} else {
				recipe.Spells = append(recipe.Spells, sc)
			}
		}
	}
}

func buildTroopFromStatic(data map[string]any, level int) Troop {
	t := Troop{}
	fillStaticUnit(&t.StaticUnit, data, level)
	t.Name, t.Level, t.MaxLevel, t.Village = t.StaticUnit.Name, t.StaticUnit.Level, t.StaticUnit.MaxLevel, t.StaticUnit.Village
	return t
}

func buildSpellFromStatic(data map[string]any, level int) Spell {
	s := Spell{}
	fillStaticUnit(&s.StaticUnit, data, level)
	s.Name, s.Level, s.MaxLevel, s.Village = s.StaticUnit.Name, s.StaticUnit.Level, s.StaticUnit.MaxLevel, s.StaticUnit.Village
	return s
}

func buildHeroFromStatic(data map[string]any, level int) Hero {
	h := Hero{}
	fillStaticUnit(&h.StaticUnit, data, level)
	h.Name, h.Level, h.MaxLevel, h.Village = h.StaticUnit.Name, h.StaticUnit.Level, h.StaticUnit.MaxLevel, h.StaticUnit.Village
	return h
}

func buildPetFromStatic(data map[string]any, level int) Pet {
	p := Pet{}
	fillStaticUnit(&p.StaticUnit, data, level)
	p.Name, p.Level, p.MaxLevel, p.Village = p.StaticUnit.Name, p.StaticUnit.Level, p.StaticUnit.MaxLevel, p.StaticUnit.Village
	return p
}

func buildEquipmentFromStatic(data map[string]any, level int) Equipment {
	e := Equipment{}
	fillStaticUnit(&e.StaticUnit, data, level)
	e.Name, e.Level, e.MaxLevel, e.Village = e.StaticUnit.Name, e.StaticUnit.Level, e.StaticUnit.MaxLevel, e.StaticUnit.Village
	return e
}

type AccountData struct {
	Raw map[string]any
}

func ParseAccountData(data map[string]any) AccountData {
	return AccountData{Raw: data}
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func decodeJSONMap(data []byte) map[string]any {
	var out map[string]any
	_ = json.Unmarshal(data, &out)
	return out
}

func parseURLValues(raw string) url.Values {
	u, err := url.Parse(raw)
	if err != nil {
		return nil
	}
	return u.Query()
}

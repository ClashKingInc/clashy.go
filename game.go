package clashy

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var armyHeroPattern = regexp.MustCompile(`^(\d+)(?:m\d+)?(?:p(\d+))?(?:e(\d+)(?:_(\d+))?)?$`)

type StaticUnit struct {
	Name        string
	Level       int
	MaxLevel    int
	Village     string
	UpgradeCost int
	UpgradeTime time.Duration
}

func fillStaticUnit(unit *StaticUnit, data map[string]any, level int) {
	if data == nil {
		return
	}
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

func (t Troop) IsHomeBase() bool    { return t.Village == "home" || t.Village == "" }
func (t Troop) IsBuilderBase() bool { return t.Village == "builderBase" }

func (t Troop) IsSuperTroop() bool {
	for _, name := range SuperTroopOrder {
		if t.Name == name {
			return true
		}
	}
	return false
}

func (t Troop) Static(c *Client) *Troop {
	if c == nil {
		return nil
	}
	return c.GetTroop(t.Name, t.IsHomeBase(), t.Level)
}

type Spell struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
	Village  string `json:"village"`
	StaticUnit
}

func (s Spell) Static(c *Client) *Spell {
	if c == nil {
		return nil
	}
	return c.GetSpell(s.Name, s.Level)
}

type Hero struct {
	Name      string      `json:"name"`
	Level     int         `json:"level"`
	MaxLevel  int         `json:"maxLevel"`
	Village   string      `json:"village"`
	Equipment []Equipment `json:"equipment"`
	StaticUnit
}

func (h Hero) Static(c *Client) *Hero {
	if c == nil {
		return nil
	}
	return c.GetHero(h.Name, h.Level)
}

type Pet struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
	Village  string `json:"village"`
	StaticUnit
}

func (p Pet) Static(c *Client) *Pet {
	if c == nil {
		return nil
	}
	return c.GetPet(p.Name, p.Level)
}

type Equipment struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	MaxLevel int    `json:"maxLevel"`
	Village  string `json:"village"`
	Rarity   string `json:"rarity"`
	StaticUnit
}

func (e Equipment) Static(c *Client) *Equipment {
	if c == nil {
		return nil
	}
	return c.GetEquipment(e.Name, e.Level)
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
	for _, match := range splitArmySections(extractArmyPayload(link)) {
		if len(match) < 2 {
			continue
		}
		switch {
		case strings.HasPrefix(match, "h"):
			for _, entry := range strings.Split(strings.TrimPrefix(match, "h"), "-") {
				g := armyHeroPattern.FindStringSubmatch(entry)
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
			parseArmyItems(static, strings.TrimPrefix(match, "i"), TroopBaseID, true, true, &recipe)
		case strings.HasPrefix(match, "d"):
			parseArmyItems(static, strings.TrimPrefix(match, "d"), SpellBaseID, false, true, &recipe)
		case strings.HasPrefix(match, "u"):
			parseArmyItems(static, strings.TrimPrefix(match, "u"), TroopBaseID, true, false, &recipe)
		case strings.HasPrefix(match, "s"):
			parseArmyItems(static, strings.TrimPrefix(match, "s"), SpellBaseID, false, false, &recipe)
		}
	}
	return recipe
}

func extractArmyPayload(link string) string {
	parsed, err := url.Parse(link)
	if err == nil {
		if army := parsed.Query().Get("army"); army != "" {
			return army
		}
	}
	return link
}

func splitArmySections(payload string) []string {
	var sections []string
	start := -1
	for i := 0; i < len(payload); i++ {
		if !isArmySectionMarker(payload[i]) {
			continue
		}
		if start >= 0 {
			sections = append(sections, payload[start:i])
		}
		start = i
	}
	if start >= 0 {
		sections = append(sections, payload[start:])
	}
	return sections
}

func isArmySectionMarker(b byte) bool {
	switch b {
	case 'h', 'i', 'd', 'u', 's':
		return true
	default:
		return false
	}
}

func parseArmyItems(static *StaticData, payload string, baseID int, troops bool, clanCastle bool, recipe *ArmyRecipe) {
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
			if clanCastle {
				recipe.ClanCastleTroops = append(recipe.ClanCastleTroops, tc)
			} else {
				recipe.Troops = append(recipe.Troops, tc)
			}
		} else {
			sc := SpellCount{Spell: buildSpellFromStatic(static.LookupByID(baseID+id), 1), Quantity: qty}
			if clanCastle {
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

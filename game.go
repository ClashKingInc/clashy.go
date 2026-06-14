package clashy

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var armyHeroPattern = regexp.MustCompile(`^(\d+)(?:m\d+)?(?:p(\d+))?(?:e(\d+)(?:_(\d+))?)?$`)

// StaticUnit contains normalized static-data fields shared by troops, spells,
// heroes, pets, and hero equipment.
type StaticUnit struct {
	// Name is the unit or equipment display name.
	Name string
	// Level is the selected level for this static lookup.
	Level int
	// MaxLevel is the maximum level found in static data.
	MaxLevel int
	// Village identifies the village this object belongs to.
	Village string
	// UpgradeCost is the cost for the selected level when static data includes
	// it.
	UpgradeCost int
	// UpgradeTime is the upgrade duration for the selected level when static data
	// includes it.
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

// Troop is a player troop or static troop lookup result.
type Troop struct {
	// Name is the troop display name.
	Name string `json:"name"`
	// Level is the player's current level or the selected static level.
	Level int `json:"level"`
	// MaxLevel is the maximum level available for the player's Town Hall or in
	// static data.
	MaxLevel int `json:"maxLevel"`
	// Village identifies home or Builder Base troops.
	Village string `json:"village"`
	// SuperTroopIsActive reports whether a super troop boost is active.
	SuperTroopIsActive bool `json:"superTroopIsActive"`
	StaticUnit
}

// IsHomeBase reports whether the troop belongs to the home village.
func (t Troop) IsHomeBase() bool { return t.Village == "home" || t.Village == "" }

// IsBuilderBase reports whether the troop belongs to Builder Base.
func (t Troop) IsBuilderBase() bool { return t.Village == "builderBase" }

// IsSuperTroop reports whether the troop name is one of the known super troops.
func (t Troop) IsSuperTroop() bool {
	for _, name := range SuperTroopOrder {
		if t.Name == name {
			return true
		}
	}
	return false
}

// Static returns the embedded static-data record matching this troop's name,
// village, and level.
func (t Troop) Static(c *Client) *Troop {
	if c == nil {
		return nil
	}
	return c.GetTroop(t.Name, t.IsHomeBase(), t.Level)
}

// Spell is a player spell or static spell lookup result.
type Spell struct {
	// Name is the spell display name.
	Name string `json:"name"`
	// Level is the player's current level or the selected static level.
	Level int `json:"level"`
	// MaxLevel is the maximum level available for the player or in static data.
	MaxLevel int `json:"maxLevel"`
	// Village identifies the spell's village when static data provides one.
	Village string `json:"village"`
	StaticUnit
}

// Static returns the embedded static-data record matching this spell's name and
// level.
func (s Spell) Static(c *Client) *Spell {
	if c == nil {
		return nil
	}
	return c.GetSpell(s.Name, s.Level)
}

// Hero is a player hero or static hero lookup result.
type Hero struct {
	// Name is the hero display name.
	Name string `json:"name"`
	// Level is the player's current level or the selected static level.
	Level int `json:"level"`
	// MaxLevel is the maximum level available for the player or in static data.
	MaxLevel int `json:"maxLevel"`
	// Village identifies the hero's village.
	Village string `json:"village"`
	// Equipment contains equipment currently assigned to this hero when the API
	// includes loadout data.
	Equipment []Equipment `json:"equipment"`
	StaticUnit
}

// Static returns the embedded static-data record matching this hero's name and
// level.
func (h Hero) Static(c *Client) *Hero {
	if c == nil {
		return nil
	}
	return c.GetHero(h.Name, h.Level)
}

// Pet is a hero pet from a player response or static lookup.
type Pet struct {
	// Name is the pet display name.
	Name string `json:"name"`
	// Level is the player's current level or the selected static level.
	Level int `json:"level"`
	// MaxLevel is the maximum level available for the player or in static data.
	MaxLevel int `json:"maxLevel"`
	// Village identifies the pet's village.
	Village string `json:"village"`
	StaticUnit
}

// Static returns the embedded static-data record matching this pet's name and
// level.
func (p Pet) Static(c *Client) *Pet {
	if c == nil {
		return nil
	}
	return c.GetPet(p.Name, p.Level)
}

// Equipment is hero equipment from a player response or static lookup.
type Equipment struct {
	// Name is the equipment display name.
	Name string `json:"name"`
	// Level is the player's current level or the selected static level.
	Level int `json:"level"`
	// MaxLevel is the maximum level available for the player or in static data.
	MaxLevel int `json:"maxLevel"`
	// Village identifies the equipment's village.
	Village string `json:"village"`
	// Rarity is the equipment rarity when static data includes it.
	Rarity string `json:"rarity"`
	StaticUnit
}

// Static returns the embedded static-data record matching this equipment's name
// and level.
func (e Equipment) Static(c *Client) *Equipment {
	if c == nil {
		return nil
	}
	return c.GetEquipment(e.Name, e.Level)
}

// HeroLoadout is one hero, pet, and equipment grouping parsed from an army link.
type HeroLoadout struct {
	// Hero is the hero selected in the army link.
	Hero Hero
	// Pet is the assigned pet when the link includes one.
	Pet *Pet
	// Equipment is the selected hero equipment in link order.
	Equipment []Equipment
}

// ArmyRecipe is the normalized representation of a Clash army link.
type ArmyRecipe struct {
	// Link is the original link or raw army payload passed by the caller.
	Link string
	// HeroesLoadout contains heroes, pets, and equipment from the link.
	HeroesLoadout []HeroLoadout
	// Troops contains home-army troops from the link.
	Troops []TroopCount
	// Spells contains home-army spells from the link.
	Spells []SpellCount
	// ClanCastleTroops contains requested Clan Castle troops.
	ClanCastleTroops []TroopCount
	// ClanCastleSpells contains requested Clan Castle spells.
	ClanCastleSpells []SpellCount
}

// TroopCount pairs a troop with a quantity from an army link.
type TroopCount struct {
	// Troop is the parsed troop.
	Troop Troop
	// Quantity is the requested troop count.
	Quantity int
}

// SpellCount pairs a spell with a quantity from an army link.
type SpellCount struct {
	// Spell is the parsed spell.
	Spell Spell
	// Quantity is the requested spell count.
	Quantity int
}

// ParseArmyRecipe parses a full Clash army link or raw army payload into a
// structured recipe using embedded static data for names and villages.
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

// AccountData is a thin wrapper around arbitrary account-link data.
type AccountData struct {
	// Raw contains the original account-link payload.
	Raw map[string]any
}

// ParseAccountData wraps account-link data without mutating it.
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

package clashy

import (
	"reflect"
	"testing"
)

func TestStaticDataAccessorsReturnIsolatedCopies(t *testing.T) {
	staticData, err := LoadStaticData()
	if err != nil {
		t.Fatalf("LoadStaticData: %v", err)
	}

	first := staticData.Section("troops")
	if len(first) == 0 {
		t.Fatal("troops section is empty")
	}
	originalName, _ := first[0]["name"].(string)
	first[0]["name"] = "mutated"
	second := staticData.Section("troops")
	if got, _ := second[0]["name"].(string); got != originalName {
		t.Fatalf("shared section was mutated: got %q, want %q", got, originalName)
	}

	lookup := staticData.LookupByID(TroopBaseID)
	if lookup != nil {
		lookup["name"] = "mutated"
		if again := staticData.LookupByID(TroopBaseID); again != nil && again["name"] == "mutated" {
			t.Fatal("LookupByID returned a mutable shared map")
		}
	}
}

func TestSeasonalTroopOrderUsesLastOccurrencesFromStaticData(t *testing.T) {
	staticData, err := LoadStaticData()
	if err != nil {
		t.Fatalf("LoadStaticData: %v", err)
	}

	seasonalNames := make([]string, 0)
	for _, troop := range staticData.Section("troops") {
		seasonal, _ := troop["is_seasonal"].(bool)
		name, _ := troop["name"].(string)
		if seasonal && name != "" {
			seasonalNames = append(seasonalNames, name)
		}
	}

	last := make(map[string]int, len(seasonalNames))
	for i, name := range seasonalNames {
		last[name] = i
	}
	expected := make([]string, 0, len(last))
	for i, name := range seasonalNames {
		if last[name] == i {
			expected = append(expected, name)
		}
	}

	if !reflect.DeepEqual(SeasonalTroopOrder, expected) {
		t.Fatalf("SeasonalTroopOrder = %v, want last-occurrence order %v", SeasonalTroopOrder, expected)
	}
}

func TestParseArmyRecipeSkipsMalformedAndUnknownItems(t *testing.T) {
	staticData, err := LoadStaticData()
	if err != nil {
		t.Fatalf("LoadStaticData: %v", err)
	}
	recipe := ParseArmyRecipe(staticData, "u0x0-badx0-2xbad-2x999999h999999p999999e999999")
	if len(recipe.Troops) != 0 || len(recipe.HeroesLoadout) != 0 {
		t.Fatalf("malformed recipe created fake units: %#v", recipe)
	}
}

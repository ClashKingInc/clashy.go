package clashy

import "testing"

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

func TestSeasonalTroopOrderUsesLastDuplicateOccurrence(t *testing.T) {
	indices := map[string][]int{}
	for i, name := range SeasonalTroopOrder {
		indices[name] = append(indices[name], i)
	}
	for _, name := range []string{"YEETer", "The Disarmer"} {
		if got := len(indices[name]); got != 1 {
			t.Fatalf("%s appears %d times, want once", name, got)
		}
	}
	if indices["YEETer"][0] >= indices["The Disarmer"][0] {
		t.Fatalf("last-occurrence order not preserved: YEETer=%d Disarmer=%d", indices["YEETer"][0], indices["The Disarmer"][0])
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

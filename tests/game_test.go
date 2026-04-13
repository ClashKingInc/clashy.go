package clashy_test

import (
	"testing"

	clashy "github.com/clashkinginc/clashy.go"
)

func TestParseArmyRecipeExampleLinks(t *testing.T) {
	static, err := clashy.LoadStaticData()
	if err != nil {
		t.Fatalf("load static data: %v", err)
	}

	tests := []struct {
		name               string
		link               string
		heroes             int
		clanCastleTroops   int
		clanCastleSpells   int
		troops             int
		spells             int
		firstHeroPet       bool
		firstHeroEquipment int
	}{
		{
			name:               "full army link",
			link:               "https://link.clashofclans.com/en?action=CopyArmy&army=h0p4e8_14-1p9e39_48-2p7e34_4-4p10e13_40i3x53-1x0-1x135-1x87d1x53-1x9u1x23-2x53-1x12-2x58-4x6-1x10-2x28-1x82-1x97-10x11-12x98-2x1-1x87-1x91-1x135s4x35-2x2-2x5-1x17",
			heroes:             4,
			clanCastleTroops:   4,
			clanCastleSpells:   2,
			troops:             15,
			spells:             4,
			firstHeroPet:       true,
			firstHeroEquipment: 2,
		},
		{
			name:               "mixed optional hero fields",
			link:               "https://link.clashofclans.com/en?action=CopyArmy&army=h6p11e35_43-2m1p16e4_24-0p4e32_14-7e52_53i1x65-6x5-1x62d4x120s5x120",
			heroes:             4,
			clanCastleTroops:   3,
			clanCastleSpells:   1,
			troops:             0,
			spells:             1,
			firstHeroPet:       true,
			firstHeroEquipment: 2,
		},
		{
			name:             "troops only",
			link:             "https://link.clashofclans.com/en?action=CopyArmy&army=u1x0-1x1-1x4",
			heroes:           0,
			clanCastleTroops: 0,
			clanCastleSpells: 0,
			troops:           3,
			spells:           0,
		},
		{
			name:               "single hero only",
			link:               "https://link.clashofclans.com/en?action=CopyArmy&army=h0p4e8_14",
			heroes:             1,
			clanCastleTroops:   0,
			clanCastleSpells:   0,
			troops:             0,
			spells:             0,
			firstHeroPet:       true,
			firstHeroEquipment: 2,
		},
		{
			name:               "hero castle troop and spell",
			link:               "https://link.clashofclans.com/en?action=CopyArmy&army=h0p4e8_14i1x52d1x70",
			heroes:             1,
			clanCastleTroops:   1,
			clanCastleSpells:   1,
			troops:             0,
			spells:             0,
			firstHeroPet:       true,
			firstHeroEquipment: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recipe := clashy.ParseArmyRecipe(static, tt.link)
			if got := len(recipe.HeroesLoadout); got != tt.heroes {
				t.Fatalf("heroes = %d, want %d", got, tt.heroes)
			}
			if got := len(recipe.ClanCastleTroops); got != tt.clanCastleTroops {
				t.Fatalf("cc troops = %d, want %d", got, tt.clanCastleTroops)
			}
			if got := len(recipe.ClanCastleSpells); got != tt.clanCastleSpells {
				t.Fatalf("cc spells = %d, want %d", got, tt.clanCastleSpells)
			}
			if got := len(recipe.Troops); got != tt.troops {
				t.Fatalf("troops = %d, want %d", got, tt.troops)
			}
			if got := len(recipe.Spells); got != tt.spells {
				t.Fatalf("spells = %d, want %d", got, tt.spells)
			}

			if tt.heroes == 0 {
				return
			}

			first := recipe.HeroesLoadout[0]
			if got := first.Pet != nil; got != tt.firstHeroPet {
				t.Fatalf("first hero pet present = %v, want %v", got, tt.firstHeroPet)
			}
			if got := len(first.Equipment); got != tt.firstHeroEquipment {
				t.Fatalf("first hero equipment = %d, want %d", got, tt.firstHeroEquipment)
			}
		})
	}
}

func TestParseArmyRecipeHeroOptionalFields(t *testing.T) {
	static, err := clashy.LoadStaticData()
	if err != nil {
		t.Fatalf("load static data: %v", err)
	}

	tests := []struct {
		name      string
		payload   string
		wantPet   bool
		wantEquip int
	}{
		{name: "hero only", payload: "h0", wantPet: false, wantEquip: 0},
		{name: "hero pet only", payload: "h0p4", wantPet: true, wantEquip: 0},
		{name: "hero one equipment", payload: "h0e8", wantPet: false, wantEquip: 1},
		{name: "hero two equipment", payload: "h0e8_14", wantPet: false, wantEquip: 2},
		{name: "hero pet one equipment", payload: "h0p4e8", wantPet: true, wantEquip: 1},
		{name: "hero pet two equipment", payload: "h0p4e8_14", wantPet: true, wantEquip: 2},
		{name: "hero skin pet and equipment", payload: "h2m1p7e34_4", wantPet: true, wantEquip: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recipe := clashy.ParseArmyRecipe(static, "https://link.clashofclans.com/en?action=CopyArmy&army="+tt.payload)
			if got := len(recipe.HeroesLoadout); got != 1 {
				t.Fatalf("heroes = %d, want 1", got)
			}

			loadout := recipe.HeroesLoadout[0]
			if got := loadout.Pet != nil; got != tt.wantPet {
				t.Fatalf("pet present = %v, want %v", got, tt.wantPet)
			}
			if got := len(loadout.Equipment); got != tt.wantEquip {
				t.Fatalf("equipment = %d, want %d", got, tt.wantEquip)
			}
		})
	}
}

func TestClientParseArmyLinkRawShareCode(t *testing.T) {
	client, err := clashy.NewClient(clashy.DefaultClientConfig())
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	recipe := client.ParseArmyLink("u1x0-1x1-1x4")
	if got := len(recipe.Troops); got != 3 {
		t.Fatalf("troops = %d, want 3", got)
	}
	if got := len(recipe.HeroesLoadout); got != 0 {
		t.Fatalf("heroes = %d, want 0", got)
	}
	if got := recipe.Link; got != "u1x0-1x1-1x4" {
		t.Fatalf("link = %q, want raw share code", got)
	}
}

package clashy

var (
	TroopBaseID          = 4000000
	SpellBaseID          = 26000000
	HeroBaseID           = 2000000
	PetBaseID            = 60000000
	EquipmentBaseID      = 30000000
	ElixirTroopOrder     = []string{"Barbarian", "Archer", "Goblin", "Giant", "Wall Breaker", "Balloon", "Wizard", "Healer", "Dragon", "P.E.K.K.A", "Baby Dragon", "Miner", "Yeti", "Electro Dragon", "Dragon Rider", "Electro Titan", "Root Rider", "Thrower", "Meteor Golem"}
	DarkElixirTroopOrder = []string{"Minion", "Hog Rider", "Valkyrie", "Golem", "Witch", "Lava Hound", "Bowler", "Ice Golem", "Headhunter", "Apprentice Warden", "Druid", "Furnace"}
	HVTroopOrder         = append(append([]string{}, ElixirTroopOrder...), DarkElixirTroopOrder...)
	SiegeMachineOrder    = []string{"Wall Wrecker", "Battle Blimp", "Stone Slammer", "Siege Barracks", "Log Launcher", "Flame Flinger", "Battle Drill", "Troop Launcher"}
	HomeTroopOrder       = append(append([]string{}, HVTroopOrder...), SiegeMachineOrder...)
	BuilderTroopOrder    = []string{"Raged Barbarian", "Sneaky Archer", "Beta Minion", "Boxer Giant", "Bomber", "Power P.E.K.K.A", "Cannon Cart", "Drop Ship", "Baby Dragon", "Night Witch", "Hog Glider", "Electrofire Wizard"}
	SpellOrder           = []string{"Lightning Spell", "Healing Spell", "Rage Spell", "Jump Spell", "Freeze Spell", "Clone Spell", "Invisibility Spell", "Recall Spell", "Revive Spell", "Totem Spell", "Poison Spell", "Earthquake Spell", "Haste Spell", "Skeleton Spell", "Bat Spell", "Overgrowth Spell", "Ice Block Spell"}
	HeroOrder            = []string{"Barbarian King", "Archer Queen", "Minion Prince", "Grand Warden", "Royal Champion", "Battle Machine", "Battle Copter"}
	AchievementOrder     = []string{"Bigger Coffers", "Get those Goblins!", "Bigger & Better", "Nice and Tidy", "Release!", "Gold Grab", "Elixir Escapade", "Sweet Victory!", "Empire Builder", "Wall Buster", "Humiliator", "Union Buster", "Conqueror", "Unbreakable", "Friend in Need", "Mortar Mauler", "Heroic Heist", "League All-Star", "X-Bow Exterminator", "Firefighter", "War Hero", "Treasurer", "Anti-Artillery", "Sharing is caring", "Keep Your Account Safe!", "Master Engineering", "Next Generation Model", "Un-Build It", "Champion Builder", "High Gear", "Hidden Treasures", "Games Champion", "Dragon Slayer", "Well Seasoned", "Shattered and Scattered", "Not So Easy This Time", "Bust This!", "Superb Work", "Siege Sharer", "Aggressive Capitalism", "Most Valuable Clanmate", "Counterspell", "Monolith Masher", "Ungrateful Child", "Super Charger", "Pump It Up", "Astral Escape", "Not a Glass Cannon", "Tuurns Out This IS a Competition", "Heavy is the Crown", "No Heroics Allowed", "Fashion Victim", "The Floor is Healing", "Overgrowth and Overlord"}
)

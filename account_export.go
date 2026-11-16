package clashy

// AccountExport is the in-game export, not the official /v1 player response.
// Nil ownership slices mean not supplied. BuildingVariants lists unlocked
// variants and must never be interpreted as the currently selected variant.
type AccountExport struct {
	Tag              string `json:"tag"`
	Timestamp        int64  `json:"timestamp"`
	Emotes           []int  `json:"emotes,omitempty"`
	BuildingVariants []int  `json:"building_variants,omitempty"`
}

// IsSquad reports whether this troop is produced by the Squad Barracks.
func (t Troop) IsSquad(c *Client) bool {
	if c == nil {
		return false
	}
	village := t.Village
	if t.IsHomeBase() {
		village = "home"
	}
	item := c.StaticData().LookupByName(t.Name, "troops", village)
	return item != nil && item["production_building"] == "Squad Barracks"
}

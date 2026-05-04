package clashy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Client struct {
	config     ClientConfig
	http       *HTTPClient
	staticData *StaticData
}

func NewClient(cfg ClientConfig) (*Client, error) {
	cfg.BaseURL = normalizeAPIBase(cfg.BaseURL)
	cfg.DeveloperBaseURL = normalizeAPIBase(cfg.DeveloperBaseURL)
	if cfg.BaseURL == "" {
		cfg = DefaultClientConfig()
	}
	staticData, err := LoadStaticData()
	if err != nil {
		return nil, err
	}
	return &Client{
		config:     cfg,
		http:       NewHTTPClient(cfg),
		staticData: staticData,
	}, nil
}

func (c *Client) Login(ctx context.Context, email, password string) error {
	return c.http.LoginDeveloper(ctx, email, password)
}

func (c *Client) LoginWithTokens(_ context.Context, tokens ...string) error {
	c.http.SetTokens(tokens...)
	return nil
}

func (c *Client) Close() error { return nil }

func (c *Client) StaticData() *StaticData { return c.staticData }

func (c *Client) requestJSON(ctx context.Context, method, path string, body any, out any, opts RequestOptions) (int, error) {
	fullURL := c.config.BaseURL + path
	payload, _, retry, err := c.http.Do(ctx, method, fullURL, body, opts)
	if err != nil {
		return 0, err
	}
	if out == nil {
		return retry, nil
	}
	if err := json.Unmarshal(payload, out); err != nil {
		return 0, err
	}
	applyResponseMeta(out, retry)
	return retry, nil
}

func (c *Client) fetchItems(ctx context.Context, path string, out any) error {
	_, err := c.requestJSON(ctx, "GET", path, nil, out, c.defaultRequestOptions())
	return err
}

func (c *Client) defaultRequestOptions() RequestOptions {
	return RequestOptions{LookupCache: c.config.LookupCache, UpdateCache: c.config.UpdateCache}
}

type SearchClansRequest struct {
	Name          string
	WarFrequency  string
	LocationID    int
	MinMembers    int
	MaxMembers    int
	MinClanPoints int
	MinClanLevel  int
	LabelIDs      []int
	Limit         int
	Before        string
	After         string
}

func (r SearchClansRequest) values() url.Values {
	values := make(url.Values)
	if r.Name != "" {
		values.Set("name", r.Name)
	}
	if r.WarFrequency != "" {
		values.Set("warFrequency", r.WarFrequency)
	}
	if r.LocationID != 0 {
		values.Set("locationId", strconv.Itoa(r.LocationID))
	}
	if r.MinMembers != 0 {
		values.Set("minMembers", strconv.Itoa(r.MinMembers))
	}
	if r.MaxMembers != 0 {
		values.Set("maxMembers", strconv.Itoa(r.MaxMembers))
	}
	if r.MinClanPoints != 0 {
		values.Set("minClanPoints", strconv.Itoa(r.MinClanPoints))
	}
	if r.MinClanLevel != 0 {
		values.Set("minClanLevel", strconv.Itoa(r.MinClanLevel))
	}
	if len(r.LabelIDs) > 0 {
		var ids []string
		for _, id := range r.LabelIDs {
			ids = append(ids, strconv.Itoa(id))
		}
		values.Set("labelIds", strings.Join(ids, ","))
	}
	if r.Limit != 0 {
		values.Set("limit", strconv.Itoa(r.Limit))
	}
	if r.Before != "" {
		values.Set("before", r.Before)
	}
	if r.After != "" {
		values.Set("after", r.After)
	}
	return values
}

func (c *Client) SearchClans(ctx context.Context, req SearchClansRequest) ([]Clan, error) {
	var response struct {
		Items []Clan `json:"items"`
	}
	path := "/clans"
	if values := req.values().Encode(); values != "" {
		path += "?" + values
	}
	if err := c.fetchItems(ctx, path, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (c *Client) GetClan(ctx context.Context, tag string) (*Clan, error) {
	var clan Clan
	_, err := c.requestJSON(ctx, "GET", "/clans/"+encodeTag(tag), nil, &clan, c.defaultRequestOptions())
	if err != nil {
		return nil, err
	}
	return &clan, nil
}

func (c *Client) GetMembers(ctx context.Context, clanTag string, limit int, after, before string) ([]ClanMember, error) {
	var response struct {
		Items []ClanMember `json:"items"`
	}
	values := make(url.Values)
	if limit != 0 {
		values.Set("limit", strconv.Itoa(limit))
	}
	if after != "" {
		values.Set("after", after)
	}
	if before != "" {
		values.Set("before", before)
	}
	path := "/clans/" + encodeTag(clanTag) + "/members"
	if encoded := values.Encode(); encoded != "" {
		path += "?" + encoded
	}
	if err := c.fetchItems(ctx, path, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (c *Client) GetWarLog(ctx context.Context, clanTag string, limit int, after, before string) ([]ClanWarLogEntry, error) {
	var response struct {
		Items []ClanWarLogEntry `json:"items"`
	}
	values := make(url.Values)
	if limit != 0 {
		values.Set("limit", strconv.Itoa(limit))
	}
	if after != "" {
		values.Set("after", after)
	}
	if before != "" {
		values.Set("before", before)
	}
	path := "/clans/" + encodeTag(clanTag) + "/warlog"
	if encoded := values.Encode(); encoded != "" {
		path += "?" + encoded
	}
	if err := c.fetchItems(ctx, path, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (c *Client) GetRaidLog(ctx context.Context, clanTag string, limit int, after, before string) ([]RaidLogEntry, error) {
	var response struct {
		Items []RaidLogEntry `json:"items"`
	}
	values := make(url.Values)
	if limit != 0 {
		values.Set("limit", strconv.Itoa(limit))
	}
	if after != "" {
		values.Set("after", after)
	}
	if before != "" {
		values.Set("before", before)
	}
	path := "/clans/" + encodeTag(clanTag) + "/capitalraidseasons"
	if encoded := values.Encode(); encoded != "" {
		path += "?" + encoded
	}
	if err := c.fetchItems(ctx, path, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (c *Client) GetClanWar(ctx context.Context, clanTag string) (*ClanWar, error) {
	return c.getWar(ctx, "/clans/"+encodeTag(clanTag)+"/currentwar", clanTag)
}

func (c *Client) GetCurrentWar(ctx context.Context, clanTag string, round ...WarRound) (*ClanWar, error) {
	if len(round) == 0 || round[0] == CurrentWar {
		return c.getWar(ctx, encodeRealtime("/clans/"+encodeTag(clanTag)+"/currentwar", c.config.Realtime), clanTag)
	}
	if round[0] == CurrentPreparation {
		return c.GetLeagueWar(ctx, clanTag, round[0])
	}
	return c.getWar(ctx, encodeRealtime("/clans/"+encodeTag(clanTag)+"/currentwar", c.config.Realtime), clanTag)
}

func (c *Client) GetClanWars(ctx context.Context, tags []string) ([]ClanWar, error) {
	var out []ClanWar
	for _, tag := range tags {
		war, err := c.GetClanWar(ctx, tag)
		if err != nil {
			return nil, err
		}
		out = append(out, *war)
	}
	return out, nil
}

func (c *Client) GetLeagueGroup(ctx context.Context, clanTag string) (*ClanWarLeagueGroup, error) {
	var group ClanWarLeagueGroup
	_, err := c.requestJSON(ctx, "GET", encodeRealtime("/clans/"+encodeTag(clanTag)+"/currentwar/leaguegroup", c.config.Realtime), nil, &group, c.defaultRequestOptions())
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (c *Client) GetLeagueWar(ctx context.Context, clanTag string, round WarRound) (*ClanWar, error) {
	group, err := c.GetLeagueGroup(ctx, clanTag)
	if err != nil {
		return nil, err
	}
	if len(group.Rounds) == 0 {
		return nil, &NotFound{newHTTPException(404, "not found", "no league wars available", nil)}
	}
	index := len(group.Rounds) - 1
	if round == PreviousWar && len(group.Rounds) > 1 {
		index = len(group.Rounds) - 2
	}
	if round == CurrentPreparation {
		index = len(group.Rounds) - 1
	}
	for _, warTag := range group.Rounds[index].WarTags {
		if warTag == "" || warTag == "#0" {
			continue
		}
		war, err := c.getWar(ctx, encodeRealtime("/clanwarleagues/wars/"+encodeTag(warTag), c.config.Realtime), clanTag)
		if err == nil && war != nil && (war.Clan != nil && war.Clan.Tag == CorrectTag(clanTag) || war.Opponent != nil && war.Opponent.Tag == CorrectTag(clanTag)) {
			return war, nil
		}
	}
	return nil, &NotFound{newHTTPException(404, "not found", "league war not found for clan", nil)}
}

func (c *Client) GetLeagueWars(ctx context.Context, warTags []string) ([]ClanWar, error) {
	var out []ClanWar
	for _, warTag := range warTags {
		war, err := c.getWar(ctx, encodeRealtime("/clanwarleagues/wars/"+encodeTag(warTag), c.config.Realtime), "")
		if err != nil {
			return nil, err
		}
		out = append(out, *war)
	}
	return out, nil
}

func (c *Client) GetCurrentWars(ctx context.Context, tags []string) ([]ClanWar, error) {
	var out []ClanWar
	for _, tag := range tags {
		war, err := c.GetCurrentWar(ctx, tag)
		if err != nil {
			return nil, err
		}
		out = append(out, *war)
	}
	return out, nil
}

func (c *Client) getWar(ctx context.Context, path, clanTag string) (*ClanWar, error) {
	var war ClanWar
	_, err := c.requestJSON(ctx, "GET", path, nil, &war, c.defaultRequestOptions())
	if err != nil {
		return nil, err
	}
	war.ClanTag = CorrectTag(clanTag)
	if war.Clan != nil {
		for i := range war.Clan.Members {
			war.Clan.Members[i].Clan = war.Clan
		}
	}
	if war.Opponent != nil {
		for i := range war.Opponent.Members {
			war.Opponent.Members[i].Clan = war.Opponent
		}
	}
	return &war, nil
}

func (c *Client) SearchLocations(ctx context.Context, limit int, before, after string) ([]Location, error) {
	var response struct {
		Items []Location `json:"items"`
	}
	values := make(url.Values)
	if limit != 0 {
		values.Set("limit", strconv.Itoa(limit))
	}
	if before != "" {
		values.Set("before", before)
	}
	if after != "" {
		values.Set("after", after)
	}
	path := "/locations"
	if encoded := values.Encode(); encoded != "" {
		path += "?" + encoded
	}
	if err := c.fetchItems(ctx, path, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (c *Client) GetLocation(ctx context.Context, locationID int) (*Location, error) {
	var out Location
	_, err := c.requestJSON(ctx, "GET", fmt.Sprintf("/locations/%d", locationID), nil, &out, c.defaultRequestOptions())
	return &out, err
}

func (c *Client) GetLocationNamed(ctx context.Context, locationName string) (*Location, error) {
	locations, err := c.SearchLocations(ctx, 0, "", "")
	if err != nil {
		return nil, err
	}
	for _, location := range locations {
		if strings.EqualFold(location.Name, locationName) {
			return &location, nil
		}
	}
	return nil, nil
}

func (c *Client) getRankingItems(ctx context.Context, path string, limit int, before, after string, out any) error {
	values := make(url.Values)
	if limit != 0 {
		values.Set("limit", strconv.Itoa(limit))
	}
	if before != "" {
		values.Set("before", before)
	}
	if after != "" {
		values.Set("after", after)
	}
	if encoded := values.Encode(); encoded != "" {
		path += "?" + encoded
	}
	return c.fetchItems(ctx, path, out)
}

func (c *Client) GetLocationClans(ctx context.Context, locationID, limit int, before, after string) ([]RankedClan, error) {
	var response struct {
		Items []RankedClan `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/locations/%d/rankings/clans", locationID), limit, before, after, &response)
	return response.Items, err
}
func (c *Client) GetLocationClansCapital(ctx context.Context, locationID, limit int, before, after string) ([]RankedClan, error) {
	var response struct {
		Items []RankedClan `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/locations/%d/rankings/capitals", locationID), limit, before, after, &response)
	return response.Items, err
}
func (c *Client) GetLocationPlayers(ctx context.Context, locationID, limit int, before, after string) ([]RankedPlayer, error) {
	var response struct {
		Items []RankedPlayer `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/locations/%d/rankings/players", locationID), limit, before, after, &response)
	return response.Items, err
}
func (c *Client) GetLocationClansBuilderBase(ctx context.Context, locationID, limit int, before, after string) ([]RankedClan, error) {
	var response struct {
		Items []RankedClan `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/locations/%d/rankings/clans-builder-base", locationID), limit, before, after, &response)
	return response.Items, err
}
func (c *Client) GetLocationPlayersBuilderBase(ctx context.Context, locationID, limit int, before, after string) ([]RankedPlayer, error) {
	var response struct {
		Items []RankedPlayer `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/locations/%d/rankings/players-builder-base", locationID), limit, before, after, &response)
	return response.Items, err
}

func (c *Client) getLeagueItems(ctx context.Context, endpoint string, limit int, before, after string) ([]League, error) {
	var response struct {
		Items []League `json:"items"`
	}
	err := c.getRankingItems(ctx, endpoint, limit, before, after, &response)
	return response.Items, err
}

func (c *Client) SearchLeagues(ctx context.Context, limit int, before, after string) ([]League, error) {
	return c.getLeagueItems(ctx, "/leaguetiers", limit, before, after)
}
func (c *Client) SearchBuilderBaseLeagues(ctx context.Context, limit int, before, after string) ([]League, error) {
	return c.getLeagueItems(ctx, "/builderbaseleagues", limit, before, after)
}
func (c *Client) SearchWarLeagues(ctx context.Context, limit int, before, after string) ([]League, error) {
	return c.getLeagueItems(ctx, "/warleagues", limit, before, after)
}
func (c *Client) SearchCapitalLeagues(ctx context.Context, limit int, before, after string) ([]League, error) {
	return c.getLeagueItems(ctx, "/capitalleagues", limit, before, after)
}

func (c *Client) GetLeague(ctx context.Context, id int) (*League, error) {
	return c.getLeague(ctx, "/leaguetiers/"+strconv.Itoa(id))
}
func (c *Client) GetBuilderBaseLeague(ctx context.Context, id int) (*League, error) {
	return c.getLeague(ctx, "/builderbaseleagues/"+strconv.Itoa(id))
}
func (c *Client) GetWarLeague(ctx context.Context, id int) (*League, error) {
	return c.getLeague(ctx, "/warleagues/"+strconv.Itoa(id))
}
func (c *Client) GetCapitalLeague(ctx context.Context, id int) (*League, error) {
	return c.getLeague(ctx, "/capitalleagues/"+strconv.Itoa(id))
}

func (c *Client) getLeague(ctx context.Context, path string) (*League, error) {
	var out League
	_, err := c.requestJSON(ctx, "GET", path, nil, &out, c.defaultRequestOptions())
	return &out, err
}

func (c *Client) GetSeasons(ctx context.Context, leagueID int) ([]string, error) {
	var response struct {
		Items []struct {
			ID string `json:"id"`
		} `json:"items"`
	}
	if leagueID == 0 {
		leagueID = 29000022
	}
	if err := c.fetchItems(ctx, fmt.Sprintf("/leagues/%d/seasons", leagueID), &response); err != nil {
		return nil, err
	}
	out := make([]string, 0, len(response.Items))
	for _, item := range response.Items {
		out = append(out, item.ID)
	}
	return out, nil
}

func (c *Client) GetSeasonRankings(ctx context.Context, leagueID int, seasonID string) ([]RankedPlayer, error) {
	var response struct {
		Items []RankedPlayer `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/leagues/%d/seasons/%s", leagueID, seasonID), 0, "", "", &response)
	return response.Items, err
}

func (c *Client) GetClanLabels(ctx context.Context, limit int, before, after string) ([]Label, error) {
	var response struct {
		Items []Label `json:"items"`
	}
	err := c.getRankingItems(ctx, "/labels/clans", limit, before, after, &response)
	return response.Items, err
}

func (c *Client) GetPlayerLabels(ctx context.Context, limit int, before, after string) ([]Label, error) {
	var response struct {
		Items []Label `json:"items"`
	}
	err := c.getRankingItems(ctx, "/labels/players", limit, before, after, &response)
	return response.Items, err
}

func (c *Client) GetPlayer(ctx context.Context, tag string) (*Player, error) {
	var out Player
	_, err := c.requestJSON(ctx, "GET", "/players/"+encodeTag(tag), nil, &out, c.defaultRequestOptions())
	return &out, err
}

func (c *Client) GetBattleLog(ctx context.Context, playerTag string) ([]BattleLogEntry, error) {
	var response struct {
		Items []BattleLogEntry `json:"items"`
	}
	err := c.fetchItems(ctx, "/players/"+encodeTag(playerTag)+"/battlelog", &response)
	return response.Items, err
}

func (c *Client) GetPlayerLeagueHistory(ctx context.Context, playerTag string) ([]LeagueHistoryEntry, error) {
	var response struct {
		Items []LeagueHistoryEntry `json:"items"`
	}
	err := c.fetchItems(ctx, "/players/"+encodeTag(playerTag)+"/leaguehistory", &response)
	return response.Items, err
}

func (c *Client) GetPlayerLeagueGroup(ctx context.Context, playerTag, leagueGroupTag string, leagueSeasonID int) (*LeagueTierGroup, error) {
	var group LeagueTierGroup
	values := make(url.Values)
	values.Set("playerTag", CorrectTag(playerTag))
	path := "/leaguegroup/" + encodeTag(leagueGroupTag) + "/" + strconv.Itoa(leagueSeasonID)
	if encoded := values.Encode(); encoded != "" {
		path += "?" + encoded
	}
	_, err := c.requestJSON(ctx, "GET", path, nil, &group, c.defaultRequestOptions())
	return &group, err
}

func (c *Client) VerifyPlayerToken(ctx context.Context, playerTag, token string) (bool, error) {
	var response struct {
		Status string `json:"status"`
	}
	_, err := c.requestJSON(ctx, "POST", "/players/"+encodeTag(playerTag)+"/verifytoken", map[string]string{"token": token}, &response, RequestOptions{})
	if err != nil {
		return false, err
	}
	return response.Status == "ok", nil
}

func (c *Client) GetCurrentGoldPassSeason(ctx context.Context) (*GoldPassSeason, error) {
	var season GoldPassSeason
	_, err := c.requestJSON(ctx, "GET", "/goldpass/seasons/current", nil, &season, c.defaultRequestOptions())
	return &season, err
}

func (c *Client) ParseArmyLink(link string) ArmyRecipe             { return ParseArmyRecipe(c.staticData, link) }
func (c *Client) ParseAccountData(data map[string]any) AccountData { return ParseAccountData(data) }

func (c *Client) GetTroop(name string, isHomeVillage bool, level int) *Troop {
	if c == nil || c.staticData == nil {
		return nil
	}
	if level == 0 {
		level = 1
	}
	village := VillageBuilderBase
	if isHomeVillage {
		village = VillageHome
	}
	if item := c.staticData.LookupByName(name, "troops", string(village)); item != nil {
		troop := buildTroopFromStatic(item, level)
		troop.Name = firstNonEmpty(troop.Name, name)
		troop.Village = firstNonEmpty(troop.Village, string(village))
		return &troop
	}
	return nil
}

func (c *Client) GetSpell(name string, level int) *Spell {
	if c == nil || c.staticData == nil {
		return nil
	}
	if item := c.staticData.LookupByName(name, "spells", ""); item != nil {
		spell := buildSpellFromStatic(item, level)
		spell.Name = firstNonEmpty(spell.Name, name)
		return &spell
	}
	return nil
}

func (c *Client) GetHero(name string, level int) *Hero {
	if c == nil || c.staticData == nil {
		return nil
	}
	if item := c.staticData.LookupByName(name, "heroes", ""); item != nil {
		hero := buildHeroFromStatic(item, level)
		hero.Name = firstNonEmpty(hero.Name, name)
		return &hero
	}
	return nil
}

func (c *Client) GetPet(name string, level int) *Pet {
	if c == nil || c.staticData == nil {
		return nil
	}
	if item := c.staticData.LookupByName(name, "pets", ""); item != nil {
		pet := buildPetFromStatic(item, level)
		pet.Name = firstNonEmpty(pet.Name, name)
		return &pet
	}
	return nil
}

func (c *Client) GetEquipment(name string, level int) *Equipment {
	if c == nil || c.staticData == nil {
		return nil
	}
	if item := c.staticData.LookupByName(name, "equipment", ""); item != nil {
		equipment := buildEquipmentFromStatic(item, level)
		equipment.Name = firstNonEmpty(equipment.Name, name)
		return &equipment
	}
	return nil
}

func (c *Client) GetTranslation(id string) *Translation {
	if c.staticData == nil || c.staticData.Translations == nil {
		return nil
	}
	languages := c.staticData.Translations[id]
	if len(languages) == 0 {
		return nil
	}
	return &Translation{ID: id, English: languages["EN"], Languages: languages}
}

func (c *Client) GetExtendedCWLGroupData(name string) *ExtendedCWLGroup {
	item := c.staticData.LookupByName(name, "war_leagues", "")
	if item == nil {
		return nil
	}
	var group ExtendedCWLGroup
	data, _ := json.Marshal(item)
	_ = json.Unmarshal(data, &group)
	if group.Name == "" {
		group.Name = name
	}
	return &group
}

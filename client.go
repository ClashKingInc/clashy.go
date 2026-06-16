package clashy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Client is the high-level Clash API client.
//
// A Client owns its configuration, HTTP transport, and embedded static-data
// indexes. It is safe to reuse a single client across request handlers as long
// as callers pass appropriate contexts.
type Client struct {
	config     ClientConfig
	http       *HTTPClient
	staticData *StaticData
}

// NewClient constructs a Client from cfg and loads embedded static data.
//
// If cfg.BaseURL is empty, DefaultClientConfig is used. BaseURL and
// DeveloperBaseURL are normalized by removing trailing slashes.
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

// Login authenticates with developer-site email and password credentials.
//
// The developer login flow discovers or creates API keys, stores them in the
// underlying HTTP client, and uses those keys for later Clash API requests.
func (c *Client) Login(ctx context.Context, email, password string) error {
	return c.http.LoginDeveloper(ctx, email, password)
}

// LoginWithTokens configures one or more existing Clash API tokens.
//
// Tokens are rotated by the underlying HTTP client. The context parameter is
// accepted for API symmetry with Login.
func (c *Client) LoginWithTokens(_ context.Context, tokens ...string) error {
	c.http.SetTokens(tokens...)
	return nil
}

// Close releases client resources.
//
// The current implementation does not hold resources that need explicit
// teardown, so Close returns nil.
func (c *Client) Close() error { return nil }

// StaticData returns the client's embedded static-data index.
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

// SearchClansRequest contains optional filters for SearchClans.
//
// Zero values are omitted from the query string, matching Clash API search
// behavior.
type SearchClansRequest struct {
	// Name filters clans by name.
	Name string
	// WarFrequency filters clans by declared war frequency.
	WarFrequency string
	// LocationID filters clans by location ID.
	LocationID int
	// MinMembers filters out clans with fewer members.
	MinMembers int
	// MaxMembers filters out clans with more members.
	MaxMembers int
	// MinClanPoints filters by minimum clan points.
	MinClanPoints int
	// MinClanLevel filters by minimum clan level.
	MinClanLevel int
	// LabelIDs filters by one or more clan label IDs.
	LabelIDs []int
	// Limit controls the number of results requested.
	Limit int
	// Before is a pagination cursor.
	Before string
	// After is a pagination cursor.
	After string
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

// SearchClans searches clans using the provided optional filters.
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

// GetClan fetches a clan profile by tag.
func (c *Client) GetClan(ctx context.Context, tag string) (*Clan, error) {
	var clan Clan
	_, err := c.requestJSON(ctx, "GET", "/clans/"+encodeTag(tag), nil, &clan, c.defaultRequestOptions())
	if err != nil {
		return nil, err
	}
	return &clan, nil
}

// GetMembers fetches a clan member page by clan tag.
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

// GetWarLog fetches public war log entries for a clan.
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

// GetRaidLog fetches Clan Capital raid weekend log entries for a clan.
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

// GetClanWar fetches the regular current-war endpoint for a clan.
//
// This method does not fall back to CWL. Use GetCurrentWar when you want the
// active normal war or the relevant Clan War League war.
func (c *Client) GetClanWar(ctx context.Context, clanTag string) (*ClanWar, error) {
	return c.getWar(ctx, encodeRealtime("/clans/"+encodeTag(clanTag)+"/currentwar", c.config.Realtime), clanTag)
}

// GetCurrentWar returns the clan's active normal war or relevant CWL war.
//
// The method first checks the regular current-war endpoint. If the clan is not
// in a regular war, or the war log is private, it loads the CWL group and
// returns the selected league round for the clan. Passing no round selects
// CurrentWar. When no current war exists, the method returns nil, nil.
func (c *Client) GetCurrentWar(ctx context.Context, clanTag string, round ...WarRound) (*ClanWar, error) {
	cwlRound := CurrentWar
	if len(round) > 0 {
		cwlRound = round[0]
	}

	// The regular current-war endpoint is authoritative for non-CWL wars.
	// Only fall through to CWL lookup when it explicitly says notInWar, or
	// when a private war log prevents us from seeing the regular war state.
	regularWar, regularErr := c.GetClanWar(ctx, clanTag)
	if regularErr != nil {
		var privateWarLog *PrivateWarLog
		if !errors.As(regularErr, &privateWarLog) {
			return nil, regularErr
		}
	} else if regularWar != nil && regularWar.State != WarStateNotInWar {
		return regularWar, nil
	}

	group, err := c.GetLeagueGroup(ctx, clanTag)
	if err != nil {
		var notFound *NotFound
		var gateway *GatewayError
		if errors.As(err, &notFound) || errors.As(err, &gateway) {
			if regularWar != nil {
				return regularWar, nil
			}
			return nil, regularErr
		}
		return nil, err
	}
	if group.State == "notInWar" || group.State == "groupNotFound" {
		return nil, nil
	}

	// A CWL group round contains every matchup in that round, not just this clan.
	// Select the logical round first, then scan that round until this clan's war is found.
	warTags, ok, err := c.selectLeagueRound(ctx, clanTag, group, cwlRound)
	if err != nil || !ok {
		return nil, err
	}
	return c.findClanLeagueWar(ctx, clanTag, group, warTags)
}

// GetClanWars fetches the regular current war for each clan tag in order.
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

// GetLeagueGroup fetches the current Clan War League group for a clan.
func (c *Client) GetLeagueGroup(ctx context.Context, clanTag string) (*ClanWarLeagueGroup, error) {
	var group ClanWarLeagueGroup
	_, err := c.requestJSON(ctx, "GET", "/clans/"+encodeTag(clanTag)+"/currentwar/leaguegroup", nil, &group, c.defaultRequestOptions())
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// GetLeagueWar fetches the selected CWL round for a clan.
//
// The returned war is oriented so Clan is the requested clan and Opponent is the
// opposing side.
func (c *Client) GetLeagueWar(ctx context.Context, clanTag string, round WarRound) (*ClanWar, error) {
	group, err := c.GetLeagueGroup(ctx, clanTag)
	if err != nil {
		return nil, err
	}
	warTags, ok, err := c.selectLeagueRound(ctx, clanTag, group, round)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, &NotFound{newHTTPException(404, "not found", "no league wars available", nil)}
	}
	war, err := c.findClanLeagueWar(ctx, clanTag, group, warTags)
	if err != nil {
		return nil, err
	}
	if war != nil {
		return war, nil
	}
	return nil, &NotFound{newHTTPException(404, "not found", "league war not found for clan", nil)}
}

// GetLeagueWars fetches CWL wars by war tag.
func (c *Client) GetLeagueWars(ctx context.Context, warTags []string) ([]ClanWar, error) {
	var out []ClanWar
	for _, warTag := range warTags {
		war, err := c.getLeagueWarByTag(ctx, warTag, nil, "")
		if err != nil {
			return nil, err
		}
		out = append(out, *war)
	}
	return out, nil
}

// GetCurrentWars fetches GetCurrentWar for each clan tag and omits clans with no
// current war.
func (c *Client) GetCurrentWars(ctx context.Context, tags []string) ([]ClanWar, error) {
	var out []ClanWar
	for _, tag := range tags {
		war, err := c.GetCurrentWar(ctx, tag)
		if err != nil {
			return nil, err
		}
		if war == nil {
			continue
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
	return &war, nil
}

func (c *Client) getLeagueWarByTag(ctx context.Context, warTag string, group *ClanWarLeagueGroup, clanTag string) (*ClanWar, error) {
	war, err := c.getWar(ctx, "/clanwarleagues/wars/"+encodeTag(warTag), clanTag)
	if err != nil {
		return nil, err
	}
	war.WarTag = CorrectTag(warTag)
	war.LeagueGroup = group
	return war, nil
}

func (c *Client) findClanLeagueWar(ctx context.Context, clanTag string, group *ClanWarLeagueGroup, warTags []string) (*ClanWar, error) {
	// A CWL round has every matchup in the group, so scan the selected round
	// until the requested clan is found on either side of a war.
	for _, warTag := range warTags {
		war, err := c.getLeagueWarByTag(ctx, warTag, group, clanTag)
		if err != nil {
			return nil, err
		}
		if orientWarForClan(war, clanTag) {
			return war, nil
		}
	}
	return nil, nil
}

// selectLeagueRound maps a requested WarRound onto the API's CWL group shape.
// The group contains all known rounds, while future rounds are usually "#0".
// When the latest real round is present, one league-war request is enough to
// tell whether that latest round is the active war or only the next preparation.
func (c *Client) selectLeagueRound(ctx context.Context, clanTag string, group *ClanWarLeagueGroup, round WarRound) ([]string, bool, error) {
	rounds := validLeagueRounds(group)
	if len(rounds) == 0 {
		return nil, false, nil
	}
	latestRoundPreparing := group.State == "preparation"
	if len(group.Rounds) == len(rounds) && group.State != "ended" {
		// If all API rounds are real war tags, the latest round might be either
		// already in war or still in preparation. All wars in a CWL round share
		// state, so probing the first matchup is enough to classify the round.
		warTag := rounds[len(rounds)-1][0]
		war, err := c.getLeagueWarByTag(ctx, warTag, group, clanTag)
		if err != nil {
			return nil, false, err
		}
		if war.State == WarStateInWar {
			latestRoundPreparing = false
		}
		if war.State == WarStatePreparation {
			latestRoundPreparing = true
		}
	}
	if round == CurrentPreparation && group.State == "ended" {
		return nil, false, nil
	}
	if round == CurrentWar {
		// During first-round preparation there is no previous round, so the prep
		// war is the best current CWL war. Between later rounds, "current war"
		// means the last completed/in-war round rather than the upcoming prep.
		if latestRoundPreparing && len(rounds) > 1 {
			return rounds[len(rounds)-2], true, nil
		}
		return rounds[len(rounds)-1], true, nil
	}
	if round == CurrentPreparation {
		// Only expose current preparation when the latest round is actually in
		// preparation; otherwise there is no upcoming prep war to return.
		if latestRoundPreparing {
			return rounds[len(rounds)-1], true, nil
		}
		return nil, false, nil
	}
	if round == PreviousWar {
		// If the newest round is only preparation, step back past it. Otherwise
		// the previous war is simply the round before the selected current round.
		if len(rounds) < 2 {
			return nil, false, nil
		}
		if latestRoundPreparing && len(rounds) > 2 {
			return rounds[len(rounds)-3], true, nil
		}
		return rounds[len(rounds)-2], true, nil
	}
	return nil, false, nil
}

// validLeagueRounds strips future "#0" placeholders and keeps each real round's
// matchup tags intact. A returned inner slice can still have many war tags.
func validLeagueRounds(group *ClanWarLeagueGroup) [][]string {
	if group == nil {
		return nil
	}
	rounds := make([][]string, 0, len(group.Rounds))
	for _, round := range group.Rounds {
		tags := make([]string, 0, len(round.WarTags))
		for _, warTag := range round.WarTags {
			if warTag != "" && warTag != "#0" {
				tags = append(tags, warTag)
			}
		}
		if len(tags) > 0 {
			rounds = append(rounds, tags)
		}
	}
	return rounds
}

func orientWarForClan(war *ClanWar, clanTag string) bool {
	// League-war responses are matchup-centric. If the requested clan is on the
	// opponent side, flip the sides so callers can always read war.Clan as self.
	if war == nil {
		return false
	}
	clanTag = CorrectTag(clanTag)
	if war.Clan != nil && CorrectTag(war.Clan.Tag) == clanTag {
		war.ClanTag = clanTag
		return true
	}
	if war.Opponent != nil && CorrectTag(war.Opponent.Tag) == clanTag {
		war.Clan, war.Opponent = war.Opponent, war.Clan
		war.ClanTag = clanTag
		return true
	}
	return false
}

// SearchLocations fetches API locations with optional pagination.
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

// GetLocation fetches a location by numeric ID.
func (c *Client) GetLocation(ctx context.Context, locationID int) (*Location, error) {
	var out Location
	_, err := c.requestJSON(ctx, "GET", fmt.Sprintf("/locations/%d", locationID), nil, &out, c.defaultRequestOptions())
	return &out, err
}

// GetLocationNamed returns the first location whose name matches
// locationName case-insensitively.
//
// It returns nil, nil when no matching location is found.
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

// GetLocationClans fetches home-village clan rankings for a numeric location ID.
func (c *Client) GetLocationClans(ctx context.Context, locationID, limit int, before, after string) ([]RankedClan, error) {
	return c.GetLocationClansByLocationID(ctx, strconv.Itoa(locationID), limit, before, after)
}

// GetLocationClansByLocationID fetches home-village clan rankings for a
// location ID string.
func (c *Client) GetLocationClansByLocationID(ctx context.Context, locationID string, limit int, before, after string) ([]RankedClan, error) {
	var response struct {
		Items []RankedClan `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/locations/%s/rankings/clans", locationID), limit, before, after, &response)
	return response.Items, err
}

// GetLocationClansCapital fetches Clan Capital clan rankings for a numeric
// location ID.
func (c *Client) GetLocationClansCapital(ctx context.Context, locationID, limit int, before, after string) ([]RankedClan, error) {
	return c.GetLocationClansCapitalByLocationID(ctx, strconv.Itoa(locationID), limit, before, after)
}

// GetLocationClansCapitalByLocationID fetches Clan Capital clan rankings for a
// location ID string.
func (c *Client) GetLocationClansCapitalByLocationID(ctx context.Context, locationID string, limit int, before, after string) ([]RankedClan, error) {
	var response struct {
		Items []RankedClan `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/locations/%s/rankings/capitals", locationID), limit, before, after, &response)
	return response.Items, err
}

// GetLocationPlayers fetches home-village player rankings for a numeric
// location ID.
func (c *Client) GetLocationPlayers(ctx context.Context, locationID, limit int, before, after string) ([]RankedPlayer, error) {
	return c.GetLocationPlayersByLocationID(ctx, strconv.Itoa(locationID), limit, before, after)
}

// GetLocationPlayersByLocationID fetches home-village player rankings for a
// location ID string.
func (c *Client) GetLocationPlayersByLocationID(ctx context.Context, locationID string, limit int, before, after string) ([]RankedPlayer, error) {
	var response struct {
		Items []RankedPlayer `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/locations/%s/rankings/players", locationID), limit, before, after, &response)
	return response.Items, err
}

// GetLocationClansBuilderBase fetches Builder Base clan rankings for a numeric
// location ID.
func (c *Client) GetLocationClansBuilderBase(ctx context.Context, locationID, limit int, before, after string) ([]RankedClan, error) {
	return c.GetLocationClansBuilderBaseByLocationID(ctx, strconv.Itoa(locationID), limit, before, after)
}

// GetLocationClansBuilderBaseByLocationID fetches Builder Base clan rankings
// for a location ID string.
func (c *Client) GetLocationClansBuilderBaseByLocationID(ctx context.Context, locationID string, limit int, before, after string) ([]RankedClan, error) {
	var response struct {
		Items []RankedClan `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/locations/%s/rankings/clans-builder-base", locationID), limit, before, after, &response)
	return response.Items, err
}

// GetLocationPlayersBuilderBase fetches Builder Base player rankings for a
// numeric location ID.
func (c *Client) GetLocationPlayersBuilderBase(ctx context.Context, locationID, limit int, before, after string) ([]RankedPlayer, error) {
	return c.GetLocationPlayersBuilderBaseByLocationID(ctx, strconv.Itoa(locationID), limit, before, after)
}

// GetLocationPlayersBuilderBaseByLocationID fetches Builder Base player
// rankings for a location ID string.
func (c *Client) GetLocationPlayersBuilderBaseByLocationID(ctx context.Context, locationID string, limit int, before, after string) ([]RankedPlayer, error) {
	var response struct {
		Items []RankedPlayer `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/locations/%s/rankings/players-builder-base", locationID), limit, before, after, &response)
	return response.Items, err
}

func (c *Client) getLeagueItems(ctx context.Context, endpoint string, limit int, before, after string) ([]League, error) {
	var response struct {
		Items []League `json:"items"`
	}
	err := c.getRankingItems(ctx, endpoint, limit, before, after, &response)
	return response.Items, err
}

// SearchLeagues fetches home-village leagues with optional pagination.
func (c *Client) SearchLeagues(ctx context.Context, limit int, before, after string) ([]League, error) {
	return c.getLeagueItems(ctx, "/leaguetiers", limit, before, after)
}

// SearchBuilderBaseLeagues fetches Builder Base leagues with optional
// pagination.
func (c *Client) SearchBuilderBaseLeagues(ctx context.Context, limit int, before, after string) ([]League, error) {
	return c.getLeagueItems(ctx, "/builderbaseleagues", limit, before, after)
}

// SearchWarLeagues fetches Clan War League tiers with optional pagination.
func (c *Client) SearchWarLeagues(ctx context.Context, limit int, before, after string) ([]League, error) {
	return c.getLeagueItems(ctx, "/warleagues", limit, before, after)
}

// SearchCapitalLeagues fetches Clan Capital leagues with optional pagination.
func (c *Client) SearchCapitalLeagues(ctx context.Context, limit int, before, after string) ([]League, error) {
	return c.getLeagueItems(ctx, "/capitalleagues", limit, before, after)
}

// GetLeague fetches a home-village league by ID.
func (c *Client) GetLeague(ctx context.Context, id int) (*League, error) {
	return c.getLeague(ctx, "/leaguetiers/"+strconv.Itoa(id))
}

// GetBuilderBaseLeague fetches a Builder Base league by ID.
func (c *Client) GetBuilderBaseLeague(ctx context.Context, id int) (*League, error) {
	return c.getLeague(ctx, "/builderbaseleagues/"+strconv.Itoa(id))
}

// GetWarLeague fetches a Clan War League tier by ID.
func (c *Client) GetWarLeague(ctx context.Context, id int) (*League, error) {
	return c.getLeague(ctx, "/warleagues/"+strconv.Itoa(id))
}

// GetCapitalLeague fetches a Clan Capital league by ID.
func (c *Client) GetCapitalLeague(ctx context.Context, id int) (*League, error) {
	return c.getLeague(ctx, "/capitalleagues/"+strconv.Itoa(id))
}

func (c *Client) getLeague(ctx context.Context, path string) (*League, error) {
	var out League
	_, err := c.requestJSON(ctx, "GET", path, nil, &out, c.defaultRequestOptions())
	return &out, err
}

// GetSeasons fetches available season IDs for a league.
//
// Passing leagueID 0 uses the default legend league ID.
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

// GetSeasonRankings fetches player rankings for a league season.
func (c *Client) GetSeasonRankings(ctx context.Context, leagueID int, seasonID string) ([]RankedPlayer, error) {
	var response struct {
		Items []RankedPlayer `json:"items"`
	}
	err := c.getRankingItems(ctx, fmt.Sprintf("/leagues/%d/seasons/%s", leagueID, url.PathEscape(seasonID)), 0, "", "", &response)
	return response.Items, err
}

// GetClanLabels fetches clan labels with optional pagination.
func (c *Client) GetClanLabels(ctx context.Context, limit int, before, after string) ([]Label, error) {
	var response struct {
		Items []Label `json:"items"`
	}
	err := c.getRankingItems(ctx, "/labels/clans", limit, before, after, &response)
	return response.Items, err
}

// GetPlayerLabels fetches player labels with optional pagination.
func (c *Client) GetPlayerLabels(ctx context.Context, limit int, before, after string) ([]Label, error) {
	var response struct {
		Items []Label `json:"items"`
	}
	err := c.getRankingItems(ctx, "/labels/players", limit, before, after, &response)
	return response.Items, err
}

// GetPlayer fetches a player profile by tag.
func (c *Client) GetPlayer(ctx context.Context, tag string) (*Player, error) {
	var out Player
	_, err := c.requestJSON(ctx, "GET", "/players/"+encodeTag(tag), nil, &out, c.defaultRequestOptions())
	return &out, err
}

// GetBattleLog fetches a player's battle log.
func (c *Client) GetBattleLog(ctx context.Context, playerTag string) ([]BattleLogEntry, error) {
	var response struct {
		Items []BattleLogEntry `json:"items"`
	}
	err := c.fetchItems(ctx, "/players/"+encodeTag(playerTag)+"/battlelog", &response)
	return response.Items, err
}

// GetPlayerLeagueHistory fetches a player's legend league history.
func (c *Client) GetPlayerLeagueHistory(ctx context.Context, playerTag string) ([]LeagueHistoryEntry, error) {
	var response struct {
		Items []LeagueHistoryEntry `json:"items"`
	}
	err := c.fetchItems(ctx, "/players/"+encodeTag(playerTag)+"/leaguehistory", &response)
	return response.Items, err
}

// GetPlayerLeagueGroup fetches a ranked group and scopes it to a player.
func (c *Client) GetPlayerLeagueGroup(ctx context.Context, playerTag, leagueGroupTag, leagueSeasonID string) (*LeagueTierGroup, error) {
	var group LeagueTierGroup
	values := make(url.Values)
	values.Set("playerTag", CorrectTag(playerTag))
	path := "/leaguegroup/" + encodeTag(leagueGroupTag) + "/" + url.PathEscape(leagueSeasonID)
	if encoded := values.Encode(); encoded != "" {
		path += "?" + encoded
	}
	_, err := c.requestJSON(ctx, "GET", path, nil, &group, c.defaultRequestOptions())
	return &group, err
}

// VerifyPlayerToken verifies an in-game player API token.
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

// GetCurrentGoldPassSeason fetches the current Gold Pass season.
func (c *Client) GetCurrentGoldPassSeason(ctx context.Context) (*GoldPassSeason, error) {
	var season GoldPassSeason
	_, err := c.requestJSON(ctx, "GET", "/goldpass/seasons/current", nil, &season, c.defaultRequestOptions())
	return &season, err
}

// ParseArmyLink parses a full Clash army link or raw army payload using the
// client's static data.
func (c *Client) ParseArmyLink(link string) ArmyRecipe { return ParseArmyRecipe(c.staticData, link) }

// ParseAccountData wraps arbitrary account-link data without mutating it.
func (c *Client) ParseAccountData(data map[string]any) AccountData { return ParseAccountData(data) }

// GetTroop looks up a troop by name, village, and level in embedded static data.
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

// GetSpell looks up a spell by name and level in embedded static data.
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

// GetHero looks up a hero by name and level in embedded static data.
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

// GetPet looks up a pet by name and level in embedded static data.
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

// GetEquipment looks up hero equipment by name and level in embedded static
// data.
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

// GetTranslation returns a translation entry by static-data translation ID.
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

// GetExtendedCWLGroupData returns static medal data for a Clan War League tier
// by name.
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

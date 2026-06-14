# Client

Client construction, authentication, request behavior, and high-level Clash API methods.

<a id="client"></a>

## Client

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Client</code></p>

Client is the high-level Clash API client.

A Client owns its configuration, HTTP transport, and embedded static-data
indexes. It is safe to reuse a single client across request handlers as long
as callers pass appropriate contexts.

<a id="clientconfig"></a>

## Client Config

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ClientConfig</code></p>

ClientConfig controls API endpoints, authentication behavior, request
throttling, response caching, and static-data loading.

Most callers should start from DefaultClientConfig and override only the
fields that differ for their service. The zero value is not the recommended
production configuration because it has no API base URL or timeout.

<div class="api-field" id="clientconfig-keycount" markdown="1">

### `KeyCount`

<p><code>int</code></p>

KeyCount is the number of developer-site API keys Login should make
available for token rotation.

</div>

<div class="api-field" id="clientconfig-keynames" markdown="1">

### `KeyNames`

<p><code>string</code></p>

KeyNames is the developer-site key name used when reusing or creating API
keys during Login.

</div>

<div class="api-field" id="clientconfig-throttlelimit" markdown="1">

### `ThrottleLimit`

<p><code>int</code></p>

ThrottleLimit is the maximum number of concurrent HTTP requests allowed by
the client. A value less than or equal to zero disables the limiter.

</div>

<div class="api-field" id="clientconfig-timeout" markdown="1">

### `Timeout`

<p><code>time.Duration</code></p>

Timeout is applied to the underlying http.Client.

</div>

<div class="api-field" id="clientconfig-baseurl" markdown="1">

### `BaseURL`

<p><code>string</code></p>

BaseURL is the Clash API or compatible proxy base URL, usually ending in
/v1 without a trailing slash.

</div>

<div class="api-field" id="clientconfig-developerbaseurl" markdown="1">

### `DeveloperBaseURL`

<p><code>string</code></p>

DeveloperBaseURL is the developer-site base URL used only by Login.

</div>

<div class="api-field" id="clientconfig-ip" markdown="1">

### `IP`

<p><code>string</code></p>

IP overrides the IP address used when Login creates developer-site API
keys. When empty, the IP is inferred from the temporary developer token.

</div>

<div class="api-field" id="clientconfig-realtime" markdown="1">

### `Realtime`

<p><code>bool</code></p>

Realtime adds realtime=true to current-war requests that support it.

</div>

<div class="api-field" id="clientconfig-correcttags" markdown="1">

### `CorrectTags`

<p><code>bool</code></p>

CorrectTags enables Clash tag normalization before tags are placed in API
paths or query strings.

</div>

<div class="api-field" id="clientconfig-cachemaxsize" markdown="1">

### `CacheMaxSize`

<p><code>int</code></p>

CacheMaxSize bounds the number of GET responses retained in memory.

</div>

<div class="api-field" id="clientconfig-lookupcache" markdown="1">

### `LookupCache`

<p><code>bool</code></p>

LookupCache allows GET requests to return fresh cached responses.

</div>

<div class="api-field" id="clientconfig-updatecache" markdown="1">

### `UpdateCache`

<p><code>bool</code></p>

UpdateCache allows successful GET responses to refresh the in-memory cache.

</div>

<div class="api-field" id="clientconfig-ignorecachederrors" markdown="1">

### `IgnoreCachedErrors`

<p><code>[]int</code></p>

IgnoreCachedErrors is reserved for compatibility with callers that model
cache behavior after coc.py; current request handling does not use it.

</div>

<div class="api-field" id="clientconfig-rawjson" markdown="1">

### `RawJSON`

<p><code>bool</code></p>

RawJSON is reserved for callers that need raw response capture; current
high-level methods unmarshal into typed models.

</div>

<div class="api-field" id="clientconfig-loadgamedata" markdown="1">

### `LoadGameData`

<p><code><a href="../enums/#loadgamedata">LoadGameData</a></code></p>

LoadGameData describes when static game data should be loaded.

</div>

<div class="api-field" id="clientconfig-useragent" markdown="1">

### `UserAgent`

<p><code>string</code></p>

UserAgent is sent with Clash API or proxy requests.

</div>

<div class="api-field" id="clientconfig-developeruseragent" markdown="1">

### `DeveloperUserAgent`

<p><code>string</code></p>

DeveloperUserAgent is sent with developer-site login and key-management
requests.

</div>

<a id="searchclansrequest"></a>

## Search Clans Request

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.SearchClansRequest</code></p>

SearchClansRequest contains optional filters for SearchClans.

Zero values are omitted from the query string, matching Clash API search
behavior.

<div class="api-field" id="searchclansrequest-name" markdown="1">

### `Name`

<p><code>string</code></p>

Name filters clans by name.

</div>

<div class="api-field" id="searchclansrequest-warfrequency" markdown="1">

### `WarFrequency`

<p><code>string</code></p>

WarFrequency filters clans by declared war frequency.

</div>

<div class="api-field" id="searchclansrequest-locationid" markdown="1">

### `LocationID`

<p><code>int</code></p>

LocationID filters clans by location ID.

</div>

<div class="api-field" id="searchclansrequest-minmembers" markdown="1">

### `MinMembers`

<p><code>int</code></p>

MinMembers filters out clans with fewer members.

</div>

<div class="api-field" id="searchclansrequest-maxmembers" markdown="1">

### `MaxMembers`

<p><code>int</code></p>

MaxMembers filters out clans with more members.

</div>

<div class="api-field" id="searchclansrequest-minclanpoints" markdown="1">

### `MinClanPoints`

<p><code>int</code></p>

MinClanPoints filters by minimum clan points.

</div>

<div class="api-field" id="searchclansrequest-minclanlevel" markdown="1">

### `MinClanLevel`

<p><code>int</code></p>

MinClanLevel filters by minimum clan level.

</div>

<div class="api-field" id="searchclansrequest-labelids" markdown="1">

### `LabelIDs`

<p><code>[]int</code></p>

LabelIDs filters by one or more clan label IDs.

</div>

<div class="api-field" id="searchclansrequest-limit" markdown="1">

### `Limit`

<p><code>int</code></p>

Limit controls the number of results requested.

</div>

<div class="api-field" id="searchclansrequest-before" markdown="1">

### `Before`

<p><code>string</code></p>

Before is a pagination cursor.

</div>

<div class="api-field" id="searchclansrequest-after" markdown="1">

### `After`

<p><code>string</code></p>

After is a pagination cursor.

</div>

<a id="requestoptions"></a>

## Request Options

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.RequestOptions</code></p>

RequestOptions controls per-request behavior for HTTPClient.Do.

<div class="api-field" id="requestoptions-lookupcache" markdown="1">

### `LookupCache`

<p><code>bool</code></p>

LookupCache allows a GET request to return a fresh cached response.

</div>

<div class="api-field" id="requestoptions-updatecache" markdown="1">

### `UpdateCache`

<p><code>bool</code></p>

UpdateCache allows a successful GET request to store or replace a cached
response.

</div>

<div class="api-field" id="requestoptions-skipauth" markdown="1">

### `SkipAuth`

<p><code>bool</code></p>

SkipAuth prevents Do from adding an Authorization header.

</div>

<a id="httpclient"></a>

## HTTPClient

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.HTTPClient</code></p>

HTTPClient performs low-level Clash API and developer-site HTTP requests.

Most callers should use Client instead. HTTPClient is exported so advanced
integrations can build compatible request flows while reusing token rotation,
throttling, compression handling, cache storage, and typed error mapping.

## Client Methods

<a id="client-close"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.Close()<span class="api-return-arrow"> -> </span><span class="api-return">error</span></code></p>

Close releases client resources.

The current implementation does not hold resources that need explicit
teardown, so Close returns nil.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>error</code> </dd>
</dl>

</div>

<a id="client-getbattlelog"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetBattleLog(<span class="api-param">ctx: context.Context</span>, <span class="api-param">playerTag: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]BattleLogEntry</span>, <span class="api-return">error</span>)</code></p>

GetBattleLog fetches a player's battle log.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>playerTag</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../battle-logs/#battlelogentry">BattleLogEntry</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getbuilderbaseleague"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetBuilderBaseLeague(<span class="api-param">ctx: context.Context</span>, <span class="api-param">id: int</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*League</span>, <span class="api-return">error</span>)</code></p>

GetBuilderBaseLeague fetches a Builder Base league by ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>id</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../locations-rankings/#league">League</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getcapitalleague"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetCapitalLeague(<span class="api-param">ctx: context.Context</span>, <span class="api-param">id: int</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*League</span>, <span class="api-return">error</span>)</code></p>

GetCapitalLeague fetches a Clan Capital league by ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>id</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../locations-rankings/#league">League</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getclan"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetClan(<span class="api-param">ctx: context.Context</span>, <span class="api-param">tag: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*Clan</span>, <span class="api-return">error</span>)</code></p>

GetClan fetches a clan profile by tag.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>tag</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../clans/#clan">Clan</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getclanlabels"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetClanLabels(<span class="api-param">ctx: context.Context</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]Label</span>, <span class="api-return">error</span>)</code></p>

GetClanLabels fetches clan labels with optional pagination.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#label">Label</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getclanwar"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetClanWar(<span class="api-param">ctx: context.Context</span>, <span class="api-param">clanTag: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*ClanWar</span>, <span class="api-return">error</span>)</code></p>

GetClanWar fetches the regular current-war endpoint for a clan.

This method does not fall back to CWL. Use GetCurrentWar when you want the
active normal war or the relevant Clan War League war.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>clanTag</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../wars/#clanwar">ClanWar</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getclanwars"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetClanWars(<span class="api-param">ctx: context.Context</span>, <span class="api-param">tags: []string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]ClanWar</span>, <span class="api-return">error</span>)</code></p>

GetClanWars fetches the regular current war for each clan tag in order.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>tags</strong> (<code>[]string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../wars/#clanwar">ClanWar</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getcurrentgoldpassseason"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetCurrentGoldPassSeason(<span class="api-param">ctx: context.Context</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*GoldPassSeason</span>, <span class="api-return">error</span>)</code></p>

GetCurrentGoldPassSeason fetches the current Gold Pass season.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../locations-rankings/#goldpassseason">GoldPassSeason</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getcurrentwar"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetCurrentWar(<span class="api-param">ctx: context.Context</span>, <span class="api-param">clanTag: string</span>, <span class="api-param">round: ...WarRound</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*ClanWar</span>, <span class="api-return">error</span>)</code></p>

GetCurrentWar returns the clan's active normal war or relevant CWL war.

The method first checks the regular current-war endpoint. If the clan is not
in a regular war, or the war log is private, it loads the CWL group and
returns the selected league round for the clan. Passing no round selects
CurrentWar. When no current war exists, the method returns nil, nil.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>clanTag</strong> (<code>string</code>)</p>
<p><strong>round</strong> (<code>...WarRound</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../wars/#clanwar">ClanWar</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getcurrentwars"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetCurrentWars(<span class="api-param">ctx: context.Context</span>, <span class="api-param">tags: []string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]ClanWar</span>, <span class="api-return">error</span>)</code></p>

GetCurrentWars fetches GetCurrentWar for each clan tag and omits clans with no
current war.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>tags</strong> (<code>[]string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../wars/#clanwar">ClanWar</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getequipment"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetEquipment(<span class="api-param">name: string</span>, <span class="api-param">level: int</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Equipment</span></code></p>

GetEquipment looks up hero equipment by name and level in embedded static
data.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
<p><strong>level</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../game-objects/#equipment">Equipment</a></code> </dd>
</dl>

</div>

<a id="client-getextendedcwlgroupdata"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetExtendedCWLGroupData(<span class="api-param">name: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*ExtendedCWLGroup</span></code></p>

GetExtendedCWLGroupData returns static medal data for a Clan War League tier
by name.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../wars/#extendedcwlgroup">ExtendedCWLGroup</a></code> </dd>
</dl>

</div>

<a id="client-gethero"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetHero(<span class="api-param">name: string</span>, <span class="api-param">level: int</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Hero</span></code></p>

GetHero looks up a hero by name and level in embedded static data.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
<p><strong>level</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../game-objects/#hero">Hero</a></code> </dd>
</dl>

</div>

<a id="client-getleague"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLeague(<span class="api-param">ctx: context.Context</span>, <span class="api-param">id: int</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*League</span>, <span class="api-return">error</span>)</code></p>

GetLeague fetches a home-village league by ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>id</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../locations-rankings/#league">League</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getleaguegroup"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLeagueGroup(<span class="api-param">ctx: context.Context</span>, <span class="api-param">clanTag: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*ClanWarLeagueGroup</span>, <span class="api-return">error</span>)</code></p>

GetLeagueGroup fetches the current Clan War League group for a clan.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>clanTag</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../wars/#clanwarleaguegroup">ClanWarLeagueGroup</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getleaguewar"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLeagueWar(<span class="api-param">ctx: context.Context</span>, <span class="api-param">clanTag: string</span>, <span class="api-param">round: WarRound</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*ClanWar</span>, <span class="api-return">error</span>)</code></p>

GetLeagueWar fetches the selected CWL round for a clan.

The returned war is oriented so Clan is the requested clan and Opponent is the
opposing side.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>clanTag</strong> (<code>string</code>)</p>
<p><strong>round</strong> (<code><a href="../wars/#warround">WarRound</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../wars/#clanwar">ClanWar</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getleaguewars"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLeagueWars(<span class="api-param">ctx: context.Context</span>, <span class="api-param">warTags: []string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]ClanWar</span>, <span class="api-return">error</span>)</code></p>

GetLeagueWars fetches CWL wars by war tag.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>warTags</strong> (<code>[]string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../wars/#clanwar">ClanWar</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocation"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocation(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: int</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*Location</span>, <span class="api-return">error</span>)</code></p>

GetLocation fetches a location by numeric ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../locations-rankings/#location">Location</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationclans"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationClans(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: int</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedClan</span>, <span class="api-return">error</span>)</code></p>

GetLocationClans fetches home-village clan rankings for a numeric location ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>int</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedclan">RankedClan</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationclansbuilderbase"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationClansBuilderBase(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: int</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedClan</span>, <span class="api-return">error</span>)</code></p>

GetLocationClansBuilderBase fetches Builder Base clan rankings for a numeric
location ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>int</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedclan">RankedClan</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationclansbuilderbasebylocationid"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationClansBuilderBaseByLocationID(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: string</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedClan</span>, <span class="api-return">error</span>)</code></p>

GetLocationClansBuilderBaseByLocationID fetches Builder Base clan rankings
for a location ID string.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>string</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedclan">RankedClan</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationclansbylocationid"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationClansByLocationID(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: string</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedClan</span>, <span class="api-return">error</span>)</code></p>

GetLocationClansByLocationID fetches home-village clan rankings for a
location ID string.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>string</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedclan">RankedClan</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationclanscapital"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationClansCapital(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: int</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedClan</span>, <span class="api-return">error</span>)</code></p>

GetLocationClansCapital fetches Clan Capital clan rankings for a numeric
location ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>int</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedclan">RankedClan</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationclanscapitalbylocationid"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationClansCapitalByLocationID(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: string</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedClan</span>, <span class="api-return">error</span>)</code></p>

GetLocationClansCapitalByLocationID fetches Clan Capital clan rankings for a
location ID string.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>string</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedclan">RankedClan</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationnamed"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationNamed(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationName: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*Location</span>, <span class="api-return">error</span>)</code></p>

GetLocationNamed returns the first location whose name matches
locationName case-insensitively.

It returns nil, nil when no matching location is found.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationName</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../locations-rankings/#location">Location</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationplayers"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationPlayers(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: int</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedPlayer</span>, <span class="api-return">error</span>)</code></p>

GetLocationPlayers fetches home-village player rankings for a numeric
location ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>int</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedplayer">RankedPlayer</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationplayersbuilderbase"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationPlayersBuilderBase(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: int</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedPlayer</span>, <span class="api-return">error</span>)</code></p>

GetLocationPlayersBuilderBase fetches Builder Base player rankings for a
numeric location ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>int</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedplayer">RankedPlayer</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationplayersbuilderbasebylocationid"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationPlayersBuilderBaseByLocationID(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: string</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedPlayer</span>, <span class="api-return">error</span>)</code></p>

GetLocationPlayersBuilderBaseByLocationID fetches Builder Base player
rankings for a location ID string.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>string</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedplayer">RankedPlayer</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getlocationplayersbylocationid"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetLocationPlayersByLocationID(<span class="api-param">ctx: context.Context</span>, <span class="api-param">locationID: string</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedPlayer</span>, <span class="api-return">error</span>)</code></p>

GetLocationPlayersByLocationID fetches home-village player rankings for a
location ID string.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>locationID</strong> (<code>string</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedplayer">RankedPlayer</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getmembers"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetMembers(<span class="api-param">ctx: context.Context</span>, <span class="api-param">clanTag: string</span>, <span class="api-param">limit: int</span>, <span class="api-param">after: string</span>, <span class="api-param">before: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]ClanMember</span>, <span class="api-return">error</span>)</code></p>

GetMembers fetches a clan member page by clan tag.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>clanTag</strong> (<code>string</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../clans/#clanmember">ClanMember</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getpet"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetPet(<span class="api-param">name: string</span>, <span class="api-param">level: int</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Pet</span></code></p>

GetPet looks up a pet by name and level in embedded static data.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
<p><strong>level</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../game-objects/#pet">Pet</a></code> </dd>
</dl>

</div>

<a id="client-getplayer"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetPlayer(<span class="api-param">ctx: context.Context</span>, <span class="api-param">tag: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*Player</span>, <span class="api-return">error</span>)</code></p>

GetPlayer fetches a player profile by tag.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>tag</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../players/#player">Player</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getplayerlabels"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetPlayerLabels(<span class="api-param">ctx: context.Context</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]Label</span>, <span class="api-return">error</span>)</code></p>

GetPlayerLabels fetches player labels with optional pagination.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#label">Label</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getplayerleaguegroup"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetPlayerLeagueGroup(<span class="api-param">ctx: context.Context</span>, <span class="api-param">playerTag: string</span>, <span class="api-param">leagueGroupTag: string</span>, <span class="api-param">leagueSeasonID: int</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*LeagueTierGroup</span>, <span class="api-return">error</span>)</code></p>

GetPlayerLeagueGroup fetches a legend league group and scopes it to a player.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>playerTag</strong> (<code>string</code>)</p>
<p><strong>leagueGroupTag</strong> (<code>string</code>)</p>
<p><strong>leagueSeasonID</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../battle-logs/#leaguetiergroup">LeagueTierGroup</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getplayerleaguehistory"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetPlayerLeagueHistory(<span class="api-param">ctx: context.Context</span>, <span class="api-param">playerTag: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]LeagueHistoryEntry</span>, <span class="api-return">error</span>)</code></p>

GetPlayerLeagueHistory fetches a player's legend league history.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>playerTag</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../battle-logs/#leaguehistoryentry">LeagueHistoryEntry</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getraidlog"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetRaidLog(<span class="api-param">ctx: context.Context</span>, <span class="api-param">clanTag: string</span>, <span class="api-param">limit: int</span>, <span class="api-param">after: string</span>, <span class="api-param">before: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RaidLogEntry</span>, <span class="api-return">error</span>)</code></p>

GetRaidLog fetches Clan Capital raid weekend log entries for a clan.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>clanTag</strong> (<code>string</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../raids/#raidlogentry">RaidLogEntry</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getseasonrankings"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetSeasonRankings(<span class="api-param">ctx: context.Context</span>, <span class="api-param">leagueID: int</span>, <span class="api-param">seasonID: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]RankedPlayer</span>, <span class="api-return">error</span>)</code></p>

GetSeasonRankings fetches player rankings for a league season.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>leagueID</strong> (<code>int</code>)</p>
<p><strong>seasonID</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#rankedplayer">RankedPlayer</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getseasons"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetSeasons(<span class="api-param">ctx: context.Context</span>, <span class="api-param">leagueID: int</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]string</span>, <span class="api-return">error</span>)</code></p>

GetSeasons fetches available season IDs for a league.

Passing leagueID 0 uses the default legend league ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>leagueID</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]string</code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getspell"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetSpell(<span class="api-param">name: string</span>, <span class="api-param">level: int</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Spell</span></code></p>

GetSpell looks up a spell by name and level in embedded static data.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
<p><strong>level</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../game-objects/#spell">Spell</a></code> </dd>
</dl>

</div>

<a id="client-gettranslation"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetTranslation(<span class="api-param">id: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Translation</span></code></p>

GetTranslation returns a translation entry by static-data translation ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>id</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../static-data/#translation">Translation</a></code> </dd>
</dl>

</div>

<a id="client-gettroop"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetTroop(<span class="api-param">name: string</span>, <span class="api-param">isHomeVillage: bool</span>, <span class="api-param">level: int</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Troop</span></code></p>

GetTroop looks up a troop by name, village, and level in embedded static data.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
<p><strong>isHomeVillage</strong> (<code>bool</code>)</p>
<p><strong>level</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../game-objects/#troop">Troop</a></code> </dd>
</dl>

</div>

<a id="client-getwarleague"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetWarLeague(<span class="api-param">ctx: context.Context</span>, <span class="api-param">id: int</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*League</span>, <span class="api-return">error</span>)</code></p>

GetWarLeague fetches a Clan War League tier by ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>id</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../locations-rankings/#league">League</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-getwarlog"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.GetWarLog(<span class="api-param">ctx: context.Context</span>, <span class="api-param">clanTag: string</span>, <span class="api-param">limit: int</span>, <span class="api-param">after: string</span>, <span class="api-param">before: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]ClanWarLogEntry</span>, <span class="api-return">error</span>)</code></p>

GetWarLog fetches public war log entries for a clan.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>clanTag</strong> (<code>string</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../wars/#clanwarlogentry">ClanWarLogEntry</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-login"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.Login(<span class="api-param">ctx: context.Context</span>, <span class="api-param">email: string</span>, <span class="api-param">password: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">error</span></code></p>

Login authenticates with developer-site email and password credentials.

The developer login flow discovers or creates API keys, stores them in the
underlying HTTP client, and uses those keys for later Clash API requests.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>email</strong> (<code>string</code>)</p>
<p><strong>password</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>error</code> </dd>
</dl>

</div>

<a id="client-loginwithtokens"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.LoginWithTokens(<span class="api-param">_: context.Context</span>, <span class="api-param">tokens: ...string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">error</span></code></p>

LoginWithTokens configures one or more existing Clash API tokens.

Tokens are rotated by the underlying HTTP client. The context parameter is
accepted for API symmetry with Login.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>_</strong> (<code>context.Context</code>)</p>
<p><strong>tokens</strong> (<code>...string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>error</code> </dd>
</dl>

</div>

<a id="client-parseaccountdata"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.ParseAccountData(<span class="api-param">data: map[string]any</span>)<span class="api-return-arrow"> -> </span><span class="api-return">AccountData</span></code></p>

ParseAccountData wraps arbitrary account-link data without mutating it.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>data</strong> (<code>map[string]any</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code><a href="../game-objects/#accountdata">AccountData</a></code> </dd>
</dl>

</div>

<a id="client-parsearmylink"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.ParseArmyLink(<span class="api-param">link: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">ArmyRecipe</span></code></p>

ParseArmyLink parses a full Clash army link or raw army payload using the
client's static data.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>link</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code><a href="../game-objects/#armyrecipe">ArmyRecipe</a></code> </dd>
</dl>

</div>

<a id="client-searchbuilderbaseleagues"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.SearchBuilderBaseLeagues(<span class="api-param">ctx: context.Context</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]League</span>, <span class="api-return">error</span>)</code></p>

SearchBuilderBaseLeagues fetches Builder Base leagues with optional
pagination.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#league">League</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-searchcapitalleagues"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.SearchCapitalLeagues(<span class="api-param">ctx: context.Context</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]League</span>, <span class="api-return">error</span>)</code></p>

SearchCapitalLeagues fetches Clan Capital leagues with optional pagination.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#league">League</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-searchclans"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.SearchClans(<span class="api-param">ctx: context.Context</span>, <span class="api-param">req: SearchClansRequest</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]Clan</span>, <span class="api-return">error</span>)</code></p>

SearchClans searches clans using the provided optional filters.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>req</strong> (<code><a href="#searchclansrequest">SearchClansRequest</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../clans/#clan">Clan</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-searchleagues"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.SearchLeagues(<span class="api-param">ctx: context.Context</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]League</span>, <span class="api-return">error</span>)</code></p>

SearchLeagues fetches home-village leagues with optional pagination.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#league">League</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-searchlocations"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.SearchLocations(<span class="api-param">ctx: context.Context</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]Location</span>, <span class="api-return">error</span>)</code></p>

SearchLocations fetches API locations with optional pagination.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#location">Location</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-searchwarleagues"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.SearchWarLeagues(<span class="api-param">ctx: context.Context</span>, <span class="api-param">limit: int</span>, <span class="api-param">before: string</span>, <span class="api-param">after: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]League</span>, <span class="api-return">error</span>)</code></p>

SearchWarLeagues fetches Clan War League tiers with optional pagination.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>limit</strong> (<code>int</code>)</p>
<p><strong>before</strong> (<code>string</code>)</p>
<p><strong>after</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../locations-rankings/#league">League</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="client-staticdata"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.StaticData()<span class="api-return-arrow"> -> </span><span class="api-return">*StaticData</span></code></p>

StaticData returns the client's embedded static-data index.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../static-data/#staticdata">StaticData</a></code> </dd>
</dl>

</div>

<a id="client-updatestatic"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.UpdateStatic(<span class="api-param">ctx: context.Context</span>)<span class="api-return-arrow"> -> </span><span class="api-return">error</span></code></p>

UpdateStatic downloads the latest ClashKing static-data and translation JSON,
writes the embedded source files, and refreshes this client's in-memory
StaticData.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>error</code> </dd>
</dl>

</div>

<a id="client-verifyplayertoken"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Client.VerifyPlayerToken(<span class="api-param">ctx: context.Context</span>, <span class="api-param">playerTag: string</span>, <span class="api-param">token: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">bool</span>, <span class="api-return">error</span>)</code></p>

VerifyPlayerToken verifies an in-game player API token.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>playerTag</strong> (<code>string</code>)</p>
<p><strong>token</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>bool</code> <code>error</code> </dd>
</dl>

</div>

## HTTPClient Methods

<a id="httpclient-do"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.HTTPClient.Do(<span class="api-param">ctx: context.Context</span>, <span class="api-param">method: string</span>, <span class="api-param">fullURL: string</span>, <span class="api-param">body: any</span>, <span class="api-param">options: RequestOptions</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">[]byte</span>, <span class="api-return">int</span>, <span class="api-return">int</span>, <span class="api-return">error</span>)</code></p>

Do sends one HTTP request and returns the response body, status code, retry
cache duration in seconds, and error.

Non-2xx API responses are converted into the package's typed HTTP errors.
Successful GET responses can be read from or written to the in-memory cache
depending on RequestOptions.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>method</strong> (<code>string</code>)</p>
<p><strong>fullURL</strong> (<code>string</code>)</p>
<p><strong>body</strong> (<code>any</code>)</p>
<p><strong>options</strong> (<code><a href="#requestoptions">RequestOptions</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]byte</code> <code>int</code> <code>int</code> <code>error</code> </dd>
</dl>

</div>

<a id="httpclient-logindeveloper"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.HTTPClient.LoginDeveloper(<span class="api-param">ctx: context.Context</span>, <span class="api-param">email: string</span>, <span class="api-param">password: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">error</span></code></p>

LoginDeveloper authenticates against the Clash developer site and configures
API tokens for subsequent Clash API requests.

The method reuses matching keys for the configured IP and key name when
possible, creating more keys until ClientConfig.KeyCount is satisfied.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
<p><strong>email</strong> (<code>string</code>)</p>
<p><strong>password</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>error</code> </dd>
</dl>

</div>

<a id="httpclient-settokens"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.HTTPClient.SetTokens(<span class="api-param">tokens: ...string</span>)</code></p>

SetTokens replaces the API tokens used for Authorization headers.

Tokens are rotated one per request. Passing no tokens clears authentication.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>tokens</strong> (<code>...string</code>)</p>
</dd>
</dl>

</div>

## Functions

<a id="newclient"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.NewClient(<span class="api-param">cfg: ClientConfig</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">*Client</span>, <span class="api-return">error</span>)</code></p>

NewClient constructs a Client from cfg and loads embedded static data.

If cfg.BaseURL is empty, DefaultClientConfig is used. BaseURL and
DeveloperBaseURL are normalized by removing trailing slashes.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>cfg</strong> (<code><a href="#clientconfig">ClientConfig</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#client">Client</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="defaultclientconfig"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.DefaultClientConfig()<span class="api-return-arrow"> -> </span><span class="api-return">ClientConfig</span></code></p>

DefaultClientConfig returns the recommended baseline configuration for the
official Clash of Clans API.

The defaults enable tag correction, GET response caching, embedded static
data, a 30 second timeout, and a conservative request throttle. Callers using
a ClashKing proxy typically override BaseURL and may enable Realtime.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code><a href="#clientconfig">ClientConfig</a></code> </dd>
</dl>

</div>

<a id="newhttpclient"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.NewHTTPClient(<span class="api-param">cfg: ClientConfig</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*HTTPClient</span></code></p>

NewHTTPClient constructs an HTTPClient from cfg.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>cfg</strong> (<code><a href="#clientconfig">ClientConfig</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#httpclient">HTTPClient</a></code> </dd>
</dl>

</div>


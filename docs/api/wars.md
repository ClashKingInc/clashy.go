# Wars

Classic war and Clan War League response models.

<a id="clanwar"></a>

## Clan War

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ClanWar</code></p>

ClanWar is the current, historical, or league war response.

For league wars found through GetCurrentWar or GetLeagueWar, the client
orients the result so Clan is the requested clan and Opponent is the opposing
side, even if the API returned the requested clan under opponent.

<div class="api-field" id="clanwar-state" markdown="1">

### `State`

<p><code><a href="#warstate">WarState</a></code> <span class="api-json">json: state</span></p>

State is the current state of the war.

</div>

<div class="api-field" id="clanwar-teamsize" markdown="1">

### `TeamSize`

<p><code>int</code> <span class="api-json">json: teamSize</span></p>

TeamSize is the roster size for each side.

</div>

<div class="api-field" id="clanwar-preparationstarttime" markdown="1">

### `PreparationStartTime`

<p><code>*<a href="../miscellaneous/#timestamp">Timestamp</a></code> <span class="api-json">json: preparationStartTime</span></p>

PreparationStartTime is when preparation day began.

</div>

<div class="api-field" id="clanwar-starttime" markdown="1">

### `StartTime`

<p><code>*<a href="../miscellaneous/#timestamp">Timestamp</a></code> <span class="api-json">json: startTime</span></p>

StartTime is when battle day starts.

</div>

<div class="api-field" id="clanwar-endtime" markdown="1">

### `EndTime`

<p><code>*<a href="../miscellaneous/#timestamp">Timestamp</a></code> <span class="api-json">json: endTime</span></p>

EndTime is when the war ends.

</div>

<div class="api-field" id="clanwar-clan" markdown="1">

### `Clan`

<p><code>*<a href="#warclan">WarClan</a></code> <span class="api-json">json: clan</span></p>

Clan is the requested clan side for oriented responses.

</div>

<div class="api-field" id="clanwar-opponent" markdown="1">

### `Opponent`

<p><code>*<a href="#warclan">WarClan</a></code> <span class="api-json">json: opponent</span></p>

Opponent is the opposing clan side for oriented responses.

</div>

<div class="api-field" id="clanwar-battlemodifier" markdown="1">

### `BattleModifier`

<p><code><a href="../clashy/#battlemodifier">BattleModifier</a></code> <span class="api-json">json: battleModifier</span></p>

BattleModifier describes event-specific modifiers when the API includes
one.

</div>

<div class="api-field" id="clanwar-wartag" markdown="1">

### `WarTag`

<p><code>string</code> <span class="api-json">json: tag</span></p>

WarTag is the CWL war tag. It is empty for normal classic wars.

</div>

<div class="api-field" id="clanwar-clantag" markdown="1">

### `ClanTag`

<p><code>string</code></p>

ClanTag is the requested clan tag associated with this response.

</div>

<div class="api-field" id="clanwar-leaguegroup" markdown="1">

### `LeagueGroup`

<p><code>*<a href="#clanwarleaguegroup">ClanWarLeagueGroup</a></code></p>

LeagueGroup is the CWL group used to find this war when available.

</div>

<a id="warclan"></a>

## War Clan

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.WarClan</code></p>

WarClan is one clan side of a classic war or CWL war.

<div class="api-field" id="warclan-tag" markdown="1">

### `Tag`

<p><code>string</code> <span class="api-json">json: tag</span></p>

Tag is the clan tag.

</div>

<div class="api-field" id="warclan-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the clan name.

</div>

<div class="api-field" id="warclan-badge" markdown="1">

### `Badge`

<p><code><a href="../locations-rankings/#badge">Badge</a></code> <span class="api-json">json: badgeUrls</span></p>

Badge contains clan badge image URLs.

</div>

<div class="api-field" id="warclan-level" markdown="1">

### `Level`

<p><code>int</code> <span class="api-json">json: clanLevel</span></p>

Level is the clan level.

</div>

<div class="api-field" id="warclan-attacks" markdown="1">

### `Attacks`

<p><code>int</code> <span class="api-json">json: attacks</span></p>

Attacks is the number of attacks used by this clan.

</div>

<div class="api-field" id="warclan-stars" markdown="1">

### `Stars`

<p><code>int</code> <span class="api-json">json: stars</span></p>

Stars is the total stars earned by this clan.

</div>

<div class="api-field" id="warclan-destruction" markdown="1">

### `Destruction`

<p><code>float64</code> <span class="api-json">json: destructionPercentage</span></p>

Destruction is the total destruction percentage earned by this clan.

</div>

<div class="api-field" id="warclan-expearned" markdown="1">

### `ExpEarned`

<p><code>int</code> <span class="api-json">json: expEarned</span></p>

ExpEarned is clan XP earned by this war when the endpoint includes it.

</div>

<div class="api-field" id="warclan-members" markdown="1">

### `Members`

<p><code>[]<a href="#clanwarmember">ClanWarMember</a></code> <span class="api-json">json: members</span></p>

Members is the war roster for this side.

</div>

<a id="clanwarmember"></a>

## Clan War Member

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ClanWarMember</code></p>

ClanWarMember is a player entry on one side of a war.

<div class="api-field" id="clanwarmember-tag" markdown="1">

### `Tag`

<p><code>string</code> <span class="api-json">json: tag</span></p>

Tag is the player's tag.

</div>

<div class="api-field" id="clanwarmember-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the player's display name at the time of the war.

</div>

<div class="api-field" id="clanwarmember-mapposition" markdown="1">

### `MapPosition`

<p><code>int</code> <span class="api-json">json: mapPosition</span></p>

MapPosition is the player's position on the war map.

</div>

<div class="api-field" id="clanwarmember-townhall" markdown="1">

### `Townhall`

<p><code>int</code> <span class="api-json">json: townhallLevel</span></p>

Townhall is the player's Town Hall level in the war response.

</div>

<div class="api-field" id="clanwarmember-opponentattacks" markdown="1">

### `OpponentAttacks`

<p><code>int</code> <span class="api-json">json: opponentAttacks</span></p>

OpponentAttacks is the number of attacks used against this base.

</div>

<div class="api-field" id="clanwarmember-attacks" markdown="1">

### `Attacks`

<p><code>[]<a href="#warattack">WarAttack</a></code> <span class="api-json">json: attacks</span></p>

Attacks contains attacks made by this member.

</div>

<div class="api-field" id="clanwarmember-bestopponentattack" markdown="1">

### `BestOpponentAttack`

<p><code>*<a href="#warattack">WarAttack</a></code> <span class="api-json">json: bestOpponentAttack</span></p>

BestOpponentAttack is the best attack received by this member.

</div>

<a id="warattack"></a>

## War Attack

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.WarAttack</code></p>

WarAttack is one attack inside a classic war or Clan War League war.

<div class="api-field" id="warattack-order" markdown="1">

### `Order`

<p><code>int</code> <span class="api-json">json: order</span></p>

Order is the attack order assigned by the API.

</div>

<div class="api-field" id="warattack-attackertag" markdown="1">

### `AttackerTag`

<p><code>string</code> <span class="api-json">json: attackerTag</span></p>

AttackerTag is the player tag of the attacker.

</div>

<div class="api-field" id="warattack-defendertag" markdown="1">

### `DefenderTag`

<p><code>string</code> <span class="api-json">json: defenderTag</span></p>

DefenderTag is the player tag of the defender.

</div>

<div class="api-field" id="warattack-stars" markdown="1">

### `Stars`

<p><code>int</code> <span class="api-json">json: stars</span></p>

Stars is the number of stars earned by the attack.

</div>

<div class="api-field" id="warattack-destruction" markdown="1">

### `Destruction`

<p><code>float64</code> <span class="api-json">json: destructionPercentage</span></p>

Destruction is the destruction percentage earned by the attack.

</div>

<div class="api-field" id="warattack-duration" markdown="1">

### `Duration`

<p><code>int</code> <span class="api-json">json: duration</span></p>

Duration is the attack duration in seconds.

</div>

<div class="api-field" id="warattack-attacker" markdown="1">

### `Attacker`

<p><code>*<a href="#clanwarmember">ClanWarMember</a></code></p>

Attacker is optionally linked to the attacker member when a caller enriches
the attack from the war member list.

</div>

<div class="api-field" id="warattack-defender" markdown="1">

### `Defender`

<p><code>*<a href="#clanwarmember">ClanWarMember</a></code></p>

Defender is optionally linked to the defender member when a caller enriches
the attack from the war member list.

</div>

<a id="clanwarlogentry"></a>

## Clan War Log Entry

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ClanWarLogEntry</code></p>

ClanWarLogEntry is one item from a clan's public war log.

<div class="api-field" id="clanwarlogentry-result" markdown="1">

### `Result`

<p><code><a href="#warresult">WarResult</a></code> <span class="api-json">json: result</span></p>

Result is the requested clan's result for this war.

</div>

<div class="api-field" id="clanwarlogentry-endtime" markdown="1">

### `EndTime`

<p><code>*<a href="../miscellaneous/#timestamp">Timestamp</a></code> <span class="api-json">json: endTime</span></p>

EndTime is when the war ended.

</div>

<div class="api-field" id="clanwarlogentry-teamsize" markdown="1">

### `TeamSize`

<p><code>int</code> <span class="api-json">json: teamSize</span></p>

TeamSize is the roster size for each side.

</div>

<div class="api-field" id="clanwarlogentry-clan" markdown="1">

### `Clan`

<p><code>*<a href="#warclan">WarClan</a></code> <span class="api-json">json: clan</span></p>

Clan is the requested clan side.

</div>

<div class="api-field" id="clanwarlogentry-opponent" markdown="1">

### `Opponent`

<p><code>*<a href="#warclan">WarClan</a></code> <span class="api-json">json: opponent</span></p>

Opponent is the opposing clan side.

</div>

<div class="api-field" id="clanwarlogentry-battlemodifier" markdown="1">

### `BattleModifier`

<p><code><a href="../clashy/#battlemodifier">BattleModifier</a></code> <span class="api-json">json: battleModifier</span></p>

BattleModifier describes event-specific modifiers when the API includes
one.

</div>

<a id="clanwarleaguegroup"></a>

## Clan War League Group

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ClanWarLeagueGroup</code></p>

ClanWarLeagueGroup is the current CWL group for a clan.

<div class="api-field" id="clanwarleaguegroup-state" markdown="1">

### `State`

<p><code>string</code> <span class="api-json">json: state</span></p>

State is the group state returned by the API.

</div>

<div class="api-field" id="clanwarleaguegroup-season" markdown="1">

### `Season`

<p><code>string</code> <span class="api-json">json: season</span></p>

Season is the CWL season identifier returned by the API.

</div>

<div class="api-field" id="clanwarleaguegroup-clans" markdown="1">

### `Clans`

<p><code>[]<a href="#clanwarleagueclan">ClanWarLeagueClan</a></code> <span class="api-json">json: clans</span></p>

Clans contains the clans participating in the group.

</div>

<div class="api-field" id="clanwarleaguegroup-rounds" markdown="1">

### `Rounds`

<p><code>[]struct {
	WarTags []string `json:"warTags,omitempty"`
}</code> <span class="api-json">json: rounds</span></p>

Rounds contains CWL war tags grouped by round. Future rounds may contain
placeholder "#0" tags.

</div>

<a id="clanwarleagueclan"></a>

## Clan War League Clan

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ClanWarLeagueClan</code></p>

ClanWarLeagueClan is a clan entry inside a CWL group.

<div class="api-field" id="clanwarleagueclan-tag" markdown="1">

### `Tag`

<p><code>string</code> <span class="api-json">json: tag</span></p>

Tag is the clan tag.

</div>

<div class="api-field" id="clanwarleagueclan-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the clan name.

</div>

<div class="api-field" id="clanwarleagueclan-badge" markdown="1">

### `Badge`

<p><code><a href="../locations-rankings/#badge">Badge</a></code> <span class="api-json">json: badgeUrls</span></p>

Badge contains clan badge image URLs.

</div>

<div class="api-field" id="clanwarleagueclan-level" markdown="1">

### `Level`

<p><code>int</code> <span class="api-json">json: clanLevel</span></p>

Level is the clan level.

</div>

<a id="extendedcwlgroup"></a>

## Extended CWLGroup

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ExtendedCWLGroup</code></p>

ExtendedCWLGroup contains static medal information for a CWL league.

<div class="api-field" id="extendedcwlgroup-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the league display name.

</div>

<div class="api-field" id="extendedcwlgroup-firstplacemedals" markdown="1">

### `FirstPlaceMedals`

<p><code>int</code> <span class="api-json">json: first_place_medals</span></p>

FirstPlaceMedals is the medal reward for first place.

</div>

<div class="api-field" id="extendedcwlgroup-secondplacemedals" markdown="1">

### `SecondPlaceMedals`

<p><code>int</code> <span class="api-json">json: second_place_medals</span></p>

SecondPlaceMedals is the medal reward for second place.

</div>

<a id="warround"></a>

## War Round

<p class="api-signature"><span class="api-kind">type</span> <code>clashy.WarRound</code></p>

WarRound identifies the logical CWL round requested from GetCurrentWar or
GetLeagueWar.

<p><code>int</code></p>

### Values

<div class="api-field" id="previouswar" markdown="1">

#### `PreviousWar`

<p><code>iota</code></p>

PreviousWar selects the previous completed or in-war CWL round.

</div>

<div class="api-field" id="currentwar" markdown="1">

#### `CurrentWar`

CurrentWar selects the active CWL war, or the latest completed/in-war
round when the latest real round is only preparation.

</div>

<div class="api-field" id="currentpreparation" markdown="1">

#### `CurrentPreparation`

CurrentPreparation selects the upcoming CWL preparation round when one is
available.

</div>

<a id="warstate"></a>

## War State

<p class="api-signature"><span class="api-kind">type</span> <code>clashy.WarState</code></p>

WarState is the lifecycle state of a classic war or CWL war.

<p><code>string</code></p>

### Values

<div class="api-field" id="warstatenotinwar" markdown="1">

#### `WarStateNotInWar`

<p><code>"notInWar"</code></p>

WarStateNotInWar means the clan is not in a regular war.

</div>

<div class="api-field" id="warstatepreparation" markdown="1">

#### `WarStatePreparation`

<p><code>"preparation"</code></p>

WarStatePreparation means the war is in preparation day.

</div>

<div class="api-field" id="warstateinwar" markdown="1">

#### `WarStateInWar`

<p><code>"inWar"</code></p>

WarStateInWar means battle day is active.

</div>

<div class="api-field" id="warstateended" markdown="1">

#### `WarStateEnded`

<p><code>"warEnded"</code></p>

WarStateEnded means the war has ended.

</div>

<a id="warresult"></a>

## War Result

<p class="api-signature"><span class="api-kind">type</span> <code>clashy.WarResult</code></p>

WarResult is the requested clan's result in a war log entry.

<p><code>string</code></p>

### Values

<div class="api-field" id="warresultwin" markdown="1">

#### `WarResultWin`

<p><code>"win"</code></p>

WarResultWin means the requested clan won.

</div>

<div class="api-field" id="warresultlose" markdown="1">

#### `WarResultLose`

<p><code>"lose"</code></p>

WarResultLose means the requested clan lost.

</div>

<div class="api-field" id="warresulttie" markdown="1">

#### `WarResultTie`

<p><code>"tie"</code></p>

WarResultTie means the war ended in a tie.

</div>

## Clan War Methods

<a id="clanwar-attacks"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.ClanWar.Attacks()<span class="api-return-arrow"> -> </span><span class="api-return">[]WarAttack</span></code></p>

Attacks returns all attacks made by both sides of the war.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="#warattack">WarAttack</a></code> </dd>
</dl>

</div>

<a id="clanwar-type"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.ClanWar.Type()<span class="api-return-arrow"> -> </span><span class="api-return">string</span></code></p>

Type returns "cwl" when the war has a CWL war tag and "random" otherwise.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>string</code> </dd>
</dl>

</div>


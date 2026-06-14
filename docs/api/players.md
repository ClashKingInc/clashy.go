# Players

Player profile models and helpers for achievements, units, spells, and labels.

<a id="player"></a>

## Player

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Player</code></p>

Player is the full player profile returned by GetPlayer.

<div class="api-field" id="player-tag" markdown="1">

### `Tag`

<p><code>string</code> <span class="api-json">json: tag</span></p>

Tag is the player's Clash tag.

</div>

<div class="api-field" id="player-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the player's current display name.

</div>

<div class="api-field" id="player-explevel" markdown="1">

### `ExpLevel`

<p><code>int</code> <span class="api-json">json: expLevel</span></p>

ExpLevel is the player's experience level.

</div>

<div class="api-field" id="player-trophies" markdown="1">

### `Trophies`

<p><code>int</code> <span class="api-json">json: trophies</span></p>

Trophies is the player's current home village trophy count.

</div>

<div class="api-field" id="player-besttrophies" markdown="1">

### `BestTrophies`

<p><code>int</code> <span class="api-json">json: bestTrophies</span></p>

BestTrophies is the player's all-time best home village trophy count.

</div>

<div class="api-field" id="player-warstars" markdown="1">

### `WarStars`

<p><code>int</code> <span class="api-json">json: warStars</span></p>

WarStars is the player's lifetime war star count.

</div>

<div class="api-field" id="player-townhall" markdown="1">

### `TownHall`

<p><code>int</code> <span class="api-json">json: townHallLevel</span></p>

TownHall is the player's home village Town Hall level.

</div>

<div class="api-field" id="player-townhallweapon" markdown="1">

### `TownHallWeapon`

<p><code>int</code> <span class="api-json">json: townHallWeaponLevel</span></p>

TownHallWeapon is the weapon level for Town Hall levels that have one.

</div>

<div class="api-field" id="player-builderhall" markdown="1">

### `BuilderHall`

<p><code>int</code> <span class="api-json">json: builderHallLevel</span></p>

BuilderHall is the player's Builder Hall level.

</div>

<div class="api-field" id="player-bestbuilderbasetrophies" markdown="1">

### `BestBuilderBaseTrophies`

<p><code>int</code> <span class="api-json">json: bestBuilderBaseTrophies</span></p>

BestBuilderBaseTrophies is the all-time best Builder Base trophy count.

</div>

<div class="api-field" id="player-versusattackwins" markdown="1">

### `VersusAttackWins`

<p><code>int</code> <span class="api-json">json: versusBattleWins</span></p>

VersusAttackWins is the legacy Builder Base attack-win field.

</div>

<div class="api-field" id="player-donations" markdown="1">

### `Donations`

<p><code>int</code> <span class="api-json">json: donations</span></p>

Donations is the number of troops donated this season.

</div>

<div class="api-field" id="player-received" markdown="1">

### `Received`

<p><code>int</code> <span class="api-json">json: donationsReceived</span></p>

Received is the number of donated troops received this season.

</div>

<div class="api-field" id="player-clancapitalcontributions" markdown="1">

### `ClanCapitalContributions`

<p><code>int</code> <span class="api-json">json: clanCapitalContributions</span></p>

ClanCapitalContributions is the lifetime Clan Capital contribution count.

</div>

<div class="api-field" id="player-clanrank" markdown="1">

### `ClanRank`

<p><code>int</code> <span class="api-json">json: clanRank</span></p>

ClanRank is the player's current rank inside their clan.

</div>

<div class="api-field" id="player-clanpreviousrank" markdown="1">

### `ClanPreviousRank`

<p><code>int</code> <span class="api-json">json: previousClanRank</span></p>

ClanPreviousRank is the player's previous rank inside their clan.

</div>

<div class="api-field" id="player-versustrophies" markdown="1">

### `VersusTrophies`

<p><code>int</code> <span class="api-json">json: versusTrophies</span></p>

VersusTrophies is the legacy Builder Base trophy field.

</div>

<div class="api-field" id="player-builderbasetrophies" markdown="1">

### `BuilderBaseTrophies`

<p><code>int</code> <span class="api-json">json: builderBaseTrophies</span></p>

BuilderBaseTrophies is the player's current Builder Base trophy count.

</div>

<div class="api-field" id="player-leaguetier" markdown="1">

### `LeagueTier`

<p><code><a href="../locations-rankings/#league">League</a></code> <span class="api-json">json: leagueTier</span></p>

LeagueTier is the player's home village league.

</div>

<div class="api-field" id="player-builderbaseleague" markdown="1">

### `BuilderBaseLeague`

<p><code>*<a href="../locations-rankings/#league">League</a></code> <span class="api-json">json: builderBaseLeague</span></p>

BuilderBaseLeague is the player's Builder Base league.

</div>

<div class="api-field" id="player-role" markdown="1">

### `Role`

<p><code><a href="../enums/#role">Role</a></code> <span class="api-json">json: role</span></p>

Role is the player's role in their current clan.

</div>

<div class="api-field" id="player-clan" markdown="1">

### `Clan`

<p><code>*<a href="../clans/#playerclan">PlayerClan</a></code> <span class="api-json">json: clan</span></p>

Clan is the compact clan object for the player's current clan.

</div>

<div class="api-field" id="player-currentleaguegrouptag" markdown="1">

### `CurrentLeagueGroupTag`

<p><code>string</code> <span class="api-json">json: currentLeagueGroupTag</span></p>

CurrentLeagueGroupTag is the active legend league group tag when present.

</div>

<div class="api-field" id="player-currentleagueseasonid" markdown="1">

### `CurrentLeagueSeasonID`

<p><code>int</code> <span class="api-json">json: currentLeagueSeasonId</span></p>

CurrentLeagueSeasonID is the active legend league season ID when present.

</div>

<div class="api-field" id="player-previousleaguegrouptag" markdown="1">

### `PreviousLeagueGroupTag`

<p><code>string</code> <span class="api-json">json: previousLeagueGroupTag</span></p>

PreviousLeagueGroupTag is the previous legend league group tag when
present.

</div>

<div class="api-field" id="player-previousleagueseasonid" markdown="1">

### `PreviousLeagueSeasonID`

<p><code>int</code> <span class="api-json">json: previousLeagueSeasonId</span></p>

PreviousLeagueSeasonID is the previous legend league season ID when
present.

</div>

<div class="api-field" id="player-legendstatistics" markdown="1">

### `LegendStatistics`

<p><code>*<a href="#legendstatistics">LegendStatistics</a></code> <span class="api-json">json: legendStatistics</span></p>

LegendStatistics contains legend trophies and seasonal legend finishes.

</div>

<div class="api-field" id="player-labels" markdown="1">

### `Labels`

<p><code>[]<a href="../locations-rankings/#label">Label</a></code> <span class="api-json">json: labels</span></p>

Labels are public player labels.

</div>

<div class="api-field" id="player-achievements" markdown="1">

### `Achievements`

<p><code>[]<a href="#achievement">Achievement</a></code> <span class="api-json">json: achievements</span></p>

Achievements contains achievement progress for both villages and Clan
Capital.

</div>

<div class="api-field" id="player-troops" markdown="1">

### `Troops`

<p><code>[]<a href="../game-objects/#troop">Troop</a></code> <span class="api-json">json: troops</span></p>

Troops contains unlocked troops with current and max levels.

</div>

<div class="api-field" id="player-heroes" markdown="1">

### `Heroes`

<p><code>[]<a href="../game-objects/#hero">Hero</a></code> <span class="api-json">json: heroes</span></p>

Heroes contains unlocked heroes with current and max levels.

</div>

<div class="api-field" id="player-spells" markdown="1">

### `Spells`

<p><code>[]<a href="../game-objects/#spell">Spell</a></code> <span class="api-json">json: spells</span></p>

Spells contains unlocked spells with current and max levels.

</div>

<div class="api-field" id="player-heroequipment" markdown="1">

### `HeroEquipment`

<p><code>[]<a href="../game-objects/#equipment">Equipment</a></code> <span class="api-json">json: heroEquipment</span></p>

HeroEquipment contains unlocked hero equipment with current and max levels.

</div>

<a id="legendstatistics"></a>

## Legend Statistics

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.LegendStatistics</code></p>

LegendStatistics contains a player's legend trophies and season snapshots.

<div class="api-field" id="legendstatistics-legendtrophies" markdown="1">

### `LegendTrophies`

<p><code>int</code> <span class="api-json">json: legendTrophies</span></p>

LegendTrophies is the player's lifetime legend trophy count.

</div>

<div class="api-field" id="legendstatistics-bestseason" markdown="1">

### `BestSeason`

<p><code>*<a href="../locations-rankings/#season">Season</a></code> <span class="api-json">json: bestSeason</span></p>

BestSeason is the player's best legend season.

</div>

<div class="api-field" id="legendstatistics-previousseason" markdown="1">

### `PreviousSeason`

<p><code>*<a href="../locations-rankings/#season">Season</a></code> <span class="api-json">json: previousSeason</span></p>

PreviousSeason is the player's previous legend season.

</div>

<div class="api-field" id="legendstatistics-bestversusseason" markdown="1">

### `BestVersusSeason`

<p><code>*<a href="../locations-rankings/#season">Season</a></code> <span class="api-json">json: bestVersusSeason</span></p>

BestVersusSeason is the player's legacy best Builder Base season.

</div>

<div class="api-field" id="legendstatistics-currentseason" markdown="1">

### `CurrentSeason`

<p><code>*<a href="../locations-rankings/#season">Season</a></code> <span class="api-json">json: currentSeason</span></p>

CurrentSeason is the player's current legend season progress.

</div>

<a id="achievement"></a>

## Achievement

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Achievement</code></p>

Achievement describes one player achievement and its current progress.

<div class="api-field" id="achievement-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the achievement display name.

</div>

<div class="api-field" id="achievement-stars" markdown="1">

### `Stars`

<p><code>int</code> <span class="api-json">json: stars</span></p>

Stars is the number of achievement stars earned.

</div>

<div class="api-field" id="achievement-value" markdown="1">

### `Value`

<p><code>int</code> <span class="api-json">json: value</span></p>

Value is the current progress value.

</div>

<div class="api-field" id="achievement-target" markdown="1">

### `Target`

<p><code>int</code> <span class="api-json">json: target</span></p>

Target is the value needed to complete the achievement.

</div>

<div class="api-field" id="achievement-info" markdown="1">

### `Info`

<p><code>string</code> <span class="api-json">json: info</span></p>

Info describes the achievement goal.

</div>

<div class="api-field" id="achievement-completioninfo" markdown="1">

### `CompletionInfo`

<p><code>string</code> <span class="api-json">json: completionInfo</span></p>

CompletionInfo describes the completed achievement state.

</div>

<div class="api-field" id="achievement-village" markdown="1">

### `Village`

<p><code>string</code> <span class="api-json">json: village</span></p>

Village identifies the village or game area for the achievement.

</div>

<a id="playerhouseelement"></a>

## Player House Element

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.PlayerHouseElement</code></p>

PlayerHouseElement is one cosmetic element of a player's house.

<div class="api-field" id="playerhouseelement-id" markdown="1">

### `ID`

<p><code>int</code> <span class="api-json">json: id</span></p>

ID is the cosmetic element identifier.

</div>

<div class="api-field" id="playerhouseelement-type" markdown="1">

### `Type`

<p><code>string</code> <span class="api-json">json: type</span></p>

Type is the cosmetic element type.

</div>

## Player Methods

<a id="player-buildertroops"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Player.BuilderTroops()<span class="api-return-arrow"> -> </span><span class="api-return">[]Troop</span></code></p>

BuilderTroops returns troops that belong to Builder Base.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../game-objects/#troop">Troop</a></code> </dd>
</dl>

</div>

<a id="player-getachievement"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Player.GetAchievement(<span class="api-param">name: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Achievement</span></code></p>

GetAchievement returns the achievement with the provided display name.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#achievement">Achievement</a></code> </dd>
</dl>

</div>

<a id="player-gethero"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Player.GetHero(<span class="api-param">name: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Hero</span></code></p>

GetHero returns the hero with the provided display name.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../game-objects/#hero">Hero</a></code> </dd>
</dl>

</div>

<a id="player-getspell"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Player.GetSpell(<span class="api-param">name: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Spell</span></code></p>

GetSpell returns the spell with the provided display name.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../game-objects/#spell">Spell</a></code> </dd>
</dl>

</div>

<a id="player-gettroop"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Player.GetTroop(<span class="api-param">name: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Troop</span></code></p>

GetTroop returns the troop with the provided display name.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="../game-objects/#troop">Troop</a></code> </dd>
</dl>

</div>

<a id="player-hometroops"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Player.HomeTroops()<span class="api-return-arrow"> -> </span><span class="api-return">[]Troop</span></code></p>

HomeTroops returns troops that belong to the home village.

Older API responses may omit Village for home-village troops, so an empty
village is treated as home.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]<a href="../game-objects/#troop">Troop</a></code> </dd>
</dl>

</div>


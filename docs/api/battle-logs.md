# Battle Logs

Player battle logs and legend league group records.

<a id="battlelogentry"></a>

## Battle Log Entry

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.BattleLogEntry</code></p>

BattleLogEntry is one player battle log entry.

<div class="api-field" id="battlelogentry-battletype" markdown="1">

### `BattleType`

<p><code>string</code> <span class="api-json">json: battleType</span></p>

BattleType describes the game mode for the battle.

</div>

<div class="api-field" id="battlelogentry-attack" markdown="1">

### `Attack`

<p><code>bool</code> <span class="api-json">json: attack</span></p>

Attack reports whether the entry is an attack made by the requested
player. False entries are defenses.

</div>

<div class="api-field" id="battlelogentry-armysharecode" markdown="1">

### `ArmyShareCode`

<p><code>string</code> <span class="api-json">json: armyShareCode</span></p>

ArmyShareCode is the in-game army share payload when available.

</div>

<div class="api-field" id="battlelogentry-opponentplayertag" markdown="1">

### `OpponentPlayerTag`

<p><code>string</code> <span class="api-json">json: opponentPlayerTag</span></p>

OpponentPlayerTag is the opponent's player tag.

</div>

<div class="api-field" id="battlelogentry-stars" markdown="1">

### `Stars`

<p><code>int</code> <span class="api-json">json: stars</span></p>

Stars is the number of stars earned by the attacker.

</div>

<div class="api-field" id="battlelogentry-destructionpercentage" markdown="1">

### `DestructionPercentage`

<p><code>int</code> <span class="api-json">json: destructionPercentage</span></p>

DestructionPercentage is the destruction percentage earned by the attacker.

</div>

<div class="api-field" id="battlelogentry-lootedresources" markdown="1">

### `LootedResources`

<p><code>[]<a href="#resource">Resource</a></code> <span class="api-json">json: lootedResources</span></p>

LootedResources contains resources actually looted.

</div>

<div class="api-field" id="battlelogentry-extralootedresources" markdown="1">

### `ExtraLootedResources`

<p><code>[]<a href="#resource">Resource</a></code> <span class="api-json">json: extraLootedResources</span></p>

ExtraLootedResources contains bonus resources awarded by the battle.

</div>

<div class="api-field" id="battlelogentry-availableloot" markdown="1">

### `AvailableLoot`

<p><code>[]<a href="#resource">Resource</a></code> <span class="api-json">json: availableLoot</span></p>

AvailableLoot contains resources that were available before the battle.

</div>

<a id="resource"></a>

## Resource

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Resource</code></p>

Resource is a named resource amount in a player battle log entry.

<div class="api-field" id="resource-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the resource name, such as gold, elixir, or dark elixir.

</div>

<div class="api-field" id="resource-amount" markdown="1">

### `Amount`

<p><code>int</code> <span class="api-json">json: amount</span></p>

Amount is the resource quantity.

</div>

<a id="leaguehistoryentry"></a>

## League History Entry

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.LeagueHistoryEntry</code></p>

LeagueHistoryEntry is one historical legend-league season result.

<div class="api-field" id="leaguehistoryentry-leagueseasonid" markdown="1">

### `LeagueSeasonID`

<p><code>int</code> <span class="api-json">json: leagueSeasonId</span></p>

LeagueSeasonID is the numeric legend season identifier.

</div>

<div class="api-field" id="leaguehistoryentry-leaguetrophies" markdown="1">

### `LeagueTrophies`

<p><code>int</code> <span class="api-json">json: leagueTrophies</span></p>

LeagueTrophies is the player's ending legend trophy count.

</div>

<div class="api-field" id="leaguehistoryentry-leaguetierid" markdown="1">

### `LeagueTierID`

<p><code>int</code> <span class="api-json">json: leagueTierId</span></p>

LeagueTierID is the league tier identifier for the season.

</div>

<div class="api-field" id="leaguehistoryentry-placement" markdown="1">

### `Placement`

<p><code>int</code> <span class="api-json">json: placement</span></p>

Placement is the player's final placement.

</div>

<div class="api-field" id="leaguehistoryentry-attackwins" markdown="1">

### `AttackWins`

<p><code>int</code> <span class="api-json">json: attackWins</span></p>

AttackWins is the number of attack wins.

</div>

<div class="api-field" id="leaguehistoryentry-attacklosses" markdown="1">

### `AttackLosses`

<p><code>int</code> <span class="api-json">json: attackLosses</span></p>

AttackLosses is the number of attack losses.

</div>

<div class="api-field" id="leaguehistoryentry-attackstars" markdown="1">

### `AttackStars`

<p><code>int</code> <span class="api-json">json: attackStars</span></p>

AttackStars is the total stars earned on attack.

</div>

<div class="api-field" id="leaguehistoryentry-defensewins" markdown="1">

### `DefenseWins`

<p><code>int</code> <span class="api-json">json: defenseWins</span></p>

DefenseWins is the number of defense wins.

</div>

<div class="api-field" id="leaguehistoryentry-defenselosses" markdown="1">

### `DefenseLosses`

<p><code>int</code> <span class="api-json">json: defenseLosses</span></p>

DefenseLosses is the number of defense losses.

</div>

<div class="api-field" id="leaguehistoryentry-defensestars" markdown="1">

### `DefenseStars`

<p><code>int</code> <span class="api-json">json: defenseStars</span></p>

DefenseStars is the total stars allowed on defense.

</div>

<div class="api-field" id="leaguehistoryentry-maxbattles" markdown="1">

### `MaxBattles`

<p><code>int</code> <span class="api-json">json: maxBattles</span></p>

MaxBattles is the maximum battle count for the season.

</div>

<a id="leaguetiergroup"></a>

## League Tier Group

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.LeagueTierGroup</code></p>

LeagueTierGroup contains members and battle logs for a legend league group.

<div class="api-field" id="leaguetiergroup-members" markdown="1">

### `Members`

<p><code>[]<a href="#leaguetiergroupmember">LeagueTierGroupMember</a></code> <span class="api-json">json: members</span></p>

Members contains the players in the legend group.

</div>

<div class="api-field" id="leaguetiergroup-attacklogs" markdown="1">

### `AttackLogs`

<p><code>[]<a href="#leaguetiergroupbattlelogentry">LeagueTierGroupBattleLogEntry</a></code> <span class="api-json">json: attackLogs</span></p>

AttackLogs contains attack entries for the requested player.

</div>

<div class="api-field" id="leaguetiergroup-defenselogs" markdown="1">

### `DefenseLogs`

<p><code>[]<a href="#leaguetiergroupbattlelogentry">LeagueTierGroupBattleLogEntry</a></code> <span class="api-json">json: defenseLogs</span></p>

DefenseLogs contains defense entries for the requested player.

</div>

<a id="leaguetiergroupmember"></a>

## League Tier Group Member

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.LeagueTierGroupMember</code></p>

LeagueTierGroupMember is one player in a legend league group.

<div class="api-field" id="leaguetiergroupmember-playertag" markdown="1">

### `PlayerTag`

<p><code>string</code> <span class="api-json">json: playerTag</span></p>

PlayerTag is the player's tag.

</div>

<div class="api-field" id="leaguetiergroupmember-playername" markdown="1">

### `PlayerName`

<p><code>string</code> <span class="api-json">json: playerName</span></p>

PlayerName is the player's display name.

</div>

<div class="api-field" id="leaguetiergroupmember-clantag" markdown="1">

### `ClanTag`

<p><code>string</code> <span class="api-json">json: clanTag</span></p>

ClanTag is the player's clan tag when present.

</div>

<div class="api-field" id="leaguetiergroupmember-clanname" markdown="1">

### `ClanName`

<p><code>string</code> <span class="api-json">json: clanName</span></p>

ClanName is the player's clan name when present.

</div>

<div class="api-field" id="leaguetiergroupmember-leaguetrophies" markdown="1">

### `LeagueTrophies`

<p><code>int</code> <span class="api-json">json: leagueTrophies</span></p>

LeagueTrophies is the player's current legend trophy count.

</div>

<div class="api-field" id="leaguetiergroupmember-attackwincount" markdown="1">

### `AttackWinCount`

<p><code>int</code> <span class="api-json">json: attackWinCount</span></p>

AttackWinCount is the player's attack win count in the group.

</div>

<div class="api-field" id="leaguetiergroupmember-attacklosecount" markdown="1">

### `AttackLoseCount`

<p><code>int</code> <span class="api-json">json: attackLoseCount</span></p>

AttackLoseCount is the player's attack loss count in the group.

</div>

<div class="api-field" id="leaguetiergroupmember-defensewincount" markdown="1">

### `DefenseWinCount`

<p><code>int</code> <span class="api-json">json: defenseWinCount</span></p>

DefenseWinCount is the player's defense win count in the group.

</div>

<div class="api-field" id="leaguetiergroupmember-defenselosecount" markdown="1">

### `DefenseLoseCount`

<p><code>int</code> <span class="api-json">json: defenseLoseCount</span></p>

DefenseLoseCount is the player's defense loss count in the group.

</div>

<a id="leaguetiergroupbattlelogentry"></a>

## League Tier Group Battle Log Entry

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.LeagueTierGroupBattleLogEntry</code></p>

LeagueTierGroupBattleLogEntry is one attack or defense inside a legend group.

<div class="api-field" id="leaguetiergroupbattlelogentry-opponentplayertag" markdown="1">

### `OpponentPlayerTag`

<p><code>string</code> <span class="api-json">json: opponentPlayerTag</span></p>

OpponentPlayerTag is the opponent's player tag.

</div>

<div class="api-field" id="leaguetiergroupbattlelogentry-opponentname" markdown="1">

### `OpponentName`

<p><code>string</code> <span class="api-json">json: opponentName</span></p>

OpponentName is the opponent's player name.

</div>

<div class="api-field" id="leaguetiergroupbattlelogentry-stars" markdown="1">

### `Stars`

<p><code>int</code> <span class="api-json">json: stars</span></p>

Stars is the number of stars earned by the attacker.

</div>

<div class="api-field" id="leaguetiergroupbattlelogentry-destructionpercentage" markdown="1">

### `DestructionPercentage`

<p><code>int</code> <span class="api-json">json: destructionPercentage</span></p>

DestructionPercentage is the destruction percentage earned by the attacker.

</div>

<div class="api-field" id="leaguetiergroupbattlelogentry-trophies" markdown="1">

### `Trophies`

<p><code>int</code> <span class="api-json">json: trophies</span></p>

Trophies is the trophy delta for the battle.

</div>

<div class="api-field" id="leaguetiergroupbattlelogentry-creationtime" markdown="1">

### `CreationTime`

<p><code>string</code> <span class="api-json">json: creationTime</span></p>

CreationTime is the API timestamp for the battle.

</div>


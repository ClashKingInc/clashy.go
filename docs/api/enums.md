# Enums

Enum-like string and integer types used across the client.

<a id="role"></a>

## Role

<p class="api-signature"><span class="api-kind">type</span> <code>clashy.Role</code></p>

Role is a member's role inside a clan.

<p><code>string</code></p>

### Values

<div class="api-field" id="rolemember" markdown="1">

#### `RoleMember`

<p><code>"member"</code></p>

RoleMember is a regular clan member.

</div>

<div class="api-field" id="roleelder" markdown="1">

#### `RoleElder`

<p><code>"admin"</code></p>

RoleElder is a clan elder. The Clash API value is "admin".

</div>

<div class="api-field" id="rolecoleader" markdown="1">

#### `RoleCoLeader`

<p><code>"coLeader"</code></p>

RoleCoLeader is a clan co-leader.

</div>

<div class="api-field" id="roleleader" markdown="1">

#### `RoleLeader`

<p><code>"leader"</code></p>

RoleLeader is the clan leader.

</div>

<a id="villagetype"></a>

## Village Type

<p class="api-signature"><span class="api-kind">type</span> <code>clashy.VillageType</code></p>

VillageType identifies the village or game area for static data and units.

<p><code>string</code></p>

### Values

<div class="api-field" id="villagehome" markdown="1">

#### `VillageHome`

<p><code>"home"</code></p>

VillageHome is the home village.

</div>

<div class="api-field" id="villagebuilderbase" markdown="1">

#### `VillageBuilderBase`

<p><code>"builderBase"</code></p>

VillageBuilderBase is Builder Base.

</div>

<div class="api-field" id="villageclancapital" markdown="1">

#### `VillageClanCapital`

<p><code>"clanCapital"</code></p>

VillageClanCapital is Clan Capital.

</div>

<a id="loadgamedata"></a>

## Load Game Data

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.LoadGameData</code></p>

LoadGameData describes when static game data should be loaded.

<div class="api-field" id="loadgamedata-default" markdown="1">

### `Default`

<p><code>bool</code></p>

Default uses the package's normal embedded static-data behavior.

</div>

<div class="api-field" id="loadgamedata-startuponly" markdown="1">

### `StartupOnly`

<p><code>bool</code></p>

StartupOnly indicates static data should be loaded during client
construction only.

</div>

<div class="api-field" id="loadgamedata-always" markdown="1">

### `Always`

<p><code>bool</code></p>

Always indicates static data should be refreshed whenever supported by the
caller's workflow.

</div>

<div class="api-field" id="loadgamedata-never" markdown="1">

### `Never`

<p><code>bool</code></p>

Never indicates static data should not be loaded.

</div>

<a id="clantype"></a>

## Clan Type

<p class="api-signature"><span class="api-kind">type</span> <code>clashy.ClanType</code></p>

ClanType describes a clan's join policy.

<p><code>string</code></p>

### Values

<div class="api-field" id="clantypeopen" markdown="1">

#### `ClanTypeOpen`

<p><code>"open"</code></p>

ClanTypeOpen means players can join directly when requirements are met.

</div>

<div class="api-field" id="clantypeclosed" markdown="1">

#### `ClanTypeClosed`

<p><code>"closed"</code></p>

ClanTypeClosed means the clan is closed to new members.

</div>

<div class="api-field" id="clantypeinviteonly" markdown="1">

#### `ClanTypeInviteOnly`

<p><code>"inviteOnly"</code></p>

ClanTypeInviteOnly means players must request or be invited to join.

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


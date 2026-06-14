# Clans

Clan profiles, clan members, clan labels, and clan search models.

<a id="clan"></a>

## Clan

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Clan</code></p>

Clan is the full clan profile returned by GetClan and search endpoints.

<div class="api-field" id="clan-tag" markdown="1">

### `Tag`

<p><code>string</code> <span class="api-json">json: tag</span></p>

Tag is the clan tag.

</div>

<div class="api-field" id="clan-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the clan display name.

</div>

<div class="api-field" id="clan-type" markdown="1">

### `Type`

<p><code><a href="#clantype">ClanType</a></code> <span class="api-json">json: type</span></p>

Type describes whether the clan is open, closed, or invite-only.

</div>

<div class="api-field" id="clan-description" markdown="1">

### `Description`

<p><code>string</code> <span class="api-json">json: description</span></p>

Description is the public clan description.

</div>

<div class="api-field" id="clan-familyfriendly" markdown="1">

### `FamilyFriendly`

<p><code>bool</code> <span class="api-json">json: isFamilyFriendly</span></p>

FamilyFriendly reports whether the clan is marked family friendly.

</div>

<div class="api-field" id="clan-level" markdown="1">

### `Level`

<p><code>int</code> <span class="api-json">json: clanLevel</span></p>

Level is the clan level.

</div>

<div class="api-field" id="clan-points" markdown="1">

### `Points`

<p><code>int</code> <span class="api-json">json: clanPoints</span></p>

Points is the clan's home village trophy score.

</div>

<div class="api-field" id="clan-builderbasepoints" markdown="1">

### `BuilderBasePoints`

<p><code>int</code> <span class="api-json">json: clanBuilderBasePoints</span></p>

BuilderBasePoints is the clan's Builder Base trophy score.

</div>

<div class="api-field" id="clan-capitalpoints" markdown="1">

### `CapitalPoints`

<p><code>int</code> <span class="api-json">json: clanCapitalPoints</span></p>

CapitalPoints is the clan's Clan Capital score.

</div>

<div class="api-field" id="clan-requiredtrophies" markdown="1">

### `RequiredTrophies`

<p><code>int</code> <span class="api-json">json: requiredTrophies</span></p>

RequiredTrophies is the home village trophy requirement to join.

</div>

<div class="api-field" id="clan-warfrequency" markdown="1">

### `WarFrequency`

<p><code>string</code> <span class="api-json">json: warFrequency</span></p>

WarFrequency is the clan's declared war frequency.

</div>

<div class="api-field" id="clan-warwinstreak" markdown="1">

### `WarWinStreak`

<p><code>int</code> <span class="api-json">json: warWinStreak</span></p>

WarWinStreak is the current classic-war win streak.

</div>

<div class="api-field" id="clan-warwins" markdown="1">

### `WarWins`

<p><code>int</code> <span class="api-json">json: warWins</span></p>

WarWins is the number of classic-war wins.

</div>

<div class="api-field" id="clan-warties" markdown="1">

### `WarTies`

<p><code>int</code> <span class="api-json">json: warTies</span></p>

WarTies is the number of classic-war ties when the API includes it.

</div>

<div class="api-field" id="clan-warlosses" markdown="1">

### `WarLosses`

<p><code>int</code> <span class="api-json">json: warLosses</span></p>

WarLosses is the number of classic-war losses when the API includes it.

</div>

<div class="api-field" id="clan-publicwarlog" markdown="1">

### `PublicWarLog`

<p><code>bool</code> <span class="api-json">json: isWarLogPublic</span></p>

PublicWarLog reports whether the clan war log is public.

</div>

<div class="api-field" id="clan-membercount" markdown="1">

### `MemberCount`

<p><code>int</code> <span class="api-json">json: members</span></p>

MemberCount is the number of current clan members.

</div>

<div class="api-field" id="clan-requiredbuilderbasetrophies" markdown="1">

### `RequiredBuilderBaseTrophies`

<p><code>int</code> <span class="api-json">json: requiredBuilderBaseTrophies</span></p>

RequiredBuilderBaseTrophies is the Builder Base trophy requirement to
join.

</div>

<div class="api-field" id="clan-requiredtownhall" markdown="1">

### `RequiredTownhall`

<p><code>int</code> <span class="api-json">json: requiredTownhallLevel</span></p>

RequiredTownhall is the minimum Town Hall level required to join.

</div>

<div class="api-field" id="clan-location" markdown="1">

### `Location`

<p><code>*<a href="../locations-rankings/#location">Location</a></code> <span class="api-json">json: location</span></p>

Location is the clan's declared location.

</div>

<div class="api-field" id="clan-badge" markdown="1">

### `Badge`

<p><code><a href="../locations-rankings/#badge">Badge</a></code> <span class="api-json">json: badgeUrls</span></p>

Badge contains the clan badge image URLs.

</div>

<div class="api-field" id="clan-labels" markdown="1">

### `Labels`

<p><code>[]<a href="../locations-rankings/#label">Label</a></code> <span class="api-json">json: labels</span></p>

Labels are public labels assigned to the clan.

</div>

<div class="api-field" id="clan-members" markdown="1">

### `Members`

<p><code>[]<a href="#clanmember">ClanMember</a></code> <span class="api-json">json: memberList</span></p>

Members is the member list embedded in full clan responses.

</div>

<div class="api-field" id="clan-warleague" markdown="1">

### `WarLeague`

<p><code><a href="../locations-rankings/#league">League</a></code> <span class="api-json">json: warLeague</span></p>

WarLeague is the clan's current Clan War League tier.

</div>

<div class="api-field" id="clan-capitalleague" markdown="1">

### `CapitalLeague`

<p><code>*<a href="../locations-rankings/#league">League</a></code> <span class="api-json">json: capitalLeague</span></p>

CapitalLeague is the clan's Clan Capital league.

</div>

<div class="api-field" id="clan-chatlanguage" markdown="1">

### `ChatLanguage`

<p><code>*<a href="../miscellaneous/#chatlanguage">ChatLanguage</a></code> <span class="api-json">json: chatLanguage</span></p>

ChatLanguage is the clan's preferred chat language.

</div>

<div class="api-field" id="clan-clancapital" markdown="1">

### `ClanCapital`

<p><code>*<a href="#clancapital">ClanCapital</a></code> <span class="api-json">json: clanCapital</span></p>

ClanCapital contains Clan Capital district information.

</div>

<a id="clanmember"></a>

## Clan Member

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ClanMember</code></p>

ClanMember describes a member entry from a clan profile or member list.

<div class="api-field" id="clanmember-tag" markdown="1">

### `Tag`

<p><code>string</code> <span class="api-json">json: tag</span></p>

Tag is the member's player tag.

</div>

<div class="api-field" id="clanmember-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the member's player name.

</div>

<div class="api-field" id="clanmember-role" markdown="1">

### `Role`

<p><code><a href="../enums/#role">Role</a></code> <span class="api-json">json: role</span></p>

Role is the member's clan role.

</div>

<div class="api-field" id="clanmember-explevel" markdown="1">

### `ExpLevel`

<p><code>int</code> <span class="api-json">json: expLevel</span></p>

ExpLevel is the player's experience level.

</div>

<div class="api-field" id="clanmember-townhall" markdown="1">

### `TownHall`

<p><code>int</code> <span class="api-json">json: townHallLevel</span></p>

TownHall is the player's home village Town Hall level.

</div>

<div class="api-field" id="clanmember-trophies" markdown="1">

### `Trophies`

<p><code>int</code> <span class="api-json">json: trophies</span></p>

Trophies is the player's home village trophy count.

</div>

<div class="api-field" id="clanmember-clanrank" markdown="1">

### `ClanRank`

<p><code>int</code> <span class="api-json">json: clanRank</span></p>

ClanRank is the player's current position in the clan trophy ranking.

</div>

<div class="api-field" id="clanmember-clanpreviousrank" markdown="1">

### `ClanPreviousRank`

<p><code>int</code> <span class="api-json">json: previousClanRank</span></p>

ClanPreviousRank is the player's previous position in the clan trophy
ranking.

</div>

<div class="api-field" id="clanmember-donations" markdown="1">

### `Donations`

<p><code>int</code> <span class="api-json">json: donations</span></p>

Donations is the number of troops donated this season.

</div>

<div class="api-field" id="clanmember-received" markdown="1">

### `Received`

<p><code>int</code> <span class="api-json">json: donationsReceived</span></p>

Received is the number of donated troops received this season.

</div>

<div class="api-field" id="clanmember-versustrophies" markdown="1">

### `VersusTrophies`

<p><code>int</code> <span class="api-json">json: versusTrophies</span></p>

VersusTrophies is the legacy Builder Base trophy field used by older API
responses.

</div>

<div class="api-field" id="clanmember-builderbasetrophies" markdown="1">

### `BuilderBaseTrophies`

<p><code>int</code> <span class="api-json">json: builderBaseTrophies</span></p>

BuilderBaseTrophies is the player's Builder Base trophy count.

</div>

<div class="api-field" id="clanmember-versusrank" markdown="1">

### `VersusRank`

<p><code>int</code> <span class="api-json">json: versusRank</span></p>

VersusRank is the legacy Builder Base rank field used by older API
responses.

</div>

<div class="api-field" id="clanmember-builderbaserank" markdown="1">

### `BuilderBaseRank`

<p><code>int</code> <span class="api-json">json: builderBaseRank</span></p>

BuilderBaseRank is the player's current Builder Base rank in the clan.

</div>

<div class="api-field" id="clanmember-leaguetier" markdown="1">

### `LeagueTier`

<p><code><a href="../locations-rankings/#league">League</a></code> <span class="api-json">json: leagueTier</span></p>

LeagueTier is the player's home village league.

</div>

<a id="playerclan"></a>

## Player Clan

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.PlayerClan</code></p>

PlayerClan is the compact clan object embedded in player responses.

<div class="api-field" id="playerclan-tag" markdown="1">

### `Tag`

<p><code>string</code> <span class="api-json">json: tag</span></p>

Tag is the clan tag, including the leading # when returned by the API.

</div>

<div class="api-field" id="playerclan-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the clan display name.

</div>

<div class="api-field" id="playerclan-level" markdown="1">

### `Level`

<p><code>int</code> <span class="api-json">json: clanLevel</span></p>

Level is the clan level.

</div>

<div class="api-field" id="playerclan-badge" markdown="1">

### `Badge`

<p><code><a href="../locations-rankings/#badge">Badge</a></code> <span class="api-json">json: badgeUrls</span></p>

Badge contains the clan badge image URLs.

</div>

<a id="clancapital"></a>

## Clan Capital

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ClanCapital</code></p>

ClanCapital describes a clan's capital districts from the clan profile.

<div class="api-field" id="clancapital-districts" markdown="1">

### `Districts`

<p><code>[]<a href="../raids/#capitaldistrict">CapitalDistrict</a></code> <span class="api-json">json: districts</span></p>

Districts contains the visible Clan Capital districts and their hall
levels.

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

## Clan Methods

<a id="clan-getmember"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Clan.GetMember(<span class="api-param">tag: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*ClanMember</span></code></p>

GetMember returns the member with the provided tag, or nil when the clan
member list does not contain that tag.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>tag</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#clanmember">ClanMember</a></code> </dd>
</dl>

</div>

<a id="clan-getmemberby"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Clan.GetMemberBy(<span class="api-param">name: string</span>, <span class="api-param">trophies: int</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*ClanMember</span></code></p>

GetMemberBy returns the first member matching the provided name and trophy
filters.

Empty name and zero trophies are treated as wildcards, which is useful when a
caller only has one of the two values from an external event.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
<p><strong>trophies</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#clanmember">ClanMember</a></code> </dd>
</dl>

</div>


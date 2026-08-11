# Locations And Rankings

Locations, leagues, seasons, labels, and ranked clan/player wrappers.

<a id="location"></a>

## Location

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Location</code></p>

Location is a country or global location used by ranking endpoints.

<div class="api-field" id="location-id" markdown="1">

### `ID`

<p><code>int</code> <span class="api-json">json: id</span></p>

ID is the numeric location identifier.

</div>

<div class="api-field" id="location-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the English location name.

</div>

<div class="api-field" id="location-iscountry" markdown="1">

### `IsCountry`

<p><code>bool</code> <span class="api-json">json: isCountry</span></p>

IsCountry reports whether the location is a country.

</div>

<div class="api-field" id="location-countrycode" markdown="1">

### `CountryCode`

<p><code>string</code> <span class="api-json">json: countryCode</span></p>

CountryCode is the ISO-style country code when the location is a country.

</div>

<div class="api-field" id="location-localised" markdown="1">

### `Localised`

<p><code>string</code> <span class="api-json">json: localizedName</span></p>

Localised is the API-provided localized display name.

</div>

<a id="league"></a>

## League

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.League</code></p>

League is a league, war league, builder-base league, or capital league.

<div class="api-field" id="league-id" markdown="1">

### `ID`

<p><code>int</code> <span class="api-json">json: id</span></p>

ID is the numeric league identifier.

</div>

<div class="api-field" id="league-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the league display name.

</div>

<div class="api-field" id="league-icon" markdown="1">

### `Icon`

<p><code>*<a href="#icon">Icon</a></code> <span class="api-json">json: iconUrls</span></p>

Icon contains league icon URLs when the endpoint provides them.

</div>

<a id="season"></a>

## Season

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Season</code></p>

Season describes one ranked season placement.

<div class="api-field" id="season-id" markdown="1">

### `ID`

<p><code>string</code> <span class="api-json">json: id</span></p>

ID is the season identifier returned by the API.

</div>

<div class="api-field" id="season-rank" markdown="1">

### `Rank`

<p><code>int</code> <span class="api-json">json: rank</span></p>

Rank is the player's season rank.

</div>

<div class="api-field" id="season-trophies" markdown="1">

### `Trophies`

<p><code>int</code> <span class="api-json">json: trophies</span></p>

Trophies is the player's trophy count for the season.

</div>

<a id="label"></a>

## Label

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Label</code></p>

Label is a player or clan label.

<div class="api-field" id="label-id" markdown="1">

### `ID`

<p><code>int</code> <span class="api-json">json: id</span></p>

ID is the label identifier.

</div>

<div class="api-field" id="label-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the label display name.

</div>

<div class="api-field" id="label-icon" markdown="1">

### `Icon`

<p><code>*<a href="#icon">Icon</a></code> <span class="api-json">json: iconUrls</span></p>

Icon contains label icon URLs.

</div>

<a id="icon"></a>

## Icon

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Icon</code></p>

Icon contains small icon URLs returned for leagues and labels.

<div class="api-field" id="icon-small" markdown="1">

### `Small`

<p><code>string</code> <span class="api-json">json: small</span></p>

Small is the small icon URL.

</div>

<div class="api-field" id="icon-medium" markdown="1">

### `Medium`

<p><code>string</code> <span class="api-json">json: medium</span></p>

Medium is the medium icon URL.

</div>

<div class="api-field" id="icon-tiny" markdown="1">

### `Tiny`

<p><code>string</code> <span class="api-json">json: tiny</span></p>

Tiny is the tiny icon URL.

</div>

<a id="badge"></a>

## Badge

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Badge</code></p>

Badge contains the common small, medium, and large image URLs for clan badges.

<div class="api-field" id="badge-small" markdown="1">

### `Small`

<p><code>string</code> <span class="api-json">json: small</span></p>

Small is the small badge image URL.

</div>

<div class="api-field" id="badge-medium" markdown="1">

### `Medium`

<p><code>string</code> <span class="api-json">json: medium</span></p>

Medium is the medium badge image URL.

</div>

<div class="api-field" id="badge-large" markdown="1">

### `Large`

<p><code>string</code> <span class="api-json">json: large</span></p>

Large is the large badge image URL.

</div>

<a id="goldpassseason"></a>

## Gold Pass Season

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.GoldPassSeason</code></p>

GoldPassSeason describes the current Gold Pass season.

<div class="api-field" id="goldpassseason-starttime" markdown="1">

### `StartTime`

<p><code>*<a href="../miscellaneous/#timestamp">Timestamp</a></code> <span class="api-json">json: startTime</span></p>

StartTime is when the Gold Pass season starts.

</div>

<div class="api-field" id="goldpassseason-endtime" markdown="1">

### `EndTime`

<p><code>*<a href="../miscellaneous/#timestamp">Timestamp</a></code> <span class="api-json">json: endTime</span></p>

EndTime is when the Gold Pass season ends.

</div>

<a id="rankedclan"></a>

## Ranked Clan

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.RankedClan</code></p>

RankedClan is a clan ranking entry.

<div class="api-field" id="rankedclan-clan" markdown="1">

### `Clan`

<p><code><a href="../clans/#clan">Clan</a></code></p>

</div>

<div class="api-field" id="rankedclan-rank" markdown="1">

### `Rank`

<p><code>int</code> <span class="api-json">json: rank</span></p>

Rank is the current ranking position.

</div>

<div class="api-field" id="rankedclan-previousrank" markdown="1">

### `PreviousRank`

<p><code>int</code> <span class="api-json">json: previousRank</span></p>

PreviousRank is the previous ranking position when the API provides it.

</div>

<a id="rankedplayer"></a>

## Ranked Player

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.RankedPlayer</code></p>

RankedPlayer is a player ranking entry.

<div class="api-field" id="rankedplayer-player" markdown="1">

### `Player`

<p><code><a href="../players/#player">Player</a></code></p>

</div>

<div class="api-field" id="rankedplayer-league" markdown="1">

### `League`

<p><code><a href="#league">League</a></code> <span class="api-json">json: league</span></p>

League is the player's ranking league when the endpoint includes it.

</div>

<div class="api-field" id="rankedplayer-attackwins" markdown="1">

### `AttackWins`

<p><code>int</code> <span class="api-json">json: attackWins</span></p>

AttackWins is the player's attack win count in the ranking.

</div>

<div class="api-field" id="rankedplayer-defensewins" markdown="1">

### `DefenseWins`

<p><code>int</code> <span class="api-json">json: defenseWins</span></p>

DefenseWins is the player's defense win count in the ranking.

</div>

<div class="api-field" id="rankedplayer-rank" markdown="1">

### `Rank`

<p><code>int</code> <span class="api-json">json: rank</span></p>

Rank is the current ranking position.

</div>

<div class="api-field" id="rankedplayer-previousrank" markdown="1">

### `PreviousRank`

<p><code>int</code> <span class="api-json">json: previousRank</span></p>

PreviousRank is the previous ranking position when the API provides it.

</div>


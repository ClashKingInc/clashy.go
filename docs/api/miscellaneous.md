# Miscellaneous

Shared timestamp, season, date, and utility helpers.

<a id="timestamp"></a>

## Timestamp

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Timestamp</code></p>

Timestamp stores both the raw Clash API timestamp string and its parsed time.

<div class="api-field" id="timestamp-rawtime" markdown="1">

### `RawTime`

<p><code>string</code></p>

RawTime is the original API timestamp.

	20060102T150405.000Z

</div>

<div class="api-field" id="timestamp-time" markdown="1">

### `Time`

<p><code>time.Time</code></p>

Time is the parsed UTC time.

</div>

<a id="timedelta"></a>

## Time Delta

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.TimeDelta</code></p>

TimeDelta represents an elapsed duration.

<a id="seasonwindow"></a>

## Season Window

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.SeasonWindow</code></p>

<div class="api-field" id="seasonwindow-seasonid" markdown="1">

### `SeasonID`

<p><code>string</code></p>

SeasonID is the season identifier in YYYY-MM form.

</div>

<div class="api-field" id="seasonwindow-starttime" markdown="1">

### `StartTime`

<p><code>time.Time</code></p>

StartTime is the inclusive season start time in UTC.

</div>

<div class="api-field" id="seasonwindow-endtime" markdown="1">

### `EndTime`

<p><code>time.Time</code></p>

EndTime is the exclusive season end time in UTC.

</div>

<a id="chatlanguage"></a>

## Chat Language

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ChatLanguage</code></p>

ChatLanguage describes the preferred language configured for a clan.

<div class="api-field" id="chatlanguage-id" markdown="1">

### `ID`

<p><code>int</code> <span class="api-json">json: id</span></p>

ID is the language identifier.

</div>

<div class="api-field" id="chatlanguage-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the language display name.

</div>

<div class="api-field" id="chatlanguage-languagecode" markdown="1">

### `LanguageCode`

<p><code>string</code> <span class="api-json">json: languageCode</span></p>

LanguageCode is the language code returned by the API.

</div>

## Timestamp Methods

<a id="timestamp-after"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Timestamp.After(<span class="api-param">other: Timestamp</span>)<span class="api-return-arrow"> -> </span><span class="api-return">bool</span></code></p>

After reports whether this timestamp occurs after another timestamp.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>other</strong> (<code><a href="#timestamp">Timestamp</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>bool</code> </dd>
</dl>

</div>

<a id="timestamp-before"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Timestamp.Before(<span class="api-param">other: Timestamp</span>)<span class="api-return-arrow"> -> </span><span class="api-return">bool</span></code></p>

Before reports whether this timestamp occurs before another timestamp.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>other</strong> (<code><a href="#timestamp">Timestamp</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>bool</code> </dd>
</dl>

</div>

<a id="timestamp-secondsuntil"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Timestamp.SecondsUntil()<span class="api-return-arrow"> -> </span><span class="api-return">int</span></code></p>

SecondsUntil returns the number of whole seconds from now until the timestamp.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>int</code> </dd>
</dl>

</div>

<a id="timestamp-unmarshaljson"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Timestamp.UnmarshalJSON(<span class="api-param">data: []byte</span>)<span class="api-return-arrow"> -> </span><span class="api-return">error</span></code></p>

UnmarshalJSON parses Clash API timestamp strings into Timestamp values.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>data</strong> (<code>[]byte</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>error</code> </dd>
</dl>

</div>

## Functions

<a id="correcttag"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.CorrectTag(<span class="api-param">tag: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">string</span></code></p>

CorrectTag normalizes a Clash tag by trimming whitespace, uppercasing,
replacing O with 0, removing invalid characters, and ensuring a leading #.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>tag</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>string</code> </dd>
</dl>

</div>

<a id="fromtimestamp"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.FromTimestamp(<span class="api-param">raw: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">time.Time</span>, <span class="api-return">error</span>)</code></p>

FromTimestamp parses a Clash API timestamp in 20060102T150405.000Z format.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>raw</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>time.Time</code> <code>error</code> </dd>
</dl>

</div>

<a id="getseasonid"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GetSeasonID()<span class="api-return-arrow"> -> </span><span class="api-return">string</span></code></p>

GetSeasonID returns the current trophy season identifier in YYYY-MM form.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>string</code> </dd>
</dl>

</div>

<a id="genseasondate"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GenSeasonDate(<span class="api-param">timestamp: time.Time</span>)<span class="api-return-arrow"> -> </span><span class="api-return">string</span></code></p>

GenSeasonDate returns the trophy season identifier for timestamp.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>timestamp</strong> (<code>time.Time</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>string</code> </dd>
</dl>

</div>

<a id="genlegenddate"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GenLegendDate(<span class="api-param">timestamp: time.Time</span>)<span class="api-return-arrow"> -> </span><span class="api-return">string</span></code></p>

GenLegendDate returns the legend-league day identifier for timestamp.

Legend days roll over at 05:00 UTC, so timestamps before that hour map to the
previous calendar date.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>timestamp</strong> (<code>time.Time</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>string</code> </dd>
</dl>

</div>

<a id="getseasonstart"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GetSeasonStart(<span class="api-param">timestamp: time.Time</span>)<span class="api-return-arrow"> -> </span><span class="api-return">time.Time</span></code></p>

GetSeasonStart returns the start time of the trophy season containing
timestamp.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>timestamp</strong> (<code>time.Time</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>time.Time</code> </dd>
</dl>

</div>

<a id="getseasonend"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GetSeasonEnd(<span class="api-param">timestamp: time.Time</span>)<span class="api-return-arrow"> -> </span><span class="api-return">time.Time</span></code></p>

GetSeasonEnd returns the end time of the trophy season containing timestamp.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>timestamp</strong> (<code>time.Time</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>time.Time</code> </dd>
</dl>

</div>

<a id="getseason"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GetSeason(<span class="api-param">timestamp: time.Time</span>, <span class="api-param">forward: bool</span>)<span class="api-return-arrow"> -> </span><span class="api-return">SeasonWindow</span></code></p>

GetSeason returns the trophy season window containing timestamp.

Passing a zero timestamp uses the current UTC time. Before the 2025 season
calendar change, seasons end on the last Monday of the month at 05:00 UTC.
From the September 2025 transition onward, seasons follow fixed 28 day
windows.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>timestamp</strong> (<code>time.Time</code>)</p>
<p><strong>forward</strong> (<code>bool</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code><a href="#seasonwindow">SeasonWindow</a></code> </dd>
</dl>

</div>

<a id="getseasonbyid"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GetSeasonByID(<span class="api-param">seasonID: string</span>)<span class="api-return-arrow"> -> </span>(<span class="api-return">SeasonWindow</span>, <span class="api-return">error</span>)</code></p>

GetSeasonByID returns the trophy season window for a YYYY-MM season ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>seasonID</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code><a href="#seasonwindow">SeasonWindow</a></code> <code>error</code> </dd>
</dl>

</div>

<a id="getclangamesstart"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GetClanGamesStart(<span class="api-param">timestamp: time.Time</span>)<span class="api-return-arrow"> -> </span><span class="api-return">time.Time</span></code></p>

GetClanGamesStart returns the Clan Games start time for the month containing
timestamp, rolling forward after that month's Clan Games end.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>timestamp</strong> (<code>time.Time</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>time.Time</code> </dd>
</dl>

</div>

<a id="getclangamesend"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GetClanGamesEnd(<span class="api-param">timestamp: time.Time</span>)<span class="api-return-arrow"> -> </span><span class="api-return">time.Time</span></code></p>

GetClanGamesEnd returns the Clan Games end time for the month containing
timestamp, rolling forward after that month's Clan Games end.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>timestamp</strong> (<code>time.Time</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>time.Time</code> </dd>
</dl>

</div>

<a id="getraidweekendstart"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GetRaidWeekendStart(<span class="api-param">timestamp: time.Time</span>)<span class="api-return-arrow"> -> </span><span class="api-return">time.Time</span></code></p>

GetRaidWeekendStart returns the start time for the raid weekend containing
timestamp.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>timestamp</strong> (<code>time.Time</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>time.Time</code> </dd>
</dl>

</div>

<a id="getraidweekendend"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.GetRaidWeekendEnd(<span class="api-param">timestamp: time.Time</span>)<span class="api-return-arrow"> -> </span><span class="api-return">time.Time</span></code></p>

GetRaidWeekendEnd returns the end time for the raid weekend containing
timestamp.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>timestamp</strong> (<code>time.Time</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>time.Time</code> </dd>
</dl>

</div>

<a id="withoutratelimit"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.WithoutRateLimit(<span class="api-param">ctx: context.Context</span>)<span class="api-return-arrow"> -> </span><span class="api-return">context.Context</span></code></p>

WithoutRateLimit returns a child context that bypasses the client's request
limiter.

Use this for trusted internal calls where the caller is already controlling
concurrency. It does not disable token rotation, caching, deadlines, or HTTP
transport behavior.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>ctx</strong> (<code>context.Context</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>context.Context</code> </dd>
</dl>

</div>


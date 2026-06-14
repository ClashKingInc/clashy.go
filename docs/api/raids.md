# Raids

Clan Capital raid weekend logs, districts, attacks, and member totals.

<a id="raidlogentry"></a>

## Raid Log Entry

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.RaidLogEntry</code></p>

RaidLogEntry is one Clan Capital raid weekend log entry.

<div class="api-field" id="raidlogentry-state" markdown="1">

### `State`

<p><code>string</code> <span class="api-json">json: state</span></p>

State is the raid weekend state.

</div>

<div class="api-field" id="raidlogentry-totalloot" markdown="1">

### `TotalLoot`

<p><code>int</code> <span class="api-json">json: capitalTotalLoot</span></p>

TotalLoot is the clan's total capital gold looted.

</div>

<div class="api-field" id="raidlogentry-completedraidcount" markdown="1">

### `CompletedRaidCount`

<p><code>int</code> <span class="api-json">json: raidsCompleted</span></p>

CompletedRaidCount is the number of completed raids.

</div>

<div class="api-field" id="raidlogentry-attackcount" markdown="1">

### `AttackCount`

<p><code>int</code> <span class="api-json">json: totalAttacks</span></p>

AttackCount is the total number of attacks used by the clan.

</div>

<div class="api-field" id="raidlogentry-destroyeddistrictcount" markdown="1">

### `DestroyedDistrictCount`

<p><code>int</code> <span class="api-json">json: enemyDistrictsDestroyed</span></p>

DestroyedDistrictCount is the number of enemy districts destroyed.

</div>

<div class="api-field" id="raidlogentry-offensivereward" markdown="1">

### `OffensiveReward`

<p><code>int</code> <span class="api-json">json: offensiveReward</span></p>

OffensiveReward is the offensive raid medal reward.

</div>

<div class="api-field" id="raidlogentry-defensivereward" markdown="1">

### `DefensiveReward`

<p><code>int</code> <span class="api-json">json: defensiveReward</span></p>

DefensiveReward is the defensive raid medal reward.

</div>

<div class="api-field" id="raidlogentry-starttime" markdown="1">

### `StartTime`

<p><code>*<a href="../miscellaneous/#timestamp">Timestamp</a></code> <span class="api-json">json: startTime</span></p>

StartTime is when the raid weekend started.

</div>

<div class="api-field" id="raidlogentry-endtime" markdown="1">

### `EndTime`

<p><code>*<a href="../miscellaneous/#timestamp">Timestamp</a></code> <span class="api-json">json: endTime</span></p>

EndTime is when the raid weekend ended.

</div>

<div class="api-field" id="raidlogentry-attacklog" markdown="1">

### `AttackLog`

<p><code>[]<a href="#raidclan">RaidClan</a></code> <span class="api-json">json: attackLog</span></p>

AttackLog contains raids made by the requested clan.

</div>

<div class="api-field" id="raidlogentry-defenselog" markdown="1">

### `DefenseLog`

<p><code>[]<a href="#raidclan">RaidClan</a></code> <span class="api-json">json: defenseLog</span></p>

DefenseLog contains raids made against the requested clan.

</div>

<div class="api-field" id="raidlogentry-members" markdown="1">

### `Members`

<p><code>[]<a href="#raidmember">RaidMember</a></code> <span class="api-json">json: members</span></p>

Members contains member-level attack and loot totals.

</div>

<a id="raidclan"></a>

## Raid Clan

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.RaidClan</code></p>

RaidClan describes one opposing clan entry in a raid attack or defense log.

Attack log entries use Attacker for the clan being attacked by the requested
clan. Defense log entries use Defender for the clan that attacked the
requested clan.

<div class="api-field" id="raidclan-attackcount" markdown="1">

### `AttackCount`

<p><code>int</code> <span class="api-json">json: attackCount</span></p>

AttackCount is the number of attacks used in this raid.

</div>

<div class="api-field" id="raidclan-districtcount" markdown="1">

### `DistrictCount`

<p><code>int</code> <span class="api-json">json: districtCount</span></p>

DistrictCount is the number of districts available.

</div>

<div class="api-field" id="raidclan-destroyeddistrictcount" markdown="1">

### `DestroyedDistrictCount`

<p><code>int</code> <span class="api-json">json: districtsDestroyed</span></p>

DestroyedDistrictCount is the number of districts destroyed.

</div>

<div class="api-field" id="raidclan-districts" markdown="1">

### `Districts`

<p><code>[]<a href="#raiddistrict">RaidDistrict</a></code> <span class="api-json">json: districts</span></p>

Districts contains district-level attack details.

</div>

<div class="api-field" id="raidclan-attacker" markdown="1">

### `Attacker`

<p><code>*struct {
	Tag	string	`json:"tag,omitempty"`
	Name	string	`json:"name,omitempty"`
	Level	int	`json:"level,omitempty"`
}</code> <span class="api-json">json: attacker</span></p>

Attacker is set on attack-log entries.

</div>

<div class="api-field" id="raidclan-defender" markdown="1">

### `Defender`

<p><code>*struct {
	Tag	string	`json:"tag,omitempty"`
	Name	string	`json:"name,omitempty"`
	Level	int	`json:"level,omitempty"`
}</code> <span class="api-json">json: defender</span></p>

Defender is set on defense-log entries.

</div>

<a id="raiddistrict"></a>

## Raid District

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.RaidDistrict</code></p>

RaidDistrict describes one district in a raid attack or defense log.

<div class="api-field" id="raiddistrict-id" markdown="1">

### `ID`

<p><code>int</code> <span class="api-json">json: id</span></p>

ID is the district identifier.

</div>

<div class="api-field" id="raiddistrict-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the district display name.

</div>

<div class="api-field" id="raiddistrict-halllevel" markdown="1">

### `HallLevel`

<p><code>int</code> <span class="api-json">json: districtHallLevel</span></p>

HallLevel is the district hall level.

</div>

<div class="api-field" id="raiddistrict-destruction" markdown="1">

### `Destruction`

<p><code>float64</code> <span class="api-json">json: destructionPercent</span></p>

Destruction is the final destruction percentage for the district.

</div>

<div class="api-field" id="raiddistrict-attackcount" markdown="1">

### `AttackCount`

<p><code>int</code> <span class="api-json">json: attackCount</span></p>

AttackCount is the number of attacks used against the district.

</div>

<div class="api-field" id="raiddistrict-looted" markdown="1">

### `Looted`

<p><code>int</code> <span class="api-json">json: totalLooted</span></p>

Looted is the total capital gold looted from the district.

</div>

<div class="api-field" id="raiddistrict-attacks" markdown="1">

### `Attacks`

<p><code>[]<a href="#raidattack">RaidAttack</a></code> <span class="api-json">json: attacks</span></p>

Attacks contains individual attacks against the district.

</div>

<a id="raidattack"></a>

## Raid Attack

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.RaidAttack</code></p>

RaidAttack is one attack against a Clan Capital district.

<div class="api-field" id="raidattack-attackertag" markdown="1">

### `AttackerTag`

<p><code>string</code></p>

AttackerTag can be populated by callers that join attacks to raid members.

</div>

<div class="api-field" id="raidattack-attackername" markdown="1">

### `AttackerName`

<p><code>string</code></p>

AttackerName can be populated by callers that join attacks to raid
members.

</div>

<div class="api-field" id="raidattack-stars" markdown="1">

### `Stars`

<p><code>int</code> <span class="api-json">json: stars</span></p>

Stars is the star count earned by this district attack.

</div>

<div class="api-field" id="raidattack-destruction" markdown="1">

### `Destruction`

<p><code>float64</code> <span class="api-json">json: destructionPercent</span></p>

Destruction is the destruction percentage earned by this district attack.

</div>

<a id="raidmember"></a>

## Raid Member

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.RaidMember</code></p>

RaidMember is one clan member's contribution in a raid weekend.

<div class="api-field" id="raidmember-tag" markdown="1">

### `Tag`

<p><code>string</code> <span class="api-json">json: tag</span></p>

Tag is the member's player tag.

</div>

<div class="api-field" id="raidmember-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the member's display name.

</div>

<div class="api-field" id="raidmember-attackcount" markdown="1">

### `AttackCount`

<p><code>int</code> <span class="api-json">json: attacks</span></p>

AttackCount is the number of attacks used.

</div>

<div class="api-field" id="raidmember-attacklimit" markdown="1">

### `AttackLimit`

<p><code>int</code> <span class="api-json">json: attackLimit</span></p>

AttackLimit is the normal attack limit.

</div>

<div class="api-field" id="raidmember-bonusattacklimit" markdown="1">

### `BonusAttackLimit`

<p><code>int</code> <span class="api-json">json: bonusAttackLimit</span></p>

BonusAttackLimit is the number of bonus attacks available.

</div>

<div class="api-field" id="raidmember-capitalresourceslooted" markdown="1">

### `CapitalResourcesLooted`

<p><code>int</code> <span class="api-json">json: capitalResourcesLooted</span></p>

CapitalResourcesLooted is the capital gold looted by the member.

</div>

<a id="capitaldistrict"></a>

## Capital District

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.CapitalDistrict</code></p>

CapitalDistrict describes a clan capital district from clan and raid data.

<div class="api-field" id="capitaldistrict-id" markdown="1">

### `ID`

<p><code>int</code> <span class="api-json">json: id</span></p>

ID is the district identifier.

</div>

<div class="api-field" id="capitaldistrict-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the district display name.

</div>

<div class="api-field" id="capitaldistrict-districthalllevel" markdown="1">

### `DistrictHallLevel`

<p><code>int</code> <span class="api-json">json: districtHallLevel</span></p>

DistrictHallLevel is the district hall level.

</div>

<div class="api-field" id="capitaldistrict-destructionpercent" markdown="1">

### `DestructionPercent`

<p><code>float64</code> <span class="api-json">json: destructionPercent</span></p>

DestructionPercent is the destruction percentage in raid contexts.

</div>

<div class="api-field" id="capitaldistrict-attackcount" markdown="1">

### `AttackCount`

<p><code>int</code> <span class="api-json">json: attackCount</span></p>

AttackCount is the number of attacks used against the district.

</div>

<div class="api-field" id="capitaldistrict-looted" markdown="1">

### `Looted`

<p><code>int</code> <span class="api-json">json: totalLooted</span></p>

Looted is the total capital gold looted from the district.

</div>

## Raid Log Entry Methods

<a id="raidlogentry-getmember"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.RaidLogEntry.GetMember(<span class="api-param">tag: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*RaidMember</span></code></p>

GetMember returns the raid member with the provided tag.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>tag</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#raidmember">RaidMember</a></code> </dd>
</dl>

</div>

## Raid Clan Methods

<a id="raidclan-level"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.RaidClan.Level()<span class="api-return-arrow"> -> </span><span class="api-return">int</span></code></p>

Level returns the attacker or defender clan level for this raid clan entry.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>int</code> </dd>
</dl>

</div>

<a id="raidclan-name"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.RaidClan.Name()<span class="api-return-arrow"> -> </span><span class="api-return">string</span></code></p>

Name returns the attacker or defender name for this raid clan entry.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>string</code> </dd>
</dl>

</div>

<a id="raidclan-tag"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.RaidClan.Tag()<span class="api-return-arrow"> -> </span><span class="api-return">string</span></code></p>

Tag returns the attacker or defender tag for this raid clan entry.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>string</code> </dd>
</dl>

</div>


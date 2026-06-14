# Game Objects

Troops, spells, heroes, pets, equipment, and parsed army-link models.

<a id="staticunit"></a>

## Static Unit

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.StaticUnit</code></p>

StaticUnit contains normalized static-data fields shared by troops, spells,
heroes, pets, and hero equipment.

<div class="api-field" id="staticunit-name" markdown="1">

### `Name`

<p><code>string</code></p>

Name is the unit or equipment display name.

</div>

<div class="api-field" id="staticunit-level" markdown="1">

### `Level`

<p><code>int</code></p>

Level is the selected level for this static lookup.

</div>

<div class="api-field" id="staticunit-maxlevel" markdown="1">

### `MaxLevel`

<p><code>int</code></p>

MaxLevel is the maximum level found in static data.

</div>

<div class="api-field" id="staticunit-village" markdown="1">

### `Village`

<p><code>string</code></p>

Village identifies the village this object belongs to.

</div>

<div class="api-field" id="staticunit-upgradecost" markdown="1">

### `UpgradeCost`

<p><code>int</code></p>

UpgradeCost is the cost for the selected level when static data includes
it.

</div>

<div class="api-field" id="staticunit-upgradetime" markdown="1">

### `UpgradeTime`

<p><code>time.Duration</code></p>

UpgradeTime is the upgrade duration for the selected level when static data
includes it.

</div>

<a id="troop"></a>

## Troop

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Troop</code></p>

Troop is a player troop or static troop lookup result.

<div class="api-field" id="troop-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the troop display name.

</div>

<div class="api-field" id="troop-level" markdown="1">

### `Level`

<p><code>int</code> <span class="api-json">json: level</span></p>

Level is the player's current level or the selected static level.

</div>

<div class="api-field" id="troop-maxlevel" markdown="1">

### `MaxLevel`

<p><code>int</code> <span class="api-json">json: maxLevel</span></p>

MaxLevel is the maximum level available for the player's Town Hall or in
static data.

</div>

<div class="api-field" id="troop-village" markdown="1">

### `Village`

<p><code>string</code> <span class="api-json">json: village</span></p>

Village identifies home or Builder Base troops.

</div>

<div class="api-field" id="troop-supertroopisactive" markdown="1">

### `SuperTroopIsActive`

<p><code>bool</code> <span class="api-json">json: superTroopIsActive</span></p>

SuperTroopIsActive reports whether a super troop boost is active.

</div>

<div class="api-field" id="troop-staticunit" markdown="1">

### `StaticUnit`

<p><code><a href="#staticunit">StaticUnit</a></code></p>

</div>

<a id="spell"></a>

## Spell

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Spell</code></p>

Spell is a player spell or static spell lookup result.

<div class="api-field" id="spell-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the spell display name.

</div>

<div class="api-field" id="spell-level" markdown="1">

### `Level`

<p><code>int</code> <span class="api-json">json: level</span></p>

Level is the player's current level or the selected static level.

</div>

<div class="api-field" id="spell-maxlevel" markdown="1">

### `MaxLevel`

<p><code>int</code> <span class="api-json">json: maxLevel</span></p>

MaxLevel is the maximum level available for the player or in static data.

</div>

<div class="api-field" id="spell-village" markdown="1">

### `Village`

<p><code>string</code> <span class="api-json">json: village</span></p>

Village identifies the spell's village when static data provides one.

</div>

<div class="api-field" id="spell-staticunit" markdown="1">

### `StaticUnit`

<p><code><a href="#staticunit">StaticUnit</a></code></p>

</div>

<a id="hero"></a>

## Hero

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Hero</code></p>

Hero is a player hero or static hero lookup result.

<div class="api-field" id="hero-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the hero display name.

</div>

<div class="api-field" id="hero-level" markdown="1">

### `Level`

<p><code>int</code> <span class="api-json">json: level</span></p>

Level is the player's current level or the selected static level.

</div>

<div class="api-field" id="hero-maxlevel" markdown="1">

### `MaxLevel`

<p><code>int</code> <span class="api-json">json: maxLevel</span></p>

MaxLevel is the maximum level available for the player or in static data.

</div>

<div class="api-field" id="hero-village" markdown="1">

### `Village`

<p><code>string</code> <span class="api-json">json: village</span></p>

Village identifies the hero's village.

</div>

<div class="api-field" id="hero-equipment" markdown="1">

### `Equipment`

<p><code>[]<a href="#equipment">Equipment</a></code> <span class="api-json">json: equipment</span></p>

Equipment contains equipment currently assigned to this hero when the API
includes loadout data.

</div>

<div class="api-field" id="hero-staticunit" markdown="1">

### `StaticUnit`

<p><code><a href="#staticunit">StaticUnit</a></code></p>

</div>

<a id="pet"></a>

## Pet

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Pet</code></p>

Pet is a hero pet from a player response or static lookup.

<div class="api-field" id="pet-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the pet display name.

</div>

<div class="api-field" id="pet-level" markdown="1">

### `Level`

<p><code>int</code> <span class="api-json">json: level</span></p>

Level is the player's current level or the selected static level.

</div>

<div class="api-field" id="pet-maxlevel" markdown="1">

### `MaxLevel`

<p><code>int</code> <span class="api-json">json: maxLevel</span></p>

MaxLevel is the maximum level available for the player or in static data.

</div>

<div class="api-field" id="pet-village" markdown="1">

### `Village`

<p><code>string</code> <span class="api-json">json: village</span></p>

Village identifies the pet's village.

</div>

<div class="api-field" id="pet-staticunit" markdown="1">

### `StaticUnit`

<p><code><a href="#staticunit">StaticUnit</a></code></p>

</div>

<a id="equipment"></a>

## Equipment

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Equipment</code></p>

Equipment is hero equipment from a player response or static lookup.

<div class="api-field" id="equipment-name" markdown="1">

### `Name`

<p><code>string</code> <span class="api-json">json: name</span></p>

Name is the equipment display name.

</div>

<div class="api-field" id="equipment-level" markdown="1">

### `Level`

<p><code>int</code> <span class="api-json">json: level</span></p>

Level is the player's current level or the selected static level.

</div>

<div class="api-field" id="equipment-maxlevel" markdown="1">

### `MaxLevel`

<p><code>int</code> <span class="api-json">json: maxLevel</span></p>

MaxLevel is the maximum level available for the player or in static data.

</div>

<div class="api-field" id="equipment-village" markdown="1">

### `Village`

<p><code>string</code> <span class="api-json">json: village</span></p>

Village identifies the equipment's village.

</div>

<div class="api-field" id="equipment-rarity" markdown="1">

### `Rarity`

<p><code>string</code> <span class="api-json">json: rarity</span></p>

Rarity is the equipment rarity when static data includes it.

</div>

<div class="api-field" id="equipment-staticunit" markdown="1">

### `StaticUnit`

<p><code><a href="#staticunit">StaticUnit</a></code></p>

</div>

<a id="heroloadout"></a>

## Hero Loadout

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.HeroLoadout</code></p>

HeroLoadout is one hero, pet, and equipment grouping parsed from an army link.

<div class="api-field" id="heroloadout-hero" markdown="1">

### `Hero`

<p><code><a href="#hero">Hero</a></code></p>

Hero is the hero selected in the army link.

</div>

<div class="api-field" id="heroloadout-pet" markdown="1">

### `Pet`

<p><code>*<a href="#pet">Pet</a></code></p>

Pet is the assigned pet when the link includes one.

</div>

<div class="api-field" id="heroloadout-equipment" markdown="1">

### `Equipment`

<p><code>[]<a href="#equipment">Equipment</a></code></p>

Equipment is the selected hero equipment in link order.

</div>

<a id="armyrecipe"></a>

## Army Recipe

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ArmyRecipe</code></p>

ArmyRecipe is the normalized representation of a Clash army link.

<div class="api-field" id="armyrecipe-link" markdown="1">

### `Link`

<p><code>string</code></p>

Link is the original link or raw army payload passed by the caller.

</div>

<div class="api-field" id="armyrecipe-heroesloadout" markdown="1">

### `HeroesLoadout`

<p><code>[]<a href="#heroloadout">HeroLoadout</a></code></p>

HeroesLoadout contains heroes, pets, and equipment from the link.

</div>

<div class="api-field" id="armyrecipe-troops" markdown="1">

### `Troops`

<p><code>[]<a href="#troopcount">TroopCount</a></code></p>

Troops contains home-army troops from the link.

</div>

<div class="api-field" id="armyrecipe-spells" markdown="1">

### `Spells`

<p><code>[]<a href="#spellcount">SpellCount</a></code></p>

Spells contains home-army spells from the link.

</div>

<div class="api-field" id="armyrecipe-clancastletroops" markdown="1">

### `ClanCastleTroops`

<p><code>[]<a href="#troopcount">TroopCount</a></code></p>

ClanCastleTroops contains requested Clan Castle troops.

</div>

<div class="api-field" id="armyrecipe-clancastlespells" markdown="1">

### `ClanCastleSpells`

<p><code>[]<a href="#spellcount">SpellCount</a></code></p>

ClanCastleSpells contains requested Clan Castle spells.

</div>

<a id="troopcount"></a>

## Troop Count

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.TroopCount</code></p>

TroopCount pairs a troop with a quantity from an army link.

<div class="api-field" id="troopcount-troop" markdown="1">

### `Troop`

<p><code><a href="#troop">Troop</a></code></p>

Troop is the parsed troop.

</div>

<div class="api-field" id="troopcount-quantity" markdown="1">

### `Quantity`

<p><code>int</code></p>

Quantity is the requested troop count.

</div>

<a id="spellcount"></a>

## Spell Count

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.SpellCount</code></p>

SpellCount pairs a spell with a quantity from an army link.

<div class="api-field" id="spellcount-spell" markdown="1">

### `Spell`

<p><code><a href="#spell">Spell</a></code></p>

Spell is the parsed spell.

</div>

<div class="api-field" id="spellcount-quantity" markdown="1">

### `Quantity`

<p><code>int</code></p>

Quantity is the requested spell count.

</div>

<a id="accountdata"></a>

## Account Data

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.AccountData</code></p>

AccountData is a thin wrapper around arbitrary account-link data.

<div class="api-field" id="accountdata-raw" markdown="1">

### `Raw`

<p><code>map[string]any</code></p>

Raw contains the original account-link payload.

</div>

## Troop Methods

<a id="troop-isbuilderbase"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Troop.IsBuilderBase()<span class="api-return-arrow"> -> </span><span class="api-return">bool</span></code></p>

IsBuilderBase reports whether the troop belongs to Builder Base.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>bool</code> </dd>
</dl>

</div>

<a id="troop-ishomebase"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Troop.IsHomeBase()<span class="api-return-arrow"> -> </span><span class="api-return">bool</span></code></p>

IsHomeBase reports whether the troop belongs to the home village.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>bool</code> </dd>
</dl>

</div>

<a id="troop-issupertroop"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Troop.IsSuperTroop()<span class="api-return-arrow"> -> </span><span class="api-return">bool</span></code></p>

IsSuperTroop reports whether the troop name is one of the known super troops.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>bool</code> </dd>
</dl>

</div>

<a id="troop-static"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Troop.Static(<span class="api-param">c: *Client</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Troop</span></code></p>

Static returns the embedded static-data record matching this troop's name,
village, and level.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>c</strong> (<code>*<a href="../client/#client">Client</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#troop">Troop</a></code> </dd>
</dl>

</div>

## Spell Methods

<a id="spell-static"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Spell.Static(<span class="api-param">c: *Client</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Spell</span></code></p>

Static returns the embedded static-data record matching this spell's name and
level.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>c</strong> (<code>*<a href="../client/#client">Client</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#spell">Spell</a></code> </dd>
</dl>

</div>

## Hero Methods

<a id="hero-static"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Hero.Static(<span class="api-param">c: *Client</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Hero</span></code></p>

Static returns the embedded static-data record matching this hero's name and
level.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>c</strong> (<code>*<a href="../client/#client">Client</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#hero">Hero</a></code> </dd>
</dl>

</div>

## Pet Methods

<a id="pet-static"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Pet.Static(<span class="api-param">c: *Client</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Pet</span></code></p>

Static returns the embedded static-data record matching this pet's name and
level.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>c</strong> (<code>*<a href="../client/#client">Client</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#pet">Pet</a></code> </dd>
</dl>

</div>

## Equipment Methods

<a id="equipment-static"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.Equipment.Static(<span class="api-param">c: *Client</span>)<span class="api-return-arrow"> -> </span><span class="api-return">*Equipment</span></code></p>

Static returns the embedded static-data record matching this equipment's name
and level.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>c</strong> (<code>*<a href="../client/#client">Client</a></code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#equipment">Equipment</a></code> </dd>
</dl>

</div>

## Functions

<a id="parsearmyrecipe"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.ParseArmyRecipe(<span class="api-param">static: *StaticData</span>, <span class="api-param">link: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">ArmyRecipe</span></code></p>

ParseArmyRecipe parses a full Clash army link or raw army payload into a
structured recipe using embedded static data for names and villages.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>static</strong> (<code>*<a href="../static-data/#staticdata">StaticData</a></code>)</p>
<p><strong>link</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code><a href="#armyrecipe">ArmyRecipe</a></code> </dd>
</dl>

</div>

<a id="parseaccountdata"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.ParseAccountData(<span class="api-param">data: map[string]any</span>)<span class="api-return-arrow"> -> </span><span class="api-return">AccountData</span></code></p>

ParseAccountData wraps account-link data without mutating it.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>data</strong> (<code>map[string]any</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code><a href="#accountdata">AccountData</a></code> </dd>
</dl>

</div>


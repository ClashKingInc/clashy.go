# Static Data

Embedded ClashKing static data, translations, and lookup helpers.

<a id="staticdata"></a>

## Static Data

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.StaticData</code></p>

StaticData is the parsed and indexed ClashKing static data embedded in the
package.

<div class="api-field" id="staticdata-raw" markdown="1">

### `Raw`

<p><code>map[string][]map[string]any</code></p>

Raw preserves static-data sections exactly as parsed from static_data.json.

</div>

<div class="api-field" id="staticdata-byid" markdown="1">

### `ByID`

<p><code>map[int]map[string]any</code></p>

ByID indexes static-data entries by their numeric _id value.

</div>

<div class="api-field" id="staticdata-byname" markdown="1">

### `ByName`

<p><code>map[string]map[string]any</code></p>

ByName indexes static-data entries by normalized name, section, and
village.

</div>

<div class="api-field" id="staticdata-translations" markdown="1">

### `Translations`

<p><code>map[string]map[string]string</code></p>

Translations maps translation IDs to language-code/value maps.

</div>

<a id="translation"></a>

## Translation

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Translation</code></p>

Translation contains one static-data translation entry.

<div class="api-field" id="translation-id" markdown="1">

### `ID`

<p><code>string</code> <span class="api-json">json: id</span></p>

ID is the translation identifier.

</div>

<div class="api-field" id="translation-english" markdown="1">

### `English`

<p><code>string</code> <span class="api-json">json: EN</span></p>

English is the EN translation value.

</div>

<div class="api-field" id="translation-languages" markdown="1">

### `Languages`

<p><code>map[string]string</code></p>

Languages maps language codes to translated strings.

</div>

## Functions

<a id="loadstaticdata"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.LoadStaticData()<span class="api-return-arrow"> -> </span>(<span class="api-return">*StaticData</span>, <span class="api-return">error</span>)</code></p>

LoadStaticData parses the embedded static-data files once and returns the
shared indexed result.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>*<a href="#staticdata">StaticData</a></code> <code>error</code> </dd>
</dl>

</div>


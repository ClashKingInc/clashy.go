# Static Data

Embedded ClashKing static data, translations, and lookup helpers.

<a id="staticdata"></a>

## Static Data

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.StaticData</code></p>

StaticData is the parsed and indexed ClashKing static data embedded in the
package.

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

## Static Data Methods

<a id="staticdata-lookupbyid"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.StaticData.LookupByID(<span class="api-param">id: int</span>)<span class="api-return-arrow"> -> </span><span class="api-return">map[string]any</span></code></p>

LookupByID returns a static-data entry by numeric static ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>id</strong> (<code>int</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>map[string]any</code> </dd>
</dl>

</div>

<a id="staticdata-lookupbyname"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.StaticData.LookupByName(<span class="api-param">name: string</span>, <span class="api-param">section: string</span>, <span class="api-param">village: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">map[string]any</span></code></p>

LookupByName returns a static-data entry by display name, section, and
village.

The lookup is case-insensitive. The section should match a top-level static
data section such as "troops", "spells", "heroes", "pets", or "equipment".

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
<p><strong>section</strong> (<code>string</code>)</p>
<p><strong>village</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>map[string]any</code> </dd>
</dl>

</div>

<a id="staticdata-section"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.StaticData.Section(<span class="api-param">name: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">[]map[string]any</span></code></p>

Section returns an isolated copy of one top-level static-data section.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>name</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>[]map[string]any</code> </dd>
</dl>

</div>

<a id="staticdata-sections"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.StaticData.Sections()<span class="api-return-arrow"> -> </span><span class="api-return">map[string][]map[string]any</span></code></p>

Sections returns an isolated copy of all top-level static-data sections.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>map[string][]map[string]any</code> </dd>
</dl>

</div>

<a id="staticdata-translation"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.StaticData.Translation(<span class="api-param">id: string</span>)<span class="api-return-arrow"> -> </span><span class="api-return">map[string]string</span></code></p>

Translation returns an isolated language map for one translation ID.

<dl class="api-parameters">
<dt>Parameters:</dt><dd>
<p><strong>id</strong> (<code>string</code>)</p>
</dd>
</dl>

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>map[string]string</code> </dd>
</dl>

</div>

<a id="staticdata-translations"></a>

<div class="api-function" markdown="1">

<p class="api-signature api-function-signature"><code>clashy.StaticData.Translations()<span class="api-return-arrow"> -> </span><span class="api-return">map[string]map[string]string</span></code></p>

Translations returns an isolated copy of all translations.

<dl class="api-parameters">
<dt>Return type:</dt><dd>
<code>map[string]map[string]string</code> </dd>
</dl>

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


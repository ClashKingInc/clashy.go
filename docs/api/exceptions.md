# Exceptions

<a id="clashofclansexception"></a>

## Clash Of Clans Exception

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.ClashOfClansException</code></p>

ClashOfClansException is the base package error type.

<div class="api-field" id="clashofclansexception-message" markdown="1">

### `Message`

<p><code>string</code></p>

Message is the human-readable error message.

</div>

<a id="httpexception"></a>

## HTTPException

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.HTTPException</code></p>

HTTPException captures a non-successful API response.

The typed HTTP errors in this package embed HTTPException so callers can
match either the specific type or inspect shared status, reason, message, and
body fields.

<div class="api-field" id="httpexception-status" markdown="1">

### `Status`

<p><code>int</code></p>

Status is the HTTP status code returned by the API. It can be zero for
transport failures mapped to GatewayError.

</div>

<div class="api-field" id="httpexception-reason" markdown="1">

### `Reason`

<p><code>string</code></p>

Reason is the API reason string when one was provided.

</div>

<div class="api-field" id="httpexception-message" markdown="1">

### `Message`

<p><code>string</code></p>

Message is the API message string when one was provided.

</div>

<div class="api-field" id="httpexception-body" markdown="1">

### `Body`

<p><code>[]byte</code></p>

Body is the raw response body retained for debugging.

</div>

<a id="invalidargument"></a>

## Invalid Argument

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.InvalidArgument</code></p>

InvalidArgument represents a 400 response from the API.

<div class="api-field" id="invalidargument-httpexception" markdown="1">

### `HTTPException`

<p><code>*<a href="#httpexception">HTTPException</a></code></p>

</div>

<a id="invalidcredentials"></a>

## Invalid Credentials

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.InvalidCredentials</code></p>

InvalidCredentials represents a developer-site authentication failure.

<div class="api-field" id="invalidcredentials-httpexception" markdown="1">

### `HTTPException`

<p><code>*<a href="#httpexception">HTTPException</a></code></p>

</div>

<a id="forbidden"></a>

## Forbidden

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Forbidden</code></p>

Forbidden represents a 403 response from the API.

<div class="api-field" id="forbidden-httpexception" markdown="1">

### `HTTPException`

<p><code>*<a href="#httpexception">HTTPException</a></code></p>

</div>

<a id="privatewarlog"></a>

## Private War Log

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.PrivateWarLog</code></p>

PrivateWarLog represents the private-war-log 403 response.

<div class="api-field" id="privatewarlog-httpexception" markdown="1">

### `HTTPException`

<p><code>*<a href="#httpexception">HTTPException</a></code></p>

</div>

<a id="notfound"></a>

## Not Found

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.NotFound</code></p>

NotFound represents a 404 response from the API.

<div class="api-field" id="notfound-httpexception" markdown="1">

### `HTTPException`

<p><code>*<a href="#httpexception">HTTPException</a></code></p>

</div>

<a id="maintenance"></a>

## Maintenance

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.Maintenance</code></p>

Maintenance represents a 503 maintenance response from the API.

<div class="api-field" id="maintenance-httpexception" markdown="1">

### `HTTPException`

<p><code>*<a href="#httpexception">HTTPException</a></code></p>

</div>

<a id="gatewayerror"></a>

## Gateway Error

<p class="api-signature"><span class="api-kind">struct</span> <code>clashy.GatewayError</code></p>

GatewayError represents transport failures and 5xx gateway responses.

<div class="api-field" id="gatewayerror-httpexception" markdown="1">

### `HTTPException`

<p><code>*<a href="#httpexception">HTTPException</a></code></p>

</div>


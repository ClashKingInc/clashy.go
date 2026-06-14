# Authentication

`clashy.go` supports two authentication flows.

## Direct Tokens

Use `LoginWithTokens` when you already have one or more Clash of Clans API tokens:

```go
err := client.LoginWithTokens(ctx, "token-a", "token-b")
```

The HTTP client rotates tokens per request. This is the simplest option for services that manage API keys outside the process.

## Developer-Site Login

Use `Login` when the process should create or reuse developer-site keys:

```go
err := client.Login(ctx, "email@example.com", "password")
```

The client logs into the developer site, discovers the current IP address from the temporary token, reuses matching keys named by `ClientConfig.KeyNames`, and creates additional keys until `ClientConfig.KeyCount` is satisfied.

## Configuration

Start from `DefaultClientConfig` and override only what your service needs:

```go
cfg := clashy.DefaultClientConfig()
cfg.BaseURL = "https://proxy.example.com/v1"
cfg.ThrottleLimit = 60
cfg.Realtime = true

client, err := clashy.NewClient(cfg)
```

`BaseURL` points at the Clash API or a compatible proxy. `DeveloperBaseURL` is only used by developer-site login. `LookupCache` and `UpdateCache` control whether successful GET responses are reused according to the response `Cache-Control` header.

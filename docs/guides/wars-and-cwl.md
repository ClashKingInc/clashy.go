# Wars And CWL

War helpers are designed around the difference between normal wars and Clan War League.

## Current War

`GetClanWar` calls the Clash API current-war endpoint directly. It is best when you only want the raw regular-war response.

`GetCurrentWar` is broader. It first checks the regular current-war endpoint and returns that war when the clan is in a normal war. If the clan is not in a normal war, or the war log is private, it checks the CWL league group and finds the relevant league war for the clan.

```go
war, err := client.GetCurrentWar(ctx, "#CLAN")
```

When the clan is not in any regular war or CWL round, `GetCurrentWar` returns `nil, nil`.

## CWL Round Selection

Pass a `WarRound` to select the logical CWL round:

```go
current, err := client.GetCurrentWar(ctx, "#CLAN", clashy.CurrentWar)
prep, err := client.GetCurrentWar(ctx, "#CLAN", clashy.CurrentPreparation)
previous, err := client.GetLeagueWar(ctx, "#CLAN", clashy.PreviousWar)
```

CWL group responses contain all matchup tags for each round and future rounds can appear as `#0`. The client filters placeholder tags, probes the latest real round when needed, and orients the returned `ClanWar` so `war.Clan` is the requested clan even when the API returned it on the opponent side.

## Batch Helpers

`GetClanWars`, `GetLeagueWars`, and `GetCurrentWars` iterate over tags or war tags and return the successful war models in order. They are intentionally simple; services that need concurrency, retries, or partial-failure reporting should compose the single-war methods with their own worker pool.

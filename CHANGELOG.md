# Changelog

All notable changes to this module are documented here.

## Unreleased

- Updated battle log, ranked player, and CWL/league-group models for the latest Clash of Clans API Swagger fields.
- Added typed battle modifier values for wars and war log entries.
- Fixed developer-site login by reusing the session cookie for API-key management requests.
- Changed `GetPlayerLeagueGroup` to accept string season IDs so full-date season identifiers can be used.
- Pulled the latest embedded static data and regenerated static constants.

## v0.1.4

- Added a daily static asset update workflow that regenerates static data, commits changed assets, increments the patch tag, and publishes a GitHub release.
- Added a pull request test workflow that runs `go test ./...` on pushes and PRs to `main`.
- Hardened current-war lookup so `GetCurrentWar` falls back from normal wars to CWL, handles private war logs and not-found/gateway league-group responses, and returns nil when no current war exists.
- Added CWL round selection helpers that ignore future `#0` placeholders, distinguish current war from current preparation, and orient league-war responses so `Clan` is always the requested clan.
- Added realtime support to `GetClanWar` without applying realtime query parameters to CWL group lookups.
- Replaced tournament-window helpers with season, legend-day, Clan Games, and raid-weekend helpers backed by tests.
- Changed `Clan.WarLeague` to a non-pointer `League` value to match the API's always-present clan league shape.
- Added Read the Docs documentation using the Read the Docs theme, generated sectioned API reference pages, expanded Go doc comments, and runnable examples.

## v0.1.3

- Enforced `CacheMaxSize` so cached GET responses cannot grow past the configured entry limit.
- Added fixed response body limits for API, developer-site, and static-data downloads.
- Escaped `seasonID` in `GetSeasonRankings`.
- Made `ClanWar.Attacks` safe for partial war payloads.
- Removed `ClanWarMember` back-references to avoid cyclic war models.

## v0.1.2

- Renamed clan member league data to `ClanMember.LeagueTier` and mapped it to the `leagueTier` API field.
- Kept `ClanMember.LeagueTier` and `Player.LeagueTier` as non-pointer values because they are always returned by the API.

## v0.1.1

- Added `ClanMember.TownHall` for the `townHallLevel` field returned by clan member payloads.

## v0.1.0

- Initial tagged release.

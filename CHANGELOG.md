# Changelog

All notable changes to this module are documented here.

## v0.1.13

- Updated battle log, ranked player, CWL, and league-group models for the latest Clash of Clans API fields, including typed battle types and battle modifiers.
- Replaced positional list pagination arguments with `PageOptions` and changed `GetPlayerLeagueGroup` to accept string season IDs.
- Changed `HTTPClient.Do` to return `HTTPResponse`, added transport connection controls, and hardened cache expiry, eviction, response-size, and CIDR handling.
- Added an exported rolling-window and in-flight `Limiter`, and made `Client.Close` release idle HTTP connections.
- Fixed developer-site login by reusing the login session cookie for API-key management requests.
- Separated player units from `StaticUnit` lookup results, made static-data access lazy and mutation-safe, and hardened army-link parsing.
- Added timestamp JSON marshaling, explicit war-member attack resolution, and CWL master-roster decoding.
- Updated the static-data generator, pulled the latest embedded data and translations from `main`, and regenerated constants and API documentation.

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

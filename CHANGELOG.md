# Changelog

All notable changes to this module are documented here.

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

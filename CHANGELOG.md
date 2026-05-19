# Changelog

All notable changes to this module are documented here.

## v0.1.2

- Renamed clan member league data to `ClanMember.LeagueTier` and mapped it to the `leagueTier` API field.
- Kept `ClanMember.LeagueTier` and `Player.LeagueTier` as non-pointer values because they are always returned by the API.

## v0.1.1

- Added `ClanMember.TownHall` for the `townHallLevel` field returned by clan member payloads.

## v0.1.0

- Initial tagged release.

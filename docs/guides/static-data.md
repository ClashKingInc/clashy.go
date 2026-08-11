# Static Data

The package embeds ClashKing static data and translations so callers can enrich API responses without another network request.

```go
archer := client.GetTroop("Archer", true, 12)
rage := client.GetSpell("Rage Spell", 6)
queen := client.GetHero("Archer Queen", 95)
```

Each helper returns `nil` when the named object is not found. Passing level `0` uses a sensible default where the helper supports one.

## Army Links

`ParseArmyLink` accepts either a full in-game army link or the raw `army=` payload:

```go
recipe := client.ParseArmyLink("https://link.clashofclans.com/?action=CopyArmy&army=u10x0-2x5s3x2")
```

The returned `ArmyRecipe` separates heroes, pets, hero equipment, home troops, spells, and Clan Castle donations.

## Updating Embedded Data

Run `go generate` from the repository root to download and validate the latest static data and translations, atomically replace the embedded files under `static/`, and regenerate `constants.go`.

Static data is parsed lazily on first use. Runtime clients cannot rewrite the embedded data, so an application sees one immutable snapshot for its lifetime.

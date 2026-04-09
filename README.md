# clashy.go

Easy-to-use Go SDK for the Clash of Clans API.

## Key Features

- Go-native client built around `context.Context`
- Broad coverage of the official Clash of Clans API
- Developer-site email/password login, or direct token login
- Built-in rate limiting, caching, and static game-data helpers
- Optional polling-based events package with pluggable snapshot stores

## Getting Started

### Installing

**Go 1.25 or higher is required**

```sh
go get github.com/clashkinginc/clashy.go
```

To use the latest development version:

```sh
go get github.com/clashkinginc/clashy.go@main
```

## Quick Example

This is the basic usage of the library. This example logs in, fetches a player, searches for clans, and loads the current war for a clan.

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	clashy "github.com/clashkinginc/clashy.go"
)

func main() {
	ctx := context.Background()

	client, err := clashy.NewClient(clashy.DefaultClientConfig())
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Login(ctx, "email", "password"); err != nil {
		var invalid *clashy.InvalidCredentials
		if errors.As(err, &invalid) {
			log.Fatal("invalid developer credentials")
		}
		log.Fatal(err)
	}

	player, err := client.GetPlayer(ctx, "#TAG")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s has %d trophies\n", player.Name, player.Trophies)

	clans, err := client.SearchClans(ctx, clashy.SearchClansRequest{
		Name:  "best clan ever",
		Limit: 5,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, clan := range clans {
		fmt.Printf("%s (%s) has %d members\n", clan.Name, clan.Tag, clan.MemberCount)
	}

	war, err := client.GetCurrentWar(ctx, "#CLANTAG")
	if err != nil {
		var privateWarLog *clashy.PrivateWarLog
		if errors.As(err, &privateWarLog) {
			fmt.Println("uh oh, they have a private war log")
			return
		}
		log.Fatal(err)
	}

	fmt.Printf("%s is currently in %s state\n", war.ClanTag, war.State)
}
```

If you already have API tokens, you can skip developer-site login:

```go
if err := client.LoginWithTokens(ctx, "your-api-token"); err != nil {
	log.Fatal(err)
}
```

## Basic Events Example

The `events` package provides polling-based trackers for clans, players, and wars. Trackers compare snapshots and dispatch handlers when a spec matches.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clashy "github.com/clashkinginc/clashy.go"
	"github.com/clashkinginc/clashy.go/events"
	"github.com/clashkinginc/clashy.go/stores"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := clashy.NewClient(clashy.DefaultClientConfig())
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Login(ctx, "email", "password"); err != nil {
		log.Fatal(err)
	}

	engine := events.NewEngine(client, stores.NewMemoryStore())

	engine.Clans().
		Group(
			"home-clans",
			events.GroupTags("#TAG1", "#TAG2", "#TAG3"),
			events.Interval(30*time.Second),
		).
		On(events.ClanMemberJoined(func(_ context.Context, change events.Change[clashy.Clan]) error {
			fmt.Printf("member joined %s (%s)\n", change.Current.Name, change.Current.Tag)
			return nil
		}))

	engine.Wars().
		Group(
			"war-tracker",
			events.GroupTags("#TAG1"),
			events.Interval(30*time.Second),
		).
		On(events.WarStateChanged(func(_ context.Context, change events.FieldChange[clashy.ClanWar, clashy.WarState]) error {
			fmt.Printf("war state changed from %s to %s for %s\n", change.OldValue, change.NewValue, change.Tag)
			return nil
		}))

	if err := engine.Start(ctx); err != nil {
		log.Fatal(err)
	}

	select {}
}
```

Available snapshot stores include:

- `stores.NewMemoryStore()`
- `stores.NewRedisStore(...)`
- `stores.NewMongoStore(...)`

## Other Features

The client also includes helpers for:

- locations, rankings, leagues, labels, and gold pass endpoints
- battle logs, raid logs, war logs, and CWL data
- player token verification
- static game-data lookups such as troops, spells, heroes, pets, equipment, and translations
- generic decode helpers like `GetClanAs[T]`, `GetPlayerAs[T]`, and `SearchClansAs[T]`

## Contributing

Contributions are welcome. If you find a bug or want to add functionality, open an issue or submit a pull request.

To run the test suite:

```sh
go test ./...
```

## Links

- [Official Clash of Clans API](https://developer.clashofclans.com/)
- [clashy.go Module](https://github.com/clashkinginc/clashy.go)

## Disclaimer

This content is not affiliated with, endorsed, sponsored, or specifically approved by Supercell, and Supercell is not responsible for it. For more information, see [Supercell's Fan Content Policy](https://www.supercell.com/fan-content-policy).

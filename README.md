# clashy.go

Easy-to-use Go SDK for the Clash of Clans API.

## Key Features

- Go-native client built around `context.Context`
- Broad coverage of the official Clash of Clans API
- Developer-site email/password login, or direct token login
- Built-in rate limiting, caching, and static game-data helpers

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

```text
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
	defer func() { _ = client.Close() }()

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

```text
package main

import (
	"context"
	"log"

	clashy "github.com/clashkinginc/clashy.go"
)

func main() {
	ctx := context.Background()

	client, err := clashy.NewClient(clashy.DefaultClientConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = client.Close() }()

	if err := client.LoginWithTokens(ctx, "your-api-token"); err != nil {
		log.Fatal(err)
	}
}
```

## Other Features

The client also includes helpers for:

- locations, rankings, leagues, labels, and gold pass endpoints
- battle logs, raid logs, war logs, and CWL data
- player token verification
- static game-data lookups such as troops, spells, heroes, pets, equipment, and translations

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

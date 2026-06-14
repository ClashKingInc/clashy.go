# Getting Started

Install the module with Go 1.25 or newer:

```sh
go get github.com/clashkinginc/clashy.go
```

Create a client from `DefaultClientConfig`, authenticate, and pass a `context.Context` into every network call:

```go
package main

import (
	"context"
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

	if err := client.LoginWithTokens(ctx, "api-token"); err != nil {
		log.Fatal(err)
	}

	player, err := client.GetPlayer(ctx, "#PLAYER")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s has %d trophies\n", player.Name, player.Trophies)
}
```

Tags can be passed with or without `#`. When `CorrectTags` is enabled in the client config, helper methods normalize tags before building API paths.

## Request Shape

List endpoints return slices, while single-resource endpoints return pointers to the response model. Pagination arguments follow the Clash API convention:

- `limit` controls page size when the endpoint supports it.
- `after` and `before` are cursor strings returned by the API.
- empty values are omitted from the query string.

## Errors

HTTP failures are returned as typed errors such as `InvalidArgument`, `Forbidden`, `PrivateWarLog`, `NotFound`, `Maintenance`, and `GatewayError`. Use `errors.As` when a workflow needs special handling:

```go
var privateWarLog *clashy.PrivateWarLog
if errors.As(err, &privateWarLog) {
	// The clan exists, but its war log is private.
}
```

package clashy_test

import (
	"context"
	"fmt"
	"log"
	"time"

	clashy "github.com/clashkinginc/clashy.go"
)

func ExampleNewClient() {
	ctx := context.Background()

	client, err := clashy.NewClient(clashy.DefaultClientConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = client.Close() }()

	if err := client.LoginWithTokens(ctx, "api-token"); err != nil {
		log.Fatal(err)
	}
}

func ExampleCorrectTag() {
	fmt.Println(clashy.CorrectTag(" o0q l "))

	// Output:
	// #00QL
}

func ExampleGetSeasonByID() {
	season, err := clashy.GetSeasonByID("2025-09")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(season.SeasonID)
	fmt.Println(season.StartTime.Format(time.RFC3339))
	fmt.Println(season.EndTime.Format(time.RFC3339))

	// Output:
	// 2025-09
	// 2025-08-25T05:00:00Z
	// 2025-10-06T05:00:00Z
}

func ExampleParseArmyRecipe() {
	staticData, err := clashy.LoadStaticData()
	if err != nil {
		log.Fatal(err)
	}

	recipe := clashy.ParseArmyRecipe(staticData, "u10x0-2x5s3x2")
	fmt.Println(len(recipe.Troops))
	fmt.Println(len(recipe.Spells))

	// Output:
	// 2
	// 1
}

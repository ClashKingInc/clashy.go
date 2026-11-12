// Verify the real SDK HTTP service against the separately running Mock API.
package main

import (
	"context"
	"errors"
	"fmt"
	clashy "github.com/clashkinginc/clashy.go"
	"os"
	"time"
)

func main() {
	base := os.Getenv("CLASH_CONTRACT_TEST_ORIGIN")
	if len(base) < 17 || base[:17] != "http://127.0.0.1:" {
		panic("owned loopback Mock API required")
	}
	client := clashy.NewHTTPClient(clashy.ClientConfig{Timeout: 5 * time.Second, ThrottleLimit: 100})
	client.SetTokens("disposable")
	for _, c := range []struct {
		s           string
		maintenance bool
	}{{"maintenance", true}, {"unavailable", false}} {
		_, err := client.Do(context.Background(), "GET", base+"/players/%232PP?scenario="+c.s, nil, clashy.RequestOptions{})
		var maintenance *clashy.Maintenance
		var gateway *clashy.GatewayError
		if c.maintenance && !errors.As(err, &maintenance) || !c.maintenance && !errors.As(err, &gateway) {
			panic(fmt.Sprintf("%s: %T", c.s, err))
		}
	}
	for _, test := range []struct {
		scenario string
		private  bool
	}{{"privateWarLog", true}, {"forbidden", false}, {"invalidIp", false}} {
		_, err := client.Do(context.Background(), "GET", base+"/clans/%232PP/warlog?scenario="+test.scenario, nil, clashy.RequestOptions{})
		var private *clashy.PrivateWarLog
		var forbidden *clashy.Forbidden
		if test.private && !errors.As(err, &private) || !test.private && !errors.As(err, &forbidden) {
			panic(fmt.Sprintf("%s: %T", test.scenario, err))
		}
	}
	cfg := clashy.DefaultClientConfig()
	cfg.BaseURL = base
	cfg.LookupCache = false
	cfg.UpdateCache = false
	full, err := clashy.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	if err = full.LoginWithTokens(context.Background(), "disposable"); err != nil {
		panic(err)
	}
	clan, err := full.GetClan(context.Background(), "#2PP")
	if err != nil || clan.Name == "" {
		panic(fmt.Sprintf("real clan decode: %v", err))
	}
	found, err := full.SearchClans(context.Background(), clashy.SearchClansRequest{Name: "test", Limit: 3})
	if err != nil || len(found) == 0 {
		panic(fmt.Sprintf("real search decode: %v", err))
	}
	fmt.Println("PASS real Go SDK: maintenance/outage and private-war/credential failures remain distinct")
}

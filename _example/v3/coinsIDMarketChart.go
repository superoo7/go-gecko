package main

import (
	"fmt"
	"log"
	"time"

	gecko "github.com/ggarcia14/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	m, err := cg.CoinsIDMarketChart("bitcoin", "usd", "1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Prices\n")
	for _, v := range *m.Prices {
		fmt.Printf("%s:%.04f\n", time.Unix(int64(v[0])/1000, int64(v[0])%1000).UTC().Format(time.RFC3339), v[1])
	}

	fmt.Printf("MarketCaps\n")
	for _, v := range *m.MarketCaps {
		fmt.Printf("%s:%.04f\n", time.Unix(int64(v[0])/1000, int64(v[0])%1000).UTC().Format(time.RFC3339), v[1])
	}

	fmt.Printf("TotalVolumes\n")
	for _, v := range *m.TotalVolumes {
		fmt.Printf("%s:%.04f\n", time.Unix(int64(v[0])/1000, int64(v[0])%1000).UTC().Format(time.RFC3339), v[1])
	}
}

package main

import (
	"fmt"
	"log"

	gecko "github.com/ggarcia14/go-gecko/v3"
	geckoTypes "github.com/ggarcia14/go-gecko/v3/types"
)

func main() {
	cg := gecko.NewClient(nil)
	// find specific coins
	vsCurrency := "usd"
	ids := []string{"bitcoin", "ethereum", "steem"}
	perPage := 1
	page := 1
	sparkline := true
	pcp := geckoTypes.PriceChangePercentageObject
	priceChangePercentage := []string{pcp.PCP1h, pcp.PCP24h, pcp.PCP7d, pcp.PCP14d, pcp.PCP30d, pcp.PCP200d, pcp.PCP1y}
	order := geckoTypes.OrderTypeObject.MarketCapDesc
	market, err := cg.CoinsMarket(vsCurrency, ids, order, perPage, page, sparkline, priceChangePercentage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total coins: ", len(*market))
	fmt.Println(*market)

	// with pagination instead
	ids = []string{}
	perPage = 10
	page = 1
	market, err = cg.CoinsMarket(vsCurrency, ids, order, perPage, page, sparkline, priceChangePercentage)
	fmt.Println("Total coins: ", len(*market))
	fmt.Println(*market)
}

package main

import (
	"fmt"
	"log"
	"time"

	gecko "github.com/ggarcia14/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	global, err := cg.Global()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("active crypto:", global.ActiveCryptocurrencies)
	fmt.Println("upcoming ico:", global.UpcomingICOs)
	fmt.Println("ended ico:", global.EndedICOs)
	fmt.Println("market:", global.Markets)
	fmt.Println("BTC Volume:", global.TotalVolume["btc"])
	fmt.Println("Global Total Market Cap in USD:", global.TotalMarketCap["usd"])
	fmt.Println("Market Cap Percentage of ETH:", global.MarketCapPercentage["eth"])
	fmt.Println("Market Cap Change Percentage 24h USD:", global.MarketCapChangePercentage24hUSD)
	fmt.Println("last updated:", time.Unix(global.UpdatedAt, 0))
}

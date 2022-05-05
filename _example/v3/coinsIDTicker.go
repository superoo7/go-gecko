package main

import (
	"fmt"
	"log"

	gecko "github.com/ggarcia14/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	tickers, err := cg.CoinsIDTickers("bitcoin", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(tickers.Tickers))
	tickers, err = cg.CoinsIDTickers("bitcoin", 1)
	fmt.Println(len(tickers.Tickers))
}

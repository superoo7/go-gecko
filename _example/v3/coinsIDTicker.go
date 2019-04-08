package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	tickers, err := gecko.CoinsIDTickers("bitcoin", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(tickers.Tickers))
	tickers, err = gecko.CoinsIDTickers("bitcoin", 1)
	fmt.Println(len(tickers.Tickers))
}

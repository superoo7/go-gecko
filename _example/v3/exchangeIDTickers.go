package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)

	coinIDs := []string{"ripple"}
	resp, err := cg.ExchangeIDTickers("korbit", coinIDs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total tickers: ", len(*&resp.Tickers))
	for _, ticker := range *&resp.Tickers {
		fmt.Println(ticker)
	}
}

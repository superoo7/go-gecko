package main

import (
	"fmt"
	"log"

	gecko "github.com/ggarcia14/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)

	price, err := cg.SimpleSinglePrice("bitcoin", "usd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("%s is worth %f %s", price.ID, price.MarketPrice, price.Currency))
}

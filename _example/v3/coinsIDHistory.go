package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	btc, err := gecko.CoinsIDHistory("bitcoin", "30-12-2018", true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*btc)
}

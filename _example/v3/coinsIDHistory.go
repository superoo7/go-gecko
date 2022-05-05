package main

import (
	"fmt"
	"log"

	gecko "github.com/ggarcia14/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	btc, err := cg.CoinsIDHistory("bitcoin", "30-12-2018", true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*btc)
}

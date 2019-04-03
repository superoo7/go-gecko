package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/src/v3"
)

func main() {
	currencies, err := gecko.SimpleSupportedVSCurrencies()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total currencies", len(*currencies))
	fmt.Println(*currencies)
}

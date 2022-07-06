package main

import (
	"fmt"
	"log"

	gecko "github.com/relative-finance/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	coin, err := cg.CoinsID("dogecoin", true, true, true, true, true, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coin)
}

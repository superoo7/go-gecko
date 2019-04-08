package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	rate, err := gecko.ExchangeRates()
	if err != nil {
		log.Fatal(err)
	}
	r := (*rate)["btc"]
	fmt.Println(r.Name)
	fmt.Println(r.Unit)
	fmt.Println(r.Value)
	fmt.Println(r.Type)
}

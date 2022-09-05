package main

import (
	"fmt"
	"log"

	gecko "github.com/tomastoth/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	rate, err := cg.ExchangeRates()
	if err != nil {
		log.Fatal(err)
	}
	r := (*rate)["btc"]
	fmt.Println(r.Name)
	fmt.Println(r.Unit)
	fmt.Println(r.Value)
	fmt.Println(r.Type)
}

package main

import (
	"fmt"
	gecko "github.com/superoo7/go-gecko/v3"
	"github.com/superoo7/go-gecko/v3/types"
	"log"
)

func main() {
	cg := gecko.NewClient(nil)
	address := "0xdac17f958d2ee523a2206206994597c13d831ec7"
	prices, err := cg.SimplePriceByAddress([]string{address}, []string{"usd"}, types.Ethereum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println((*prices)[address]["usd"])
}

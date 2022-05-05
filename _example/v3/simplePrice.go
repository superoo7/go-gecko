package main

import (
	"fmt"
	"log"

	gecko "github.com/ggarcia14/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)

	ids := []string{"bitcoin", "ethereum"}
	vc := []string{"usd", "myr"}
	sp, err := cg.SimplePrice(ids, vc)
	if err != nil {
		log.Fatal(err)
	}
	bitcoin := (*sp)["bitcoin"]
	eth := (*sp)["ethereum"]
	fmt.Println(fmt.Sprintf("Bitcoin is worth %f usd (myr %f)", bitcoin["usd"], bitcoin["myr"]))
	fmt.Println(fmt.Sprintf("Ethereum is worth %f usd (myr %f)", eth["usd"], eth["myr"]))
}

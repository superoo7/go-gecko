package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	price, err := gecko.SimpleSinglePrice("bitcoin", "usd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(price)
}

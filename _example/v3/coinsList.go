package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/src/v3"
)

func main() {
	list, err := gecko.CoinsList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available coins:", len(*list))
}

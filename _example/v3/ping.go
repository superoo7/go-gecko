package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/src/v3"
)

func main() {
	ping, err := gecko.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ping.GeckoSays)
}

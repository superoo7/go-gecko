package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	ping, err := gecko.Ping()
	if err != null {
		log.Fatal(err)
	}
	fmt.Println(ping.GeckoSays)
}

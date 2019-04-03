package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/src/v3"
)

func main() {
	t, err := gecko.EventsTypes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.Count)
	fmt.Println(t.Data)
}

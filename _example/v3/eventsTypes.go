package main

import (
	"fmt"
	"log"

	gecko "github.com/ggarcia14/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	t, err := cg.EventsTypes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.Count)
	fmt.Println(t.Data)
}

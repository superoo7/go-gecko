package main

import (
	"fmt"
	"log"

	gecko "github.com/ggarcia14/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	e, err := cg.EventsCountries()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(e))
	fmt.Println(e[0])
	fmt.Println(e[1])
	fmt.Println(e[2].Code)
	fmt.Println(e[2].Country)
}

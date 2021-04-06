package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)

	exchanges, err := cg.ExchangesList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total exchanges: ", len(*exchanges))
	for _, exchange := range *exchanges {
		fmt.Println(exchange)
	}
}

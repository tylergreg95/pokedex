package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous
	return nil
}

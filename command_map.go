package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Println()
	cfg.nextLocationAreaURL = locationAreas.Next
	cfg.previousLocationAreaURL = locationAreas.Previous
	return nil

}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("You're on the first page\n")
	}

	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Println()
	cfg.nextLocationAreaURL = locationAreas.Next
	cfg.previousLocationAreaURL = locationAreas.Previous
	return nil

}

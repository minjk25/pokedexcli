package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for key, command := range getCommand() {
		fmt.Printf(" - %s: %s\n", key, command.description)
	}
	fmt.Println()
	return nil
}

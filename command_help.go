package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range getCommand() {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

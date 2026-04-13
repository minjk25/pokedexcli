package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("No pokemon name provided\n")
	}
	if len(args) > 1 {
		return errors.New("One input only! You can inspect one pokemon at a time\n")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("You haven't caught this pokemon ('%s') yet!\n", pokemonName)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, pokeStat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", pokeStat.Stat.Name, pokeStat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokeType.Type.Name)
	}

	fmt.Println()
	return nil

}

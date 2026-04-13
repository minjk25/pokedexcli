package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("No pokemon name provided\n")
	}
	if len(args) > 1 {
		return errors.New("One input only! You can catch one pokemon at a time\n")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("%s escaped!\n", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!", pokemonName)
	fmt.Println()
	return nil

}

package main

import (
	"errors"
	"fmt"
	"log/slog"
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
	if pokemon, ok := cfg.caughtPokemon[pokemonName]; ok {
		return fmt.Errorf("You already caught '%s'!\n", pokemon.Name)
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)

	// for debugging
	slog.Debug("catch attempt",
		"baseExperience", pokemon.BaseExperience,
		"randNum", randNum,
		"threshold", threshold,
	)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if randNum > threshold {
		return fmt.Errorf("%s escaped!\n", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println()
	return nil

}

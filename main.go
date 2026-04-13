package main

import (
	"time"

	"github.com/minjk25/pokedexcli/internal/pokeapi"
)

func main() {
	cacheInterval := time.Hour
	timeout := 5 * time.Second
	cfg := config{
		pokeapiClient: pokeapi.NewClient(cacheInterval, timeout),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}

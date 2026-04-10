package main

import (
	"time"

	"github.com/minjk25/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	startRepl(&cfg)
}

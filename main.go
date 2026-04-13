package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/minjk25/pokedexcli/internal/pokeapi"
)

func main() {
	// for debugging
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	cacheInterval := time.Hour
	timeout := 5 * time.Second
	cfg := config{
		pokeapiClient: pokeapi.NewClient(cacheInterval, timeout),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}

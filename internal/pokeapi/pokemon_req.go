package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)

	// for debugging
	slog.Debug("cache lookup", "url", fullURL, "cache hit", ok)

	if ok {
		// cache hit = true, using cache data instead of fetching the new one
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	// if there is no cache data then fetch the new one with new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("Error: bad status code: %v\n", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// save data to the cache before return
	c.cache.Add(fullURL, data)

	return pokemon, nil

}

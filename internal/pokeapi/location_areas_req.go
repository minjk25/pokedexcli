package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	data, ok := c.cache.Get(fullURL)

	// for debugging
	slog.Debug("cache lookup", "url", fullURL, "cache hit", ok)

	if ok {
		// cache hit = true, using cache data instead of fetching the new one
		locationAreas := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreas, nil
	}

	// if there is no cache data then fetch the new one with new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("Error: bad status code: %v\n", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreas := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// save data to the cache before return
	c.cache.Add(fullURL, data)

	return locationAreas, nil

}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)

	// for debugging
	slog.Debug("cache lookup", "url", fullURL, "cache hit", ok)

	if ok {
		// cache hit = true, using cache data instead of fetching the new one
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	// if there is no cache data then fetch the new one with new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("Error: bad status code: %v\n", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	// save data to the cache before return
	c.cache.Add(fullURL, data)

	return locationArea, nil

}

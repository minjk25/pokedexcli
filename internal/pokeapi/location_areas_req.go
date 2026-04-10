package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

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
		return LocationAreasResp{}, fmt.Errorf("Error: bad status code: %v", resp.StatusCode)
	}

	var locationAreas LocationAreasResp
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationAreas); err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreas, nil

}

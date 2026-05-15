package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(pageURL string) (LocationArea, error) {
	url := baseURL + pageURL

	if val, ok := c.cache.Get(url); ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("bad status code(invalid location name): %d", res.StatusCode)
	}

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	if err := json.Unmarshal(jsonData, &locationArea); err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, jsonData)
	return locationArea, nil
}

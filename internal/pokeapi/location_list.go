package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ShallowLocations, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ShallowLocations{}, err
	}
	defer res.Body.Close()

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return ShallowLocations{}, err
	}

	locations := ShallowLocations{}
	if err := json.Unmarshal(jsonData, &locations); err != nil {
		return ShallowLocations{}, nil
	}

	return locations, nil
}

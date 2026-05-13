package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(c *config) error {
	url := ""

	if c.next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error while reading json body: %w", err)
	}

	var locationData LocationData
	if err := json.Unmarshal(data, &locationData); err != nil {
		return fmt.Errorf("error reading from unmarshaling: %w", err)
	}

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

	c.previous = url
	c.next = locationData.Next

	return nil
}

type LocationData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

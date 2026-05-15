package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pageURL string) (PokemonData, error) {
	url := baseURL + pageURL

	if val, ok := c.cache.Get(url); ok {
		pokemonData := PokemonData{}
		err := json.Unmarshal(val, &pokemonData)
		if err != nil {
			return PokemonData{}, err
		}

		return pokemonData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return PokemonData{}, fmt.Errorf("bad status code(invalid pokemon name): %d", res.StatusCode)
	}

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonData{}, err
	}

	pokemonData := PokemonData{}
	if err := json.Unmarshal(jsonData, &pokemonData); err != nil {
		return PokemonData{}, err
	}

	c.cache.Add(url, jsonData)
	return pokemonData, nil
}

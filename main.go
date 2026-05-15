package main

import (
	"ZafirChowdhury/pokedexGo/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.PokemonData),
	}

	startRepl(cfg)
}

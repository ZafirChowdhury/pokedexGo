package main

import (
	"ZafirChowdhury/pokedexGo/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{pokeapiClient: pokeClient}

	startRepl(cfg)
}

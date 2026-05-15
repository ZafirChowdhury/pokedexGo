package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 2 {
		return errors.New("you must provide a location name")
	}

	locationName := args[1]

	if locationName == "" {
		return errors.New("Location cannot be blank")
	}

	endpoint := "/location-area/" + locationName

	locationArea, err := cfg.pokeapiClient.ListPokemons(endpoint)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")

	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

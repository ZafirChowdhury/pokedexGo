package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(c *config, args []string) error {
	if len(args) != 2 {
		return errors.New("Invalid command. Usage: catch <pokemon-name>")
	}

	pokemonName := args[1]

	// check if user allreay has the pokemon
	_, ok := c.caughtPokemon[pokemonName]
	if ok {
		fmt.Printf("You allready caught %s\n", pokemonName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := "/pokemon/" + pokemonName
	pokemonData, err := c.pokeapiClient.GetPokemon(url)
	if err != nil {
		return err
	}

	userCatchRate := rand.IntN(100)

	if userCatchRate < pokemonData.BaseExperience {
		// user fails to catch
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	c.caughtPokemon[pokemonName] = pokemonData
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}

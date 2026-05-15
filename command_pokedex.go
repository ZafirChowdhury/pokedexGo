package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have yet to catch any pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, val := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", val.Name)
	}

	return nil
}

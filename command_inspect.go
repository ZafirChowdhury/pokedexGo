package main

import (
	"errors"
	"fmt"
)

func commendInspect(c *config, args []string) error {
	if len(args) != 2 {
		return errors.New("Invalid command. Usage: inspect <pokemon-name>")
	}

	pokemonName := args[1]

	p, ok := c.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("You dont have that pokemon!")
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	for _, s := range p.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

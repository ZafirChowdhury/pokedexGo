package main

import (
	"ZafirChowdhury/pokedexGo/internal/pokeapi"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	prevURL       *string
	caughtPokemon map[string]pokeapi.PokemonData
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		ok := scanner.Scan()
		if !ok {
			break
		}

		input := scanner.Text()
		cleanInput := cleanInput(input)

		if len(cleanInput) == 0 {
			continue
		}

		userCmd := cleanInput[0]

		command, ok := getCommands()[userCmd]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		if err := command.callback(cfg, cleanInput); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)

	lowered := strings.ToLower(trimmed)

	words := strings.Fields(lowered)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays the program usage",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Displays next 20 location",
			callback:    commandMapf,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location",
			callback:    commandMapb,
		},

		"explore": {
			name:        "explore",
			description: "Displayes pokemon in the given area | usage: explore <location-name>",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon | usage: catch <pokemon-name>",
			callback:    commandCatch,
		},

		"inspect": {
			name:        "inspect",
			description: "Displays stats of one of your pokemon | usage: inspect <pokemon-name>",
			callback:    commendInspect,
		},
	}
}

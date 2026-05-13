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

		if err := command.callback(cfg); err != nil {
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
	callback    func(*config) error
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
	}
}

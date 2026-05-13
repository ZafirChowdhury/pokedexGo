package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}

	for {
		fmt.Print("Pokedex > ")

		ok := scanner.Scan()
		if !ok {
			break
		}

		input := scanner.Text()
		cleanInput := cleanInput(input)
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

type config struct {
	next     string
	previous string
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
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location",
			callback:    commandMapb,
		},
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Pokedex > ")

		ok := scanner.Scan()
		if !ok {
			break
		}

		input := scanner.Text()
		cleanInput := cleanInput(input)

		for _, command := range getCommands() {
			for _, userInput := range cleanInput {
				if command.name == userInput {
					command.callback()
				}
			}
		}

		fmt.Println("Invalid commend!")
	}
}

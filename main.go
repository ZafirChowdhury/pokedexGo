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

		fmt.Printf("Your command was: %s\n", cleanInput[0])
	}
}

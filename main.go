package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commandMap := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Get info on how to use the Pokedex",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		command, exists := commandMap[words[0]]

		if exists {
			err := command.callback()

			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

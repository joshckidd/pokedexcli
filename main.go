package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var currentConfig config
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		command, exists := getCommandMap()[words[0]]

		if exists {
			err := command.callback(&currentConfig)

			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		} else {
			fmt.Println("Unknown command")
		}

	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/joshckidd/pokedexcli/internal/pokecache"
)

func main() {

	var currentConfig config
	currentConfig.cache = pokecache.NewCache(5 * time.Second)
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

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joshckidd/pokedexcli/internal/pokeapi"
	"github.com/joshckidd/pokedexcli/internal/pokecache"
)

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Get the next page in a list of Pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "help",
			description: "Get the previous page in a list of Pokemon locations",
			callback:    commandMapb,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	previousURL string
	nextUrl     string
	cache       pokecache.Cache
}

func cleanInput(text string) []string {
	var cleanString []string
	words := strings.Fields(text)

	for _, word := range words {
		cleanWord := strings.ToLower(word)
		cleanString = append(cleanString, cleanWord)
	}

	return cleanString
}

func commandExit(_ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandMap(currentConfig *config) error {
	locations, err := pokeapi.GetLocations(currentConfig.nextUrl, currentConfig.cache)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	currentConfig.nextUrl = locations.Next
	currentConfig.previousURL = locations.Previous

	return nil
}

func commandMapb(currentConfig *config) error {
	locations, err := pokeapi.GetLocations(currentConfig.previousURL, currentConfig.cache)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	currentConfig.nextUrl = locations.Next
	currentConfig.previousURL = locations.Previous

	return nil
}

package main

import (
	"fmt"
	"math/rand"
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
			name:        "mapb",
			description: "Get the previous page in a list of Pokemon locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get a list of Pokemon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(string, *config) error
}

type config struct {
	previousURL string
	nextUrl     string
	pokedex     map[string]pokeapi.Pokemon
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

func commandExit(_ string, _ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ string, _ *config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandMap(_ string, currentConfig *config) error {
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

func commandMapb(_ string, currentConfig *config) error {
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

func commandExplore(arg string, currentConfig *config) error {
	apiUrl := "https://pokeapi.co/api/v2/location-area/" + arg

	locations, err := pokeapi.GetLocationDetails(apiUrl, currentConfig.cache)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locations.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(arg string, currentConfig *config) error {
	apiUrl := "https://pokeapi.co/api/v2/pokemon/" + arg

	pokemon, err := pokeapi.GetPokemon(apiUrl, currentConfig.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	xp := pokemon.BaseExperience
	r := rand.Intn(500)
	if r > xp {
		fmt.Printf("%s was caught!\n", arg)
		currentConfig.pokedex[arg] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", arg)
	}

	return nil
}

func commandInspect(arg string, currentConfig *config) error {
	pokemon, ok := currentConfig.pokedex[arg]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %v\nWeight: %v\n", arg, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(arg string, currentConfig *config) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range currentConfig.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}

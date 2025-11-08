package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/joshckidd/pokedexcli/internal/pokecache"
)

type pokeapiLocation struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type pokeapiLocationDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocations(url string, cache pokecache.Cache) (pokeapiLocation, error) {
	var apiURL string
	if url == "" {
		apiURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		apiURL = url
	}
	var body []byte
	val, ok := cache.Get(apiURL)
	if ok {
		body = val
	} else {
		res, err := http.Get(apiURL)
		if err != nil {
			return pokeapiLocation{}, err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return pokeapiLocation{}, fmt.Errorf("Response code: %v", res.StatusCode)
		}
		if err != nil {
			return pokeapiLocation{}, err
		}
		cache.Add(apiURL, body)
	}

	var locations pokeapiLocation
	err := json.Unmarshal(body, &locations)
	if err != nil {
		return pokeapiLocation{}, err
	}
	return locations, nil
}

func GetLocationDetails(apiURL string, cache pokecache.Cache) (pokeapiLocationDetails, error) {
	var body []byte
	val, ok := cache.Get(apiURL)
	if ok {
		body = val
	} else {
		res, err := http.Get(apiURL)
		if err != nil {
			return pokeapiLocationDetails{}, err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return pokeapiLocationDetails{}, fmt.Errorf("Response code: %v", res.StatusCode)
		}
		if err != nil {
			return pokeapiLocationDetails{}, err
		}
		cache.Add(apiURL, body)
	}

	var locations pokeapiLocationDetails
	err := json.Unmarshal(body, &locations)
	if err != nil {
		return pokeapiLocationDetails{}, err
	}
	return locations, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func getLocations(url string) (pokeapiLocation, error) {
	var apiURL string
	if url == "" {
		apiURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		apiURL = url
	}

	res, err := http.Get(apiURL)
	if err != nil {
		return pokeapiLocation{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return pokeapiLocation{}, fmt.Errorf("Response code: %v", res.StatusCode)
	}
	if err != nil {
		return pokeapiLocation{}, err
	}

	var locations pokeapiLocation
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return pokeapiLocation{}, err
	}
	return locations, nil
}

package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var cleanString []string
	words := strings.Fields(text)

	for _, word := range words {
		cleanWord := strings.ToLower(word)
		cleanString = append(cleanString, cleanWord)
	}

	return cleanString
}

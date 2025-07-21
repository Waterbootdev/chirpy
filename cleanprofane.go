package main

import (
	"slices"
	"strings"
)

const PROFANE_MASK string = "****"

func currentProfaneWords() []string {
	return []string{"kerfuffle", "sharbert", "fornax"}
}

func cleanProfaneWords(body string, profaneWords []string) string {

	words := strings.Split(body, " ")

	for i, word := range words {
		if slices.Contains(profaneWords, strings.ToLower(word)) {
			words[i] = PROFANE_MASK
		}
	}

	return strings.Join(words, " ")
}

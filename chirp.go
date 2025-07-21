package main

import (
	"slices"
	"strings"
)

type chirp struct {
	Body string `json:"body"`
}

type chirpError struct {
	Error string `json:"error"`
}

type cleanedChirp struct {
	CleanedBody string `json:"cleaned_body"`
}

const MAX_CHIRP_LENGTH int = 140

func ProfaneWords() []string {
	return []string{"kerfuffle", "sharbert", "fornax"}
}

const PROFANE_MASK string = "****"

func getCleanChirp(body string, profaneWords []string) string {

	words := strings.Split(body, " ")

	for i, word := range words {
		if slices.Contains(profaneWords, strings.ToLower(word)) {
			words[i] = PROFANE_MASK
		}
	}

	return strings.Join(words, " ")
}

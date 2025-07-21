package chirp

import (
	"slices"
	"strings"
)

const MAX_CHIRP_LENGTH int = 140

const PROFANE_MASK string = "****"

func (c Chirp) CleanProfaneWords(profaneWords []string) string {

	words := strings.Split(c.Body, " ")

	for i, word := range words {
		if slices.Contains(profaneWords, strings.ToLower(word)) {
			words[i] = PROFANE_MASK
		}
	}

	return strings.Join(words, " ")
}

func (c Chirp) IsToLong() bool {
	return len(c.Body) > MAX_CHIRP_LENGTH
}

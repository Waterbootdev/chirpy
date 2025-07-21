package chirp

import (
	"net/http"
	"slices"
	"strings"
)

const MAX_CHIRP_LENGTH int = 140

const PROFANE_MASK string = "****"

func (c Chirp) cleanProfaneWords(profaneWords []string) string {

	words := strings.Split(c.Body, " ")

	for i, word := range words {
		if slices.Contains(profaneWords, strings.ToLower(word)) {
			words[i] = PROFANE_MASK
		}
	}

	return strings.Join(words, " ")
}

func (c Chirp) isToLong() bool {
	return len(c.Body) > MAX_CHIRP_LENGTH
}

func (c Chirp) cleanedResponse(writer http.ResponseWriter) {
	cleanedResponse(writer, c.cleanProfaneWords(currentProfaneWords()))
}

func (c Chirp) isToLongErrorResponse(writer http.ResponseWriter) bool {
	toLong := c.isToLong()

	if toLong {
		errorResponse(writer, http.StatusBadRequest, "Chirp is too long")
	}

	return toLong
}

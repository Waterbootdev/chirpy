package apiconfig

import (
	"net/http"

	"slices"
	"strings"

	"github.com/google/uuid"

	"github.com/Waterbootdev/chirpy/internal/response"
)

type chirpRequest struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

const MAX_CHIRP_LENGTH int = 140

const PROFANE_MASK string = "****"

func (r *chirpRequest) cleanProfaneWords(profaneWords []string) {

	words := strings.Split(r.Body, " ")

	for i, word := range words {
		if slices.Contains(profaneWords, strings.ToLower(word)) {
			words[i] = PROFANE_MASK
		}
	}

	r.Body = strings.Join(words, " ")
}

func (r *chirpRequest) isToLong() bool {
	return len(r.Body) > MAX_CHIRP_LENGTH
}

func (r *chirpRequest) isToLongErrorResponse(writer http.ResponseWriter) bool {
	toLong := r.isToLong()

	if toLong {
		response.ErrorResponse(writer, http.StatusBadRequest, "Chirp is too long")
	}

	return toLong
}

func currentProfaneWords() []string {
	return []string{"kerfuffle", "sharbert", "fornax"}
}

func (r chirpRequest) IsValidResponse(_ *ApiConfig, writer http.ResponseWriter) (*chirpRequest, bool) {

	valid := !r.isToLongErrorResponse(writer)

	if valid {
		r.cleanProfaneWords(currentProfaneWords())
	}

	return &r, valid
}

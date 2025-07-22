package apiconfig

import (
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) CreateChirp(request *http.Request, chirpRequest *chirpRequest) (*chirp, error) {
	timeNow := time.Now()
	c, err := cfg.queries.CreateChirp(request.Context(), database.CreateChirpParams{
		ID:        uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Body:      chirpRequest.Body,
		UserID:    chirpRequest.UserID,
	})
	return fromDatabaseChirp(&c), err
}

const MAX_CHIRP_LENGTH int = 140

const PROFANE_MASK string = "****"

func (c *chirpRequest) cleanProfaneWords(profaneWords []string) {

	words := strings.Split(c.Body, " ")

	for i, word := range words {
		if slices.Contains(profaneWords, strings.ToLower(word)) {
			words[i] = PROFANE_MASK
		}
	}

	c.Body = strings.Join(words, " ")
}

func (c *chirpRequest) isToLong() bool {
	return len(c.Body) > MAX_CHIRP_LENGTH
}

func (c *chirpRequest) isToLongErrorResponse(writer http.ResponseWriter) bool {
	toLong := c.isToLong()

	if toLong {
		response.ErrorResponse(writer, http.StatusBadRequest, "Chirp is too long")
	}

	return toLong
}
func currentProfaneWords() []string {
	return []string{"kerfuffle", "sharbert", "fornax"}
}

func (c chirpRequest) IsValidResponse(writer http.ResponseWriter) bool {

	valid := !c.isToLongErrorResponse(writer)

	if valid {
		c.cleanProfaneWords(currentProfaneWords())
	}

	return valid
}

package apiconfig

import (
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

type chirpRequest struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
	ID     uuid.UUID `json:"id"`
	At     time.Time `json:"at"`
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
	return response.ErrorResponse(r.isToLong(), writer, http.StatusBadRequest, "Chirp is too long")
}

func (cfg *ApiConfig) chirpRequestValidator(writer http.ResponseWriter, request *http.Request, chirpRequest *chirpRequest) bool {

	if userID, ok := cfg.validateJWTResponse(request, writer); ok {
		chirpRequest.UserID = userID
		return !chirpRequest.isToLongErrorResponse(writer)
	} else {
		return ok
	}
}

func currentProfaneWords() []string {
	return []string{"kerfuffle", "sharbert", "fornax"}
}

func (cfg *ApiConfig) createChirpHandle(request *http.Request, chirpRequest *chirpRequest) (*chirp, error) {
	chirpRequest.cleanProfaneWords(currentProfaneWords())
	chirpRequest.At = time.Now()
	chirpRequest.ID = uuid.New()
	return cfg.createChirp(request, chirpRequest)
}

func (cfg *ApiConfig) CreateChirpHandler(writer http.ResponseWriter, request *http.Request) {
	response.ContentBodyHandler(writer, request, cfg.createChirpHandle, cfg.chirpRequestValidator, http.StatusCreated)
}

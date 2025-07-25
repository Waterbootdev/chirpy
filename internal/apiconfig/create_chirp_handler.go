package apiconfig

import (
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/generic_handler"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
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

func (cfg *ApiConfig) validateJWTResponse(request *http.Request, writer http.ResponseWriter) (uuid.UUID, bool) {
	userID, err := cfg.validateJWT(request)
	ok := err == nil
	if !ok {
		response.ErrorResponse(writer, http.StatusUnauthorized, "Unauthorized")
	}
	return userID, ok
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

func (cfg *ApiConfig) CreateChirpHandler(writer http.ResponseWriter, request *http.Request) {
	generic_handler.HandlerBody(writer, request, cfg.createChirpHandle, cfg.chirpRequestValidator, http.StatusCreated)
}

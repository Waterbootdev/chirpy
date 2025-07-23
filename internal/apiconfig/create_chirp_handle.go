package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

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

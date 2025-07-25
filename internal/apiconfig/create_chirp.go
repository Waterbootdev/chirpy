package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/database"
)

func (cfg *ApiConfig) createChirp(request *http.Request, chirpRequest *chirpRequest) (*chirp, error) {
	c, err := cfg.queries.CreateChirp(request.Context(), database.CreateChirpParams{
		ID:        chirpRequest.ID,
		CreatedAt: chirpRequest.At,
		UpdatedAt: chirpRequest.At,
		Body:      chirpRequest.Body,
		UserID:    chirpRequest.UserID,
	})
	return fromDatabaseChirp(&c), err
}

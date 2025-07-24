package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
)

func (cfg *ApiConfig) webhookHandle(request *http.Request, webhook *webhook) (*error, error) {
	return nil, cfg.queries.UpdateIsChirpyRed(request.Context(), database.UpdateIsChirpyRedParams{
		ID:          webhook.User.ID,
		IsChirpyRed: true,
		UpdatedAt:   time.Now(),
	})
}

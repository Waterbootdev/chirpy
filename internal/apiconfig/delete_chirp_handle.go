package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/database"
)

func (cfg *ApiConfig) deleteChirpHandle(request *http.Request, chirp *database.Chirp) (*error, error) {
	return nil, cfg.queries.DeleteChirp(request.Context(), chirp.ID)
}

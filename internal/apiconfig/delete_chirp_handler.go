package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) chirpIDValidator(writer http.ResponseWriter, request *http.Request) (*database.Chirp, bool) {
	userID, ok := cfg.validateJWTResponse(request, writer)

	if !ok {
		return nil, false
	}

	chirp, err := cfg.queries.GetChirp(request.Context(), uuid.MustParse(request.PathValue("chirpID")))

	if response.WriteHeaderContentText(err != nil, writer, http.StatusForbidden) {
		return nil, false
	}

	if response.WriteHeaderContentText(chirp.UserID != userID, writer, http.StatusForbidden) {
		return nil, false
	}

	return &chirp, true
}

func (cfg *ApiConfig) deleteChirpHandle(request *http.Request, chirp *database.Chirp) error {
	return cfg.queries.DeleteChirp(request.Context(), chirp.ID)
}

func (cfg *ApiConfig) DeleteChirpHandler(writer http.ResponseWriter, request *http.Request) {
	response.NoContentNoBodyHandler(writer, request, cfg.deleteChirpHandle, cfg.chirpIDValidator)
}

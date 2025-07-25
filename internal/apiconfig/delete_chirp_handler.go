package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/generic_handler"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) chirpIDValidator(writer http.ResponseWriter, request *http.Request) (*database.Chirp, bool) {
	userID, ok := cfg.validateJWTResponse(request, writer)

	if !ok {
		return nil, false
	}

	chirp, err := cfg.queries.GetChirp(request.Context(), uuid.MustParse(request.PathValue("chirpID")))

	if err != nil {
		response.WriteHeaderContentText(writer, response.PLAIN, http.StatusForbidden)
		return nil, false
	}

	ok = chirp.UserID == userID

	if !ok {
		response.WriteHeaderContentText(writer, response.PLAIN, http.StatusForbidden)
		return nil, false
	}

	return &chirp, ok
}

func (cfg *ApiConfig) deleteChirpHandle(request *http.Request, chirp *database.Chirp) error {
	return cfg.queries.DeleteChirp(request.Context(), chirp.ID)
}

func (cfg *ApiConfig) DeleteChirpHandler(writer http.ResponseWriter, request *http.Request) {
	generic_handler.HeaderHandler(writer, request, cfg.deleteChirpHandle, cfg.chirpIDValidator)
}

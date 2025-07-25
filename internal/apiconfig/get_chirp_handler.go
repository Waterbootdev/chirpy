package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

func fromDatabaseChirp(dbChirp *database.Chirp) *chirp {
	return &chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	}
}

func (cfg *ApiConfig) GetChirpHandler(writer http.ResponseWriter, request *http.Request) {
	id := uuid.MustParse(request.PathValue("chirpID"))
	chirp, err := cfg.queries.GetChirp(request.Context(), id)

	if err != nil {
		response.ErrorResponse(writer, http.StatusNotFound, "Chirp not found")
		return
	}

	response.ResponseJsonMarshal(writer, http.StatusOK, fromDatabaseChirp(&chirp))
}

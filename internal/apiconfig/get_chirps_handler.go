package apiconfig

import (
	"context"
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

type SortOrder int

const (
	ASC SortOrder = iota
	DESC
)

func getSortOrder(request *http.Request) SortOrder {

	sort := request.URL.Query().Get("sort")

	if sort == "asc" || (sort != "asc" && sort != "desc") {
		return ASC
	}
	return DESC
}

func getChirps(request *http.Request, getChirps func(context.Context) ([]database.Chirp, error), getChirpsByAuthorID func(context.Context, uuid.UUID) ([]database.Chirp, error)) ([]database.Chirp, error) {
	authorID := request.URL.Query().Get("author_id")
	if authorID == "" {
		return getChirps(request.Context())
	} else {
		return getChirpsByAuthorID(request.Context(), uuid.MustParse(authorID))
	}
}

func (cfg *ApiConfig) getChirpsOrder(sortOrder SortOrder, request *http.Request) ([]database.Chirp, error) {

	switch sortOrder {
	case ASC:
		return getChirps(request, cfg.queries.GetChirpsASC, cfg.queries.GetChirpsASCByUserID)
	case DESC:
		return getChirps(request, cfg.queries.GetChirpsDESC, cfg.queries.GetChirpsDESCByUserID)
	}
	return nil, nil
}

func fromDatabaseChirps(dbChirp []database.Chirp) []chirp {
	chirps := make([]chirp, len(dbChirp))
	for i, dbChirp := range dbChirp {
		chirps[i] = *fromDatabaseChirp(&dbChirp)
	}
	return chirps
}

func (cfg *ApiConfig) GetChirpsHandler(writer http.ResponseWriter, request *http.Request) {

	chirps, err := cfg.getChirpsOrder(getSortOrder(request), request)

	if err != nil {
		response.InternalServerErrorResponse(writer, err)
		return
	}

	response.ResponseJsonMarshal(writer, http.StatusOK, fromDatabaseChirps(chirps))
}

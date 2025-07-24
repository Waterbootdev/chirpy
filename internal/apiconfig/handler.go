package apiconfig

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

const PRINTERROR bool = true

func (cfg *ApiConfig) ResetHandler(writer http.ResponseWriter, request *http.Request) {

	if cfg.platform != "dev" {
		response.ErrorResponse(writer, http.StatusForbidden, "Forbidden")
		return
	}

	cfg.fileserverHits.Store(0)
	response.FprintOKResponse(PRINTERROR, writer, response.PLAIN, "Hits reset")

	if err := cfg.queries.DeleteUsers(request.Context()); err != nil {
		response.InternalServerErrorResponse(writer, err)
	}
}

func (cfg *ApiConfig) MetricsHandler(writer http.ResponseWriter, _ *http.Request) {
	response.FprintfOKResponse(PRINTERROR, writer, response.HTML, METRICSFORMAT, cfg.fileserverHits.Load())
}

func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(writer, request)
	})
}

func (cfg *ApiConfig) getChirps(authorId string, request *http.Request) ([]database.Chirp, error) {
	if authorId == "" {
		return cfg.queries.GetChirps(request.Context())
	} else {
		return cfg.queries.GetChirpsByUserID(request.Context(), uuid.MustParse(authorId))
	}
}

func (cfg *ApiConfig) GetChirpsHandler(writer http.ResponseWriter, request *http.Request) {

	chirps, err := cfg.getChirps(request.URL.Query().Get("author_id"), request)

	if err != nil {
		response.InternalServerErrorResponse(writer, err)
		return
	}

	response.ResponseJsonMarshal(writer, http.StatusOK, fromDatabaseChirps(chirps))

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

func (cfg *ApiConfig) CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	handler(writer, request, cfg.createUserHandle, allways[userRequest], http.StatusCreated)
}

func (cfg *ApiConfig) CreateChirpHandler(writer http.ResponseWriter, request *http.Request) {
	handler(writer, request, cfg.createChirpHandle, cfg.chirpRequestValidator, http.StatusCreated)
}

func (cfg *ApiConfig) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	handler(writer, request, cfg.loginHandle, cfg.loginRequestValidator, http.StatusOK)
}

func (cfg *ApiConfig) RefreshHandler(writer http.ResponseWriter, request *http.Request) {
	headerHandler(writer, request, cfg.refreshHandler, cfg.refreshTokenValidator, http.StatusOK)
}

func (cfg *ApiConfig) RevokeHandler(writer http.ResponseWriter, request *http.Request) {
	token, ok := cfg.getRefreshToken(request)

	if unauthorizedResponse(!ok, writer) {
		return
	}

	err := cfg.queries.RevokeRefreshToken(request.Context(), database.RevokeRefreshTokenParams{
		Token: token.Token,
		RevokedAt: sql.NullTime{
			Valid: true,
			Time:  time.Now(),
		},
	})

	if err == nil {
		response.WriteHeaderContentText(writer, response.PLAIN, http.StatusNoContent)
	} else {
		response.InternalServerErrorResponse(writer, err)
	}
}
func (cfg *ApiConfig) UpdateUserHandler(writer http.ResponseWriter, request *http.Request) {
	handler(writer, request, cfg.updateUserHandle, cfg.updateUserValidator, http.StatusOK)
}

func (cfg *ApiConfig) DeleteChirpHandler(writer http.ResponseWriter, request *http.Request) {
	headerHandler(writer, request, cfg.deleteChirpHandle, cfg.chirpIDValidator, http.StatusNoContent)
}

func (cfg *ApiConfig) WebhookHandler(writer http.ResponseWriter, request *http.Request) {
	handler(writer, request, cfg.webhookHandle, cfg.webhookValidator, http.StatusNoContent)
}

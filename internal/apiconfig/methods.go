package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

const PRINTERROR bool = true

func (cfg *ApiConfig) ResetHandle(writer http.ResponseWriter, request *http.Request) {

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

func (cfg *ApiConfig) MetricsHandle(writer http.ResponseWriter, _ *http.Request) {
	response.FprintfOKResponse(PRINTERROR, writer, response.HTML, METRICSFORMAT, cfg.fileserverHits.Load())
}

func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(writer, request)
	})
}

func (cfg *ApiConfig) CreateUserHandle(writer http.ResponseWriter, request *http.Request) {
	CreateHandle(cfg, writer, request, cfg.CreateUser)
}
func (cfg *ApiConfig) CreateChirpHandle(writer http.ResponseWriter, request *http.Request) {
	CreateHandle(cfg, writer, request, cfg.CreateChirp)
}

func (cfg *ApiConfig) GetChirpsHandle(writer http.ResponseWriter, request *http.Request) {
	chirps, err := cfg.queries.GetChirps(request.Context())

	if err != nil {
		response.InternalServerErrorResponse(writer, err)
		return
	}

	response.ResponseJsonMarshal(writer, http.StatusOK, fromDatabaseChirps(chirps))
}

func (cfg *ApiConfig) GetChirpHandle(writer http.ResponseWriter, request *http.Request) {
	id := uuid.MustParse(request.PathValue("chirpID"))
	chirps, err := cfg.queries.GetChirp(request.Context(), id)

	if err != nil {
		response.InternalServerErrorResponse(writer, err)
		return
	}

	if len(chirps) == 0 {
		response.ErrorResponse(writer, http.StatusNotFound, "Chirp not found")
		return
	}

	response.ResponseJsonMarshal(writer, http.StatusOK, fromDatabaseChirps(chirps))
}

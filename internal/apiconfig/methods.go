package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
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

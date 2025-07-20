package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
)

func (cfg *ApiConfig) ResetHandle(writer http.ResponseWriter, _ *http.Request) {
	cfg.fileserverHits.Store(0)
	response.FprintOKResponse(writer, response.PLAIN, "Hits reset")
}

func (cfg *ApiConfig) MetricsHandler(writer http.ResponseWriter, _ *http.Request) {
	response.FprintfOKResponse(writer, response.HTML, METRICSFORMAT, cfg.fileserverHits.Load())
}

func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(writer, request)
	})
}

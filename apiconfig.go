package main

import (
	"net/http"
	"sync/atomic"

	. "github.com/Waterbootdev/chirpy/internal/response"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) resetHandle(writer http.ResponseWriter, request *http.Request) {
	cfg.fileserverHits.Store(0)
	FprintResponse(writer, WriteHeaderContentTextOK(PLAIN), "Hits reset")
}

const METRICSFORMAT = `<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`

func (cfg *apiConfig) metricsHandler(writer http.ResponseWriter, request *http.Request) {
	FprintfResponse(writer, WriteHeaderContentTextOK(HTML), METRICSFORMAT, cfg.fileserverHits.Load())
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(writer, request)
	})
}

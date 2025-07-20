package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

func healthzHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("OK"))
}

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) resetHandle(writer http.ResponseWriter, request *http.Request) {
	cfg.fileserverHits.Store(0)
	healthzHandler(writer, request)
}

func (cfg *apiConfig) metricsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Hits: %d", cfg.fileserverHits.Load())
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(writer, request)
	})
}

func main() {
	apiCfg := apiConfig{}

	serveMux := http.NewServeMux()

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	httpFileServer := apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir("."))))

	serveMux.Handle("/app/", httpFileServer)
	serveMux.Handle("/app/assets", httpFileServer)

	serveMux.HandleFunc("/healthz", healthzHandler)
	serveMux.HandleFunc("/metrics", apiCfg.metricsHandler)
	serveMux.HandleFunc("/reset", apiCfg.resetHandle)

	httpServer.ListenAndServe()
}

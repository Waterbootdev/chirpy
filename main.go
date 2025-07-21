package main

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/apiconfig"
)

func main() {
	apiCfg := apiconfig.ApiConfig{}

	serveMux := http.NewServeMux()

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	httpFileServer := apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir("."))))

	serveMux.Handle("/app/", httpFileServer)
	serveMux.Handle("/app/assets", httpFileServer)

	serveMux.HandleFunc("GET /api/healthz", healthzHandler)
	serveMux.HandleFunc("GET /admin/metrics", apiCfg.MetricsHandler)
	serveMux.HandleFunc("POST /admin/reset", apiCfg.ResetHandle)
	serveMux.HandleFunc("POST /api/validate_chirp", validateChirpLength)

	httpServer.ListenAndServe()
}

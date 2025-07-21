package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Waterbootdev/chirpy/internal/apiconfig"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiCfg := apiconfig.NewApiConfig()

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

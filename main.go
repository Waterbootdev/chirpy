package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Waterbootdev/chirpy/internal/apiconfig"
	"github.com/joho/godotenv"
)

func setFileServerHandle(serveMux *http.ServeMux, apiCfg *apiconfig.ApiConfig) {

	httpFileServer := apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir("."))))

	serveMux.Handle("/app/", httpFileServer)
	serveMux.Handle("/app/assets", httpFileServer)
}

func setAdminHandleFuncs(serveMux *http.ServeMux, apiCfg *apiconfig.ApiConfig) {
	serveMux.HandleFunc("GET /admin/metrics", apiCfg.MetricsHandler)
	serveMux.HandleFunc("POST /admin/reset", apiCfg.ResetHandle)

}

func setApiHandleFuncs(serveMux *http.ServeMux) {
	serveMux.HandleFunc("GET /api/healthz", healthzHandler)
	serveMux.HandleFunc("POST /api/validate_chirp", validateChirpLengthAndCleanProfaneWords)
}

func newServeMux(apiCfg *apiconfig.ApiConfig) *http.ServeMux {
	serveMux := http.NewServeMux()
	setFileServerHandle(serveMux, apiCfg)
	setAdminHandleFuncs(serveMux, apiCfg)
	setApiHandleFuncs(serveMux)
	return serveMux
}

func httpServer(apiCfg *apiconfig.ApiConfig) (httpServer *http.Server) {
	return &http.Server{
		Addr:    ":8080",
		Handler: newServeMux(apiCfg),
	}
}

func main() {
	godotenv.Load()

	httpServer(apiconfig.NewApiConfig()).ListenAndServe()
}

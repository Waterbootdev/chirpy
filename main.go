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
	serveMux.HandleFunc("POST /admin/reset", apiCfg.ResetHandler)
}

func setApiHandleFuncs(serveMux *http.ServeMux, apiCfg *apiconfig.ApiConfig) {
	serveMux.HandleFunc("GET /api/healthz", healthzHandler)
	serveMux.HandleFunc("POST /api/chirps", apiCfg.CreateChirpHandler)
	serveMux.HandleFunc("POST /api/users", apiCfg.CreateUserHandler)
	serveMux.HandleFunc("GET /api/chirps", apiCfg.GetChirpsHandler)
	serveMux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.GetChirpHandler)
	serveMux.HandleFunc("POST /api/login", apiCfg.LoginHandler)
	serveMux.HandleFunc("POST /api/refresh", apiCfg.RefreshHandler)
	serveMux.HandleFunc("POST /api/revoke", apiCfg.RevokeHandler)
}
func newServeMux(apiCfg *apiconfig.ApiConfig) *http.ServeMux {
	serveMux := http.NewServeMux()
	setFileServerHandle(serveMux, apiCfg)
	setAdminHandleFuncs(serveMux, apiCfg)
	setApiHandleFuncs(serveMux, apiCfg)
	return serveMux
}

func newHttpServer(apiCfg *apiconfig.ApiConfig) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: newServeMux(apiCfg),
	}
}

func main() {
	godotenv.Load()

	newHttpServer(apiconfig.NewApiConfig()).ListenAndServe()
}

package main

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/apiconfig"
	"github.com/Waterbootdev/chirpy/internal/response"
)

func healthzHandler(writer http.ResponseWriter, request *http.Request) {
	response.FprintOKResponse(apiconfig.PRINTERROR, writer, response.PLAIN, "OK")
}

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

	httpServer.ListenAndServe()
}

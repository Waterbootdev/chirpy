package main

import (
	"net/http"

	. "github.com/Waterbootdev/chirpy/internal/response"
)

func healthzHandler(writer http.ResponseWriter, request *http.Request) {
	FprintResponse(writer, WriteHeaderContentTextPlainOK, "OK")
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

package main

import (
	"net/http"
)

func healthzHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("OK"))
}
func main() {
	serveMux := http.NewServeMux()

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	httpFileServer := http.StripPrefix("/app/", http.FileServer(http.Dir(".")))

	serveMux.HandleFunc("/app/", httpFileServer.ServeHTTP)
	serveMux.HandleFunc("/app/assets", httpFileServer.ServeHTTP)

	serveMux.HandleFunc("/healthz", healthzHandler)

	httpServer.ListenAndServe()
}

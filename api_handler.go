package main

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/apiconfig"
	"github.com/Waterbootdev/chirpy/internal/response"
)

func healthzHandler(writer http.ResponseWriter, request *http.Request) {
	response.FprintOKResponse(apiconfig.PRINTERROR, writer, response.PLAIN, "OK")
}

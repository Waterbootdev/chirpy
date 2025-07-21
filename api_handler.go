package main

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/apiconfig"
	"github.com/Waterbootdev/chirpy/internal/chirp"
	"github.com/Waterbootdev/chirpy/internal/response"
)

func healthzHandler(writer http.ResponseWriter, request *http.Request) {
	response.FprintOKResponse(apiconfig.PRINTERROR, writer, response.PLAIN, "OK")
}

func validateChirpLength(writer http.ResponseWriter, request *http.Request) {

	chirpFromRequest, err := chirp.ChirpFromRequest(request)

	if err != nil {
		chirp.ChirpErrorResponse(writer, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if chirpFromRequest.IsToLong() {
		chirp.ChirpErrorResponse(writer, http.StatusBadRequest, "Chirp is too long")
		return
	}

	chirp.ChirpCleanedResponse(writer, chirpFromRequest.CleanProfaneWords(chirp.CurrentProfaneWords()))
}

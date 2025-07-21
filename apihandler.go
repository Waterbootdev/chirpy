package main

import (
	"encoding/json"
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/apiconfig"
	"github.com/Waterbootdev/chirpy/internal/response"
)

func healthzHandler(writer http.ResponseWriter, request *http.Request) {
	response.FprintOKResponse(apiconfig.PRINTERROR, writer, response.PLAIN, "OK")
}

func chirpErrorResponse(writer http.ResponseWriter, statusCode int, currentChirpError string) {
	data, err := json.Marshal(chirpError{Error: currentChirpError})
	if err != nil {
		response.InternalServerErrorResponse(writer, err)
		return
	}
	response.JsonResponse(writer, statusCode, data)
}

func chirpCleanedResponse(writer http.ResponseWriter, cleanedBody string) {
	data, err := json.Marshal(cleanedChirp{CleanedBody: cleanedBody})

	if err != nil {
		response.InternalServerErrorResponse(writer, err)
		return
	}
	response.JsonResponse(writer, http.StatusOK, data)
}

func validateChirpLength(writer http.ResponseWriter, request *http.Request) {

	chirp, err := getChirp(request)

	if err != nil {
		chirpErrorResponse(writer, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if len(chirp.Body) > MAX_CHIRP_LENGTH {
		chirpErrorResponse(writer, http.StatusBadRequest, "Chirp is too long")
		return
	}

	chirpCleanedResponse(writer, getCleanChirp(chirp.Body, ProfaneWords()))
}

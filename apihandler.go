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

func chirpJsonResponse(writer http.ResponseWriter, statusCode int, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		response.InternalServerErrorResponse(writer, err)
		return
	}
	response.JsonResponse(writer, statusCode, data)
}

func chirpErrorResponse(writer http.ResponseWriter, statusCode int, currentChirpError string) {
	chirpJsonResponse(writer, statusCode, chirpError{Error: currentChirpError})
}
func chirpCleanedResponse(writer http.ResponseWriter, cleanedBody string) {
	chirpJsonResponse(writer, http.StatusOK, cleanedChirp{CleanedBody: cleanedBody})
}

func validateChirpLength(writer http.ResponseWriter, request *http.Request) {

	chirp, err := chirpFromRequest(request)

	if err != nil {
		chirpErrorResponse(writer, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if len(chirp.Body) > MAX_CHIRP_LENGTH {
		chirpErrorResponse(writer, http.StatusBadRequest, "Chirp is too long")
		return
	}

	chirpCleanedResponse(writer, chirp.cleanProfaneWords(currentProfaneWords()))
}

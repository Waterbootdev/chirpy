package chirp

import (
	"encoding/json"
	"net/http"
)

func chirpFromRequestErrorResponse(writer http.ResponseWriter, request *http.Request) (chirp Chirp, wasError bool) {
	decoder := json.NewDecoder(request.Body)
	wasError = decoder.Decode(&chirp) != nil

	if wasError {
		errorResponse(writer, http.StatusInternalServerError, "Something went wrong")
	}

	return chirp, wasError
}

func validatedChirpFromRequestNotValidResponse(writer http.ResponseWriter, request *http.Request) (Chirp, bool) {
	chirpFromRequest, wasError := chirpFromRequestErrorResponse(writer, request)
	notValid := wasError || chirpFromRequest.isToLongErrorResponse(writer)
	return chirpFromRequest, !notValid
}

func ValidateChirpLengthAndCleanProfaneWords(writer http.ResponseWriter, request *http.Request) {
	if chirpFromRequest, valid := validatedChirpFromRequestNotValidResponse(writer, request); valid {
		chirpFromRequest.cleanedResponse(writer)
	}
}

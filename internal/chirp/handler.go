package chirp

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
)

func validatedChirpFromRequestNotValidResponse(writer http.ResponseWriter, request *http.Request) (*Chirp, bool) {
	chirpFromRequest, ok := response.FromRequestErrorResponse[Chirp](writer, request)
	valid := ok && !chirpFromRequest.isToLongErrorResponse(writer)
	return chirpFromRequest, valid
}

func ValidateChirpLengthAndCleanProfaneWords(writer http.ResponseWriter, request *http.Request) {
	if chirpFromRequest, valid := validatedChirpFromRequestNotValidResponse(writer, request); valid {
		chirpFromRequest.cleanedResponse(writer)
	}
}

package chirp

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
)

func errorResponse(writer http.ResponseWriter, statusCode int, currentChirpError string) {
	response.ResponseJsonMarshal(writer, statusCode, chirpError{Error: currentChirpError})
}
func cleanedResponse(writer http.ResponseWriter, cleanedBody string) {
	response.ResponseJsonMarshal(writer, http.StatusOK, cleanedChirp{CleanedBody: cleanedBody})
}

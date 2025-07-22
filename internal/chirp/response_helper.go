package chirp

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
)

func cleanedResponse(writer http.ResponseWriter, cleanedBody string) {
	response.ResponseJsonMarshal(writer, http.StatusOK, cleanedChirp{CleanedBody: cleanedBody})
}

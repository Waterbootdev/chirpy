package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) validateJWTResponse(request *http.Request, writer http.ResponseWriter) (uuid.UUID, bool) {
	userID, err := cfg.validateJWT(request)
	if response.ErrorResponse(err != nil, writer, http.StatusUnauthorized, "Unauthorized") {
		return userID, false
	}
	return userID, true
}

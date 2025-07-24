package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
)

func unauthorizedResponse(unauthorized bool, writer http.ResponseWriter) bool {
	if unauthorized {
		response.ErrorResponse(writer, http.StatusUnauthorized, "Unauthorized")
	}
	return unauthorized
}

func (cfg *ApiConfig) getRefreshToken(request *http.Request) (*database.RefreshToken, bool) {

	token, err := auth.GetBearerToken(request.Header)

	if err != nil {
		return nil, false
	}

	dbToken, err := cfg.queries.GetRefreshToken(request.Context(), token)

	if err != nil {
		return nil, false
	}

	return &dbToken, true
}

func (cfg *ApiConfig) refreshTokenValidator(writer http.ResponseWriter, request *http.Request) (*database.RefreshToken, bool) {

	refreshToken, ok := cfg.getRefreshToken(request)

	if unauthorizedResponse(!ok, writer) {
		return nil, false
	}

	if unauthorizedResponse(refreshToken.RevokedAt.Valid || refreshToken.ExpiresAt.Before(time.Now()), writer) {
		return nil, false
	}

	return refreshToken, true
}

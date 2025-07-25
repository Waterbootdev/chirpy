package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
)

type accesToken struct {
	Token string `json:"token"`
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

	if response.UnauthorizedResponse(!ok, writer) {
		return nil, false
	}

	if response.UnauthorizedResponse(refreshToken.RevokedAt.Valid || refreshToken.ExpiresAt.Before(time.Now()), writer) {
		return nil, false
	}

	return refreshToken, true
}

func (cfg *ApiConfig) refreshHandle(_ *http.Request, refreshToken *database.RefreshToken) (*accesToken, error) {

	token, err := cfg.makeJWT(refreshToken.UserID)

	if err != nil {
		return nil, err
	}

	return &accesToken{Token: token}, nil
}

func (cfg *ApiConfig) RefreshHandler(writer http.ResponseWriter, request *http.Request) {
	response.ContentNoBodyHandler(writer, request, cfg.refreshHandle, cfg.refreshTokenValidator, http.StatusOK)
}

package apiconfig

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
)

const EXPIRE_DURATION = 60 * 24 * time.Hour

func (cfg *ApiConfig) loginHandle(request *http.Request, loginRequest *loginRequest) (*userToken, error) {

	accesToken, err := cfg.makeJWT(loginRequest.User.ID)

	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.MakeRefreshToken()

	if err != nil {
		return nil, err
	}
	timeNow := time.Now()

	databaseRefreshToken, err := cfg.queries.CreateRefreshToken(request.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		UserID:    loginRequest.User.ID,
		ExpiresAt: timeNow.Add(EXPIRE_DURATION),
		RevokedAt: sql.NullTime{
			Valid: false,
			Time:  time.Time{},
		},
	})

	if err != nil {
		return nil, err
	}

	return fromDatabaseUserToken(loginRequest, &databaseRefreshToken, accesToken), nil
}

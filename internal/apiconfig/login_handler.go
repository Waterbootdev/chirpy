package apiconfig

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/generic_handler"
	"github.com/Waterbootdev/chirpy/internal/response"
)

const EXPIRE_DURATION = 60 * 24 * time.Hour

type loginRequest struct {
	Password string         `json:"password"`
	Email    string         `json:"email"`
	User     *database.User `json:"user"`
}

type userToken struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	IsChirpyRed  bool      `json:"is_chirpy_red"`
}

func fromDatabaseUserToken(loginRequest *loginRequest, refreshToken *database.RefreshToken, accesToken string) *userToken {
	return &userToken{
		ID:           loginRequest.User.ID,
		CreatedAt:    refreshToken.CreatedAt,
		UpdatedAt:    refreshToken.UpdatedAt,
		Email:        loginRequest.Email,
		Token:        accesToken,
		RefreshToken: refreshToken.Token,
		IsChirpyRed:  loginRequest.User.IsChirpyRed,
	}
}

func (cfg *ApiConfig) loginRequestValidator(writer http.ResponseWriter, request *http.Request, loginRequest *loginRequest) (ok bool) {

	user, err := cfg.queries.GetUserByEmail(context.Background(), loginRequest.Email)

	if ok = err == nil; !ok {
		response.ErrorResponse(writer, http.StatusUnauthorized, "Incorrect email or password")
		return ok
	}

	err = auth.CheckPasswordHash(loginRequest.Password, user.PasswordHash)

	if ok = err == nil; !ok {
		response.ErrorResponse(writer, http.StatusUnauthorized, "Incorrect email or password")
		return ok
	}

	loginRequest.User = &user

	return ok
}

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

func (cfg *ApiConfig) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	generic_handler.ContentBodyHandler(writer, request, cfg.loginHandle, cfg.loginRequestValidator, http.StatusOK)
}

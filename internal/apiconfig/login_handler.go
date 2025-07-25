package apiconfig

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
)

const EXPIRE_DURATION = 60 * 24 * time.Hour

type loginRequest struct {
	Password     string         `json:"password"`
	Email        string         `json:"email"`
	User         *database.User `json:"user"`
	RefreshToken string         `json:"refresh_token"`
	AccesToken   string         `json:"acces_token"`
	At           time.Time      `json:"at"`
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

func (cfg *ApiConfig) loginRequestValidator(writer http.ResponseWriter, request *http.Request, loginRequest *loginRequest) bool {

	user, err := cfg.queries.GetUserByEmail(context.Background(), loginRequest.Email)

	if response.ErrorResponse(err != nil, writer, http.StatusUnauthorized, "Incorrect email or password") {
		return false
	}

	err = auth.CheckPasswordHash(loginRequest.Password, user.PasswordHash)

	if response.ErrorResponse(err != nil, writer, http.StatusUnauthorized, "Incorrect email or password") {
		return false
	}

	refreshToken, err := auth.MakeRefreshToken()

	if response.ErrorResponse(err != nil, writer, http.StatusInternalServerError, "Internal server error") {
		return false
	}

	accesToken, err := cfg.makeJWT(user.ID)

	if response.ErrorResponse(err != nil, writer, http.StatusInternalServerError, "Internal server error") {
		return false
	}

	loginRequest.User = &user
	loginRequest.At = time.Now()
	loginRequest.RefreshToken = refreshToken
	loginRequest.AccesToken = accesToken

	return true
}

func (cfg *ApiConfig) createRefreshToken(request *http.Request, loginRequest *loginRequest) (*database.RefreshToken, error) {

	databaseRefreshToken, err := cfg.queries.CreateRefreshToken(request.Context(), database.CreateRefreshTokenParams{
		Token:     loginRequest.RefreshToken,
		CreatedAt: loginRequest.At,
		UpdatedAt: loginRequest.At,
		UserID:    loginRequest.User.ID,
		ExpiresAt: loginRequest.At.Add(EXPIRE_DURATION),
		RevokedAt: sql.NullTime{
			Valid: false,
			Time:  time.Time{},
		},
	})

	return &databaseRefreshToken, err
}

func (cfg *ApiConfig) loginHandle(request *http.Request, loginRequest *loginRequest) (*userToken, error) {
	accesToken := loginRequest.AccesToken
	loginRequest.AccesToken = ""
	databaseRefreshToken, err := cfg.createRefreshToken(request, loginRequest)
	return fromDatabaseUserToken(loginRequest, databaseRefreshToken, accesToken), err
}

func (cfg *ApiConfig) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	response.ContentBodyHandler(writer, request, cfg.loginHandle, cfg.loginRequestValidator, http.StatusOK)
}

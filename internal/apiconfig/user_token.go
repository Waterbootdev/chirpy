package apiconfig

import (
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

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

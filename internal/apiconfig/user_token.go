package apiconfig

import (
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

type userToken struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
}

func fromDatabaseUserToken(dbUser *database.User, token string) *userToken {
	return &userToken{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
		Token:     token,
	}
}

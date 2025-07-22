package apiconfig

import (
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

type email struct {
	Email string `json:"email"`
}

type user struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func fromChirpy(chirpy *database.Chirpy) *user {
	return &user{
		ID:        chirpy.ID,
		CreatedAt: chirpy.CreatedAt,
		UpdatedAt: chirpy.UpdatedAt,
		Email:     chirpy.Email,
	}
}

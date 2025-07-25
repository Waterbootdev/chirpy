package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

type userRequest struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	ID       uuid.UUID `json:"id"`
	At       time.Time `json:"at"`
}

func (cfg *ApiConfig) createUser(request *http.Request, userRequest *userRequest) (*user, error) {
	c, err := cfg.queries.CreateUser(request.Context(), database.CreateUserParams{
		ID:           userRequest.ID,
		CreatedAt:    userRequest.At,
		UpdatedAt:    userRequest.At,
		Email:        userRequest.Email,
		PasswordHash: userRequest.Password,
	})
	return fromDatabaseUser(&c), err
}

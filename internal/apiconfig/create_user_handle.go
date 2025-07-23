package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) createUserHandle(request *http.Request, userRequest *userRequest) (*user, error) {
	password, err := auth.HashPassword(userRequest.Password)

	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	c, err := cfg.queries.CreateUser(request.Context(), database.CreateUserParams{
		ID:           uuid.New(),
		CreatedAt:    timeNow,
		UpdatedAt:    timeNow,
		Email:        userRequest.Email,
		PasswordHash: password,
	})
	return fromDatabaseUser(&c), err
}

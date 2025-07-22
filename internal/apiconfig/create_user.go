package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) CreateUser(request *http.Request, userRequest *userRequest) (*user, error) {
	timeNow := time.Now()
	c, err := cfg.queries.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Email:     userRequest.Email,
	})
	return fromDatabaseUser(&c), err
}

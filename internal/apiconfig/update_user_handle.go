package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
)

func (cfg *ApiConfig) updateUserHandle(request *http.Request, userRequest *updateUserRequest) (*user, error) {

	if hash, err := auth.HashPassword(userRequest.Password); err == nil {
		if dbUser, err := cfg.queries.UpdateUser(request.Context(), database.UpdateUserParams{
			ID:           userRequest.UserID,
			Email:        userRequest.Email,
			PasswordHash: hash,
			UpdatedAt:    time.Now()}); err == nil {
			return fromDatabaseUser(&dbUser), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

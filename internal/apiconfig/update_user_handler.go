package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/generic_handler"
	"github.com/google/uuid"
)

type updateUserRequest struct {
	UserID   uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (cfg *ApiConfig) updateUserValidator(writer http.ResponseWriter, request *http.Request, userRequest *updateUserRequest) bool {

	userID, ok := cfg.validateJWTResponse(request, writer)

	if ok {
		userRequest.UserID = userID
	}

	return ok
}

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

func (cfg *ApiConfig) UpdateUserHandler(writer http.ResponseWriter, request *http.Request) {
	generic_handler.HandlerBody(writer, request, cfg.updateUserHandle, cfg.updateUserValidator, http.StatusOK)
}

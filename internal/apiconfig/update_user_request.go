package apiconfig

import (
	"net/http"

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

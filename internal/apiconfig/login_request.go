package apiconfig

import (
	"context"
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
)

type loginRequest struct {
	Password         string         `json:"password"`
	Email            string         `json:"email"`
	ExpiresInSeconds int            `json:"expires_in_seconds"`
	User             *database.User `json:"user"`
}

func loginRequestValidator(cfg *ApiConfig, writer http.ResponseWriter, request *http.Request, loginRequest *loginRequest) (ok bool) {

	user, err := cfg.queries.GetUserByEmail(context.Background(), loginRequest.Email)

	if ok = err == nil; !ok {
		response.ErrorResponse(writer, http.StatusUnauthorized, "Incorrect email or password")
		return ok
	}

	err = auth.CheckPasswordHash(loginRequest.Password, user.PasswordHash)

	if ok = err == nil; !ok {
		response.ErrorResponse(writer, http.StatusUnauthorized, "Incorrect email or password")
		return ok
	}

	loginRequest.User = &user

	return ok
}

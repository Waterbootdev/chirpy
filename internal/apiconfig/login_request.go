package apiconfig

import (
	"context"
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

type loginRequest struct {
	Password string    `json:"password"`
	Email    string    `json:"email"`
	ID       uuid.UUID `json:"id"`
}

func loginRequestValidator(cfg *ApiConfig, writer http.ResponseWriter, loginRequest *loginRequest) (ok bool) {

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

	loginRequest.ID = user.ID

	return ok
}

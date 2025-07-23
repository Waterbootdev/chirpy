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

func (lr loginRequest) IsValidResponse(cfg *ApiConfig, writer http.ResponseWriter) (lrp *loginRequest, ok bool) {
	lrp = &lr

	user, err := cfg.queries.GetUserByEmail(context.Background(), lr.Email)

	if ok = err == nil; !ok {
		response.ErrorResponse(writer, http.StatusUnauthorized, "Incorrect email or password")
		return lrp, ok
	}

	err = auth.CheckPasswordHash(lr.Password, user.PasswordHash)

	if ok = err == nil; !ok {
		response.ErrorResponse(writer, http.StatusUnauthorized, "Incorrect email or password")
		return lrp, ok
	}

	lr.ID = user.ID

	return lrp, ok
}

package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) createUserValidator(writer http.ResponseWriter, request *http.Request, userRequest *userRequest) bool {

	password, err := auth.HashPassword(userRequest.Password)

	if response.ErrorResponse(err != nil, writer, http.StatusBadRequest, "Invalid password") {
		return false
	}

	userRequest.Password = password

	return true
}

func (cfg *ApiConfig) createUserHandle(request *http.Request, userRequest *userRequest) (*user, error) {
	userRequest.ID = uuid.New()
	userRequest.At = time.Now()
	return cfg.createUser(request, userRequest)
}

func (cfg *ApiConfig) CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	response.ContentBodyHandler(writer, request, cfg.createUserHandle, cfg.createUserValidator, http.StatusCreated)
}

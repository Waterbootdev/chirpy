package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

const PRINTERROR bool = true

func (cfg *ApiConfig) ResetHandle(writer http.ResponseWriter, request *http.Request) {

	if cfg.platform != "dev" {
		response.ErrorResponse(writer, http.StatusForbidden, "Forbidden")
		return
	}

	cfg.fileserverHits.Store(0)
	response.FprintOKResponse(PRINTERROR, writer, response.PLAIN, "Hits reset")

	if err := cfg.queries.DeleteUsers(request.Context()); err != nil {
		response.InternalServerErrorResponse(writer, err)
	}
}

func (cfg *ApiConfig) MetricsHandle(writer http.ResponseWriter, _ *http.Request) {
	response.FprintfOKResponse(PRINTERROR, writer, response.HTML, METRICSFORMAT, cfg.fileserverHits.Load())
}

func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(writer, request)
	})
}

func (cfg *ApiConfig) CreateUser(request *http.Request, email *email) (*database.User, error) {
	timeNow := time.Now()
	c, err := cfg.queries.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Email:     email.Email,
	})
	return &c, err
}

func (cfg *ApiConfig) CreateUserErrorResponse(writer http.ResponseWriter, request *http.Request, email *email) (*user, bool) {
	chirpy, err := cfg.CreateUser(request, email)
	wasError := err != nil
	if wasError {
		response.InternalServerErrorResponse(writer, err)
	}
	return fromDatabase(chirpy), !wasError
}

func (cfg *ApiConfig) CreateUserFromRequest(writer http.ResponseWriter, request *http.Request) (*user, bool) {
	if emailFromRequest, ok := response.FromRequestErrorResponse[email](writer, request); ok {
		return cfg.CreateUserErrorResponse(writer, request, emailFromRequest)
	} else {
		return nil, ok
	}
}

func (cfg *ApiConfig) CreateUserHandle(writer http.ResponseWriter, request *http.Request) {
	if user, ok := cfg.CreateUserFromRequest(writer, request); ok {
		response.ResponseJsonMarshal(writer, http.StatusCreated, user)
	}
}

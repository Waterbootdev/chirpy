package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
)

type Validator[T any] interface {
	IsValidResponse(http.ResponseWriter) bool
}

func CreateErrorResponse[T any, U any](cfg *ApiConfig, writer http.ResponseWriter, request *http.Request, t *T, create func(request *http.Request, t *T) (*U, error)) (*U, bool) {
	u, err := create(request, t)
	wasError := err != nil
	if wasError {
		response.InternalServerErrorResponse(writer, err)
	}
	return u, !wasError
}

func CreateFromRequest[T Validator[T], U any](cfg *ApiConfig, writer http.ResponseWriter, request *http.Request, create func(request *http.Request, t *T) (*U, error)) (*U, bool) {
	if t, ok := response.FromRequestErrorResponse[T](writer, request); ok && (*t).IsValidResponse(writer) {
		return CreateErrorResponse(cfg, writer, request, t, create)
	} else {
		return nil, ok
	}
}

func CreateHandle[T Validator[T], U any](cfg *ApiConfig, writer http.ResponseWriter, request *http.Request, create func(request *http.Request, t *T) (*U, error)) {
	if u, ok := CreateFromRequest(cfg, writer, request, create); ok {
		response.ResponseJsonMarshal(writer, http.StatusCreated, u)
	}
}

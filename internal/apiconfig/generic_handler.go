package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
)

type Validator[T any] interface {
	IsValidResponse(*ApiConfig, http.ResponseWriter) (*T, bool)
}

func handleErrorResponse[T any, U any](writer http.ResponseWriter, request *http.Request, t *T, handle func(request *http.Request, t *T) (*U, error)) (*U, bool) {
	u, err := handle(request, t)
	wasError := err != nil
	if wasError {
		response.InternalServerErrorResponse(writer, err)
	}
	return u, !wasError
}

func handelFromRequest[T Validator[T], U any](cfg *ApiConfig, writer http.ResponseWriter, request *http.Request, handle func(request *http.Request, t *T) (*U, error)) (*U, bool) {
	if t, ok := response.FromRequestErrorResponse[T](writer, request); ok {
		if t, ok := (*t).IsValidResponse(cfg, writer); ok {
			return handleErrorResponse(writer, request, t, handle)
		} else {
			return nil, false
		}
	} else {
		return nil, false
	}
}

func handler[T Validator[T], U any](cfg *ApiConfig, writer http.ResponseWriter, request *http.Request, handel func(request *http.Request, t *T) (*U, error), statusCode int) {
	if u, ok := handelFromRequest(cfg, writer, request, handel); ok {
		response.ResponseJsonMarshal(writer, statusCode, u)
	}
}

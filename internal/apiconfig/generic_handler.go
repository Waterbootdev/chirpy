package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
)

func handleErrorResponse[T, U any](writer http.ResponseWriter, request *http.Request, t *T, handle func(request *http.Request, t *T) (*U, error)) (*U, bool) {
	u, err := handle(request, t)
	wasError := err != nil
	if wasError {
		response.InternalServerErrorResponse(writer, err)
	}
	return u, !wasError
}

func handelFromRequest[T, U any](cfg *ApiConfig, writer http.ResponseWriter, request *http.Request, handle func(request *http.Request, t *T) (*U, error), requestValidator func(cfg *ApiConfig, writer http.ResponseWriter, t *T) bool) (*U, bool) {
	if t, ok := response.FromRequestErrorResponse[T](writer, request); ok && requestValidator(cfg, writer, t) {
		return handleErrorResponse(writer, request, t, handle)
	} else {
		return nil, false
	}
}

func handler[T, U any](cfg *ApiConfig, writer http.ResponseWriter, request *http.Request, handel func(request *http.Request, t *T) (*U, error), requestValidator func(cfg *ApiConfig, writer http.ResponseWriter, t *T) bool, statusCode int) {
	if u, ok := handelFromRequest(cfg, writer, request, handel, requestValidator); ok {
		response.ResponseJsonMarshal(writer, statusCode, u)
	}
}

func allways[T any](cfg *ApiConfig, writer http.ResponseWriter, t *T) bool { return true }

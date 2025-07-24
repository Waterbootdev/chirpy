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

func handelFromRequest[T, U any](writer http.ResponseWriter, request *http.Request, handle func(request *http.Request, t *T) (*U, error), requestValidator func(writer http.ResponseWriter, request *http.Request, t *T) bool) (*U, bool) {
	if t, ok := response.FromRequestErrorResponse[T](writer, request); ok && requestValidator(writer, request, t) {
		return handleErrorResponse(writer, request, t, handle)
	} else {
		return nil, false
	}
}

func handler[T, U any](writer http.ResponseWriter, request *http.Request, handel func(request *http.Request, t *T) (*U, error), requestValidator func(writer http.ResponseWriter, request *http.Request, t *T) bool, statusCode int) {
	if u, ok := handelFromRequest(writer, request, handel, requestValidator); ok {
		response.ResponseJsonMarshal(writer, statusCode, u)
	}
}

func allways[T any](writer http.ResponseWriter, request *http.Request, t *T) bool {
	return true
}

func headerHandler[T, U any](writer http.ResponseWriter, request *http.Request, handel func(request *http.Request, t *T) (*U, error), handelValidator func(writer http.ResponseWriter, request *http.Request) (*T, bool), statusCode int) {
	if t, ok := handelValidator(writer, request); ok {
		if u, ok := handleErrorResponse(writer, request, t, handel); ok {
			response.ResponseJsonMarshal(writer, statusCode, u)
		}
	}
}

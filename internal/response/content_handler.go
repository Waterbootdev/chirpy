package response

import (
	"net/http"
)

func handleBodyErrorResponse[T, U any](writer http.ResponseWriter, request *http.Request, t *T, handle func(request *http.Request, t *T) (*U, error)) (*U, bool) {
	u, err := handle(request, t)
	wasError := err != nil
	if wasError {
		InternalServerErrorResponse(writer, err)
	}
	return u, !wasError
}

func handelBodyFromRequest[T, U any](writer http.ResponseWriter, request *http.Request, handle func(request *http.Request, t *T) (*U, error), requestValidator func(writer http.ResponseWriter, request *http.Request, t *T) bool) (*U, bool) {
	if t, ok := FromRequestErrorResponse[T](writer, request); ok && requestValidator(writer, request, t) {
		return handleBodyErrorResponse(writer, request, t, handle)
	} else {
		return nil, false
	}
}

func ContentBodyHandler[T, U any](writer http.ResponseWriter, request *http.Request, handel func(request *http.Request, t *T) (*U, error), requestValidator func(writer http.ResponseWriter, request *http.Request, t *T) bool, statusCode int) {
	if u, ok := handelBodyFromRequest(writer, request, handel, requestValidator); ok {
		ResponseJsonMarshal(writer, statusCode, u)
	}
}

func ContentNoBodyHandler[T, U any](writer http.ResponseWriter, request *http.Request, handel func(request *http.Request, t *T) (*U, error), handelValidator func(writer http.ResponseWriter, request *http.Request) (*T, bool), statusCode int) {
	if t, ok := handelValidator(writer, request); ok {
		if u, ok := handleBodyErrorResponse(writer, request, t, handel); ok {
			ResponseJsonMarshal(writer, statusCode, u)
		}
	}
}

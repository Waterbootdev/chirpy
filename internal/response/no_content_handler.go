package response

import (
	"net/http"
)

func handleErrorResponse[T any](writer http.ResponseWriter, request *http.Request, t *T, handle func(request *http.Request, t *T) error) bool {
	err := handle(request, t)
	wasError := err != nil
	if wasError {
		InternalServerErrorResponse(writer, err)
	}
	return !wasError
}

func handelFromRequest[T any](writer http.ResponseWriter, request *http.Request, handle func(request *http.Request, t *T) error, requestValidator func(writer http.ResponseWriter, request *http.Request, t *T) bool) bool {
	if t, ok := FromRequestErrorResponse[T](writer, request); ok && requestValidator(writer, request, t) {
		return handleErrorResponse(writer, request, t, handle)
	} else {
		return false
	}
}

func NoContentBodyHandler[T any](writer http.ResponseWriter, request *http.Request, handel func(request *http.Request, t *T) error, requestValidator func(writer http.ResponseWriter, request *http.Request, t *T) bool) {
	if ok := handelFromRequest(writer, request, handel, requestValidator); ok {
		WriteHeaderContentText(writer, PLAIN, http.StatusNoContent)
	}
}

func NoContentNoBodyHandler[T any](writer http.ResponseWriter, request *http.Request, handel func(request *http.Request, t *T) error, handelValidator func(writer http.ResponseWriter, request *http.Request) (*T, bool)) {
	if t, ok := handelValidator(writer, request); ok {
		if ok := handleErrorResponse(writer, request, t, handel); ok {
			WriteHeaderContentText(writer, PLAIN, http.StatusNoContent)
		}
	}
}

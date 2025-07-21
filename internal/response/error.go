package response

import "net/http"

func InternalServerErrorResponse(writer http.ResponseWriter, err error) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Write([]byte(err.Error()))
}

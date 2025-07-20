package response

import "net/http"

func WriteHeaderContentTextPlainOK(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
}

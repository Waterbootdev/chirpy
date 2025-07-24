package response

import (
	"fmt"
	"net/http"
)

const PLAIN = "plain"
const HTML = "html"

func WriteHeaderContentTextOK(textType string) func(http.ResponseWriter) {
	return func(writer http.ResponseWriter) {
		writeHeaderContentTextOK(writer, textType)
	}
}

func writeHeaderContentTextOK(writer http.ResponseWriter, textType string) {
	writer.Header().Set("Content-Type", fmt.Sprintf("text/%s; charset=utf-8", textType))
	writer.WriteHeader(http.StatusOK)
}

func WriteHeaderContentText(writer http.ResponseWriter, textType string, statusCode int) {
	writer.Header().Set("Content-Type", fmt.Sprintf("text/%s; charset=utf-8", textType))
	writer.WriteHeader(statusCode)
}

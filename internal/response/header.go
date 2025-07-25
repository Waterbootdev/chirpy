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

func writeHeaderContentText(writer http.ResponseWriter, textType string, statusCode int) {
	writer.Header().Set("Content-Type", fmt.Sprintf("text/%s; charset=utf-8", textType))
	writer.WriteHeader(statusCode)
}

func WriteHeaderContentText(notOk bool, writer http.ResponseWriter, statusCode int) bool {
	if notOk {
		writeHeaderContentText(writer, PLAIN, statusCode)
	}
	return notOk
}

func WriteHeaderNoContent(writer http.ResponseWriter) {
	writeHeaderContentText(writer, PLAIN, http.StatusNoContent)
}

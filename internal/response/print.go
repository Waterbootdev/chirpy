package response

import (
	"fmt"
	"net/http"
)

func FprintOKResponse(writer http.ResponseWriter, textType string, a ...any) {
	FprintResponse(writer, WriteHeaderContentTextOK(textType), a...)
}

func FprintfOKResponse(writer http.ResponseWriter, textType string, format string, a ...any) {
	FprintfResponse(writer, WriteHeaderContentTextOK(textType), format, a...)
}

func FprintResponse(writer http.ResponseWriter, writeHeader func(http.ResponseWriter), a ...any) {
	writeHeader(writer)
	fmt.Fprint(writer, a...)
}

func FprintfResponse(writer http.ResponseWriter, writeHeader func(http.ResponseWriter), format string, a ...any) {
	writeHeader(writer)
	fmt.Fprintf(writer, format, a...)
}

package response

import (
	"fmt"
	"net/http"
)

func FprintResponse(writer http.ResponseWriter, writeHeader func(http.ResponseWriter), a ...any) {
	writeHeader(writer)
	fmt.Fprint(writer, a...)
}

func FprintfResponse(writer http.ResponseWriter, writeHeader func(http.ResponseWriter), format string, a ...any) {
	writeHeader(writer)
	fmt.Fprintf(writer, format, a...)
}

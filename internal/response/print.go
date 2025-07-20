package response

import (
	"fmt"
	"net/http"
)

func PrintPrintError(printPrintError bool, n int, err error) {
	if printPrintError && err != nil {
		fmt.Printf("Print Error %d: %v\n", n, err)
	}
}
func FprintOKResponse(printPrintError bool, writer http.ResponseWriter, textType string, a ...any) {
	FprintResponse(printPrintError, writer, WriteHeaderContentTextOK(textType), a...)
}

func FprintfOKResponse(printPrintError bool, writer http.ResponseWriter, textType string, format string, a ...any) {
	FprintfResponse(printPrintError, writer, WriteHeaderContentTextOK(textType), format, a...)
}

func FprintResponse(printPrintError bool, writer http.ResponseWriter, writeHeader func(http.ResponseWriter), a ...any) {
	writeHeader(writer)
	n, err := fmt.Fprint(writer, a...)
	PrintPrintError(printPrintError, n, err)
}

func FprintfResponse(printPrintError bool, writer http.ResponseWriter, writeHeader func(http.ResponseWriter), format string, a ...any) {
	writeHeader(writer)
	n, err := fmt.Fprintf(writer, format, a...)
	PrintPrintError(printPrintError, n, err)
}

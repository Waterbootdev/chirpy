package response

import (
	"encoding/json"
	"net/http"
)

func ResponseJsonData(writer http.ResponseWriter, statusCode int, data []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(data)
}

func ResponseJsonMarshal(writer http.ResponseWriter, statusCode int, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		InternalServerErrorResponse(writer, err)
		return
	}
	ResponseJsonData(writer, statusCode, data)
}

type responseError struct {
	Error string `json:"error"`
}

func ErrorResponse(writer http.ResponseWriter, statusCode int, currentResponseError string) {
	ResponseJsonMarshal(writer, statusCode, responseError{Error: currentResponseError})
}

func FromRequestErrorResponse[T any](writer http.ResponseWriter, request *http.Request) (t *T, wasError bool) {
	decoder := json.NewDecoder(request.Body)
	wasError = decoder.Decode(&t) != nil
	if wasError {
		ErrorResponse(writer, http.StatusInternalServerError, "Something went wrong")
	}

	return t, !wasError
}

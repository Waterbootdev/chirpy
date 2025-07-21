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

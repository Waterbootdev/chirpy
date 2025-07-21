package response

import "net/http"

func JsonResponse(writer http.ResponseWriter, statusCode int, data []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(data)
}

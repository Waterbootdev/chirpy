package apiconfig

import "net/http"

func Allways[T any](writer http.ResponseWriter, request *http.Request, t *T) bool {
	return true
}

package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authorization := strings.Split(headers.Get("Authorization"), " ")
	authorizationLength := len(authorization)
	if authorizationLength != 2 || authorization[0] != "Bearer" {
		return "", errors.New("Invalid authorization header")
	}
	return authorization[1], nil
}

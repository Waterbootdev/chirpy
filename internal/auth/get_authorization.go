package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAuthorization(headers http.Header, key string) (string, error) {
	authorization := strings.Split(headers.Get("Authorization"), " ")
	authorizationLength := len(authorization)
	if authorizationLength != 2 || authorization[0] != key {
		return "", errors.New("invalid authorization header")
	}
	return authorization[1], nil
}

func GetBearerToken(headers http.Header) (string, error) {
	return GetAuthorization(headers, "Bearer")
}

func GetApiKey(headers http.Header) (string, error) {
	return GetAuthorization(headers, "ApiKey")
}

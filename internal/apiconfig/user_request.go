package apiconfig

import "net/http"

type userRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r userRequest) IsValidResponse(_ *ApiConfig, _ http.ResponseWriter) (*userRequest, bool) {
	return &r, true
}

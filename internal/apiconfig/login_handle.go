package apiconfig

import (
	"net/http"
)

func (cfg *ApiConfig) loginHandle(request *http.Request, loginRequest *loginRequest) (*userToken, error) {

	token, err := cfg.makeJWT(loginRequest.ExpiresInSeconds, loginRequest.User.ID)

	if err != nil {
		return nil, err
	}

	return fromDatabaseUserToken(loginRequest.User, token), nil
}

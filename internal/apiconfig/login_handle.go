package apiconfig

import (
	"net/http"
)

func (cfg *ApiConfig) loginHandle(request *http.Request, loginRequest *loginRequest) (*user, error) {

	user, err := cfg.queries.GetUser(request.Context(), loginRequest.ID)

	if err != nil {
		return nil, err
	}

	return fromDatabaseUser(&user), nil
}

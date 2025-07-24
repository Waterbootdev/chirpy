package apiconfig

import "github.com/Waterbootdev/chirpy/internal/database"

type accesToken struct {
	Token string `json:"token"`
}

func (cfg *ApiConfig) refreshHandler(refreshToken *database.RefreshToken) (*accesToken, error) {

	token, err := cfg.makeJWT(refreshToken.UserID)

	if err != nil {
		return nil, err
	}

	return &accesToken{Token: token}, nil
}

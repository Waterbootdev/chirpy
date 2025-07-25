package apiconfig

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
)

func (cfg *ApiConfig) RevokeHandler(writer http.ResponseWriter, request *http.Request) {
	token, ok := cfg.getRefreshToken(request)

	if response.UnauthorizedResponse(!ok, writer) {
		return
	}

	err := cfg.queries.RevokeRefreshToken(request.Context(), database.RevokeRefreshTokenParams{
		Token: token.Token,
		RevokedAt: sql.NullTime{
			Valid: true,
			Time:  time.Now(),
		},
	})

	if err == nil {
		response.WriteHeaderNoContent(writer)
	} else {
		response.InternalServerErrorResponse(writer, err)
	}
}

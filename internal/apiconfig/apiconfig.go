package apiconfig

import (
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

type ApiConfig struct {
	fileserverHits atomic.Int32
	queries        *database.Queries
	platform       string
	secret         string
	polkaKey       string
}

func (cfg *ApiConfig) validateJWT(request *http.Request) (uuid.UUID, error) {
	bearerToken, err := auth.GetBearerToken(request.Header)
	if err != nil {
		return uuid.Nil, err
	}
	return auth.ValidateJWT(bearerToken, cfg.secret)
}

func (cfg *ApiConfig) validateJWTResponse(request *http.Request, writer http.ResponseWriter) (uuid.UUID, bool) {
	userID, err := cfg.validateJWT(request)
	ok := err == nil
	if !ok {
		response.ErrorResponse(writer, http.StatusUnauthorized, "Unauthorized")
	}
	return userID, ok
}

func (cfg *ApiConfig) makeJWT(userID uuid.UUID) (string, error) {
	return auth.MakeJWT(userID, cfg.secret, time.Hour)
}

func NewApiConfig() *ApiConfig {
	return &ApiConfig{
		fileserverHits: atomic.Int32{},
		queries:        database.GetDatabaseQueries(),
		platform:       os.Getenv("PLATFORM"),
		secret:         os.Getenv("SECRET"),
		polkaKey:       os.Getenv("POLKA_KEY"),
	}
}

const METRICSFORMAT = `<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`

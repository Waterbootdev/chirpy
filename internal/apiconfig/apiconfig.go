package apiconfig

import (
	"os"
	"sync/atomic"

	"github.com/Waterbootdev/chirpy/internal/database"
)

type ApiConfig struct {
	fileserverHits atomic.Int32
	queries        *database.Queries
	platform       string
}

func NewApiConfig() *ApiConfig {
	return &ApiConfig{
		fileserverHits: atomic.Int32{},
		queries:        database.GetDatabaseQueries(),
		platform:       os.Getenv("PLATFORM"),
	}
}

const METRICSFORMAT = `<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`

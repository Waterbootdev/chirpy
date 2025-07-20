package apiconfig

import (
	"sync/atomic"
)

type ApiConfig struct {
	fileserverHits atomic.Int32
}

const METRICSFORMAT = `<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`

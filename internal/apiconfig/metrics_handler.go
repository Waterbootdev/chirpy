package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
)

const METRICSFORMAT = `<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`

func (cfg *ApiConfig) MetricsHandler(writer http.ResponseWriter, _ *http.Request) {
	response.FprintfOKResponse(PRINTERROR, writer, response.HTML, METRICSFORMAT, cfg.fileserverHits.Load())
}

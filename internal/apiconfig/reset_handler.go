package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
)

const PRINTERROR bool = true

func (cfg *ApiConfig) ResetHandler(writer http.ResponseWriter, request *http.Request) {

	if response.ForbiddenErrorResponse(cfg.platform != "dev", writer) {
		return
	}

	cfg.fileserverHits.Store(0)
	response.FprintOKResponse(PRINTERROR, writer, response.PLAIN, "Hits reset")

	if err := cfg.queries.DeleteUsers(request.Context()); err != nil {
		response.InternalServerErrorResponse(writer, err)
	}
}

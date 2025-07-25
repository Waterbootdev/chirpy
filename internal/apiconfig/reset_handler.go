package apiconfig

import (
	"net/http"

	"github.com/Waterbootdev/chirpy/internal/response"
)

const PRINTERROR bool = true

func (cfg *ApiConfig) ResetHandler(writer http.ResponseWriter, request *http.Request) {

	if response.ErrorResponse(cfg.platform != "dev", writer, http.StatusForbidden, "Forbidden") {
		return
	}

	cfg.fileserverHits.Store(0)
	response.FprintOKResponse(PRINTERROR, writer, response.PLAIN, "Hits reset")

	if err := cfg.queries.DeleteUsers(request.Context()); err != nil {
		response.InternalServerErrorResponse(writer, err)
	}
}

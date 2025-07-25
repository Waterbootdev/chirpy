package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/auth"
	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/Waterbootdev/chirpy/internal/generic_handler"
	"github.com/Waterbootdev/chirpy/internal/response"
	"github.com/google/uuid"
)

type webhook struct {
	Event string `json:"event"`
	Data  struct {
		UserID string `json:"user_id"`
	} `json:"data"`
	User database.User `json:"user"`
}

func writeHeaderContentText(notOk bool, writer http.ResponseWriter, statusCode int) bool {

	if notOk {
		response.WriteHeaderContentText(writer, response.PLAIN, statusCode)
	}
	return notOk
}

func (cfg *ApiConfig) validatePolkaKey(request *http.Request) bool {
	apiKey, err := auth.GetApiKey(request.Header)

	if err != nil {
		return false
	}

	return apiKey == cfg.polkaKey
}

func (cfg *ApiConfig) webhookValidator(writer http.ResponseWriter, request *http.Request, webhook *webhook) bool {

	if writeHeaderContentText(!cfg.validatePolkaKey(request), writer, http.StatusUnauthorized) {
		return false
	}

	if writeHeaderContentText(webhook.Event != "user.upgraded", writer, http.StatusNoContent) {
		return false
	}

	user, err := cfg.queries.GetUser(request.Context(), uuid.MustParse(webhook.Data.UserID))

	if writeHeaderContentText(err != nil, writer, http.StatusNotFound) {
		return false
	}

	webhook.User = user

	return true
}

func (cfg *ApiConfig) webhookHandle(request *http.Request, webhook *webhook) error {
	return cfg.queries.UpdateIsChirpyRed(request.Context(), database.UpdateIsChirpyRedParams{
		ID:          webhook.User.ID,
		IsChirpyRed: true,
		UpdatedAt:   time.Now(),
	})
}

func (cfg *ApiConfig) WebhookHandler(writer http.ResponseWriter, request *http.Request) {
	generic_handler.NoContentBodyHandler(writer, request, cfg.webhookHandle, cfg.webhookValidator)
}

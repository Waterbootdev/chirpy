package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
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

func (cfg *ApiConfig) webhookValidator(writer http.ResponseWriter, request *http.Request, webhook *webhook) bool {

	if response.WriteHeaderContentText(!cfg.validatePolkaKey(request), writer, http.StatusNoContent) {
		return false
	}

	if response.WriteHeaderContentText(webhook.Event != "user.upgraded", writer, http.StatusNoContent) {
		return false
	}

	user, err := cfg.queries.GetUser(request.Context(), uuid.MustParse(webhook.Data.UserID))

	if response.WriteHeaderContentText(err != nil, writer, http.StatusNotFound) {
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
	response.NoContentBodyHandler(writer, request, cfg.webhookHandle, cfg.webhookValidator)
}

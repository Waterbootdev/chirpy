package chirp

import (
	"encoding/json"
	"net/http"
)

func ChirpFromRequest(request *http.Request) (chirp Chirp, err error) {
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&chirp)
	return chirp, err
}

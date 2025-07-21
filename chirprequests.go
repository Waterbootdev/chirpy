package main

import (
	"encoding/json"
	"net/http"
)

func getChirp(request *http.Request) (chirp chirp, err error) {
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&chirp)
	return chirp, err
}

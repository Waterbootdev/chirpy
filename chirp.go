package main

const MAX_CHIRP_LENGTH int = 140

type chirp struct {
	Body string `json:"body"`
}

type chirpError struct {
	Error string `json:"error"`
}

type cleanedChirp struct {
	CleanedBody string `json:"cleaned_body"`
}

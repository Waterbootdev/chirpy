package chirp

type Chirp struct {
	Body string `json:"body"`
}

type chirpError struct {
	Error string `json:"error"`
}

type cleanedChirp struct {
	CleanedBody string `json:"cleaned_body"`
}

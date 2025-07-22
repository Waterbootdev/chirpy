package chirp

type Chirp struct {
	Body string `json:"body"`
}

type cleanedChirp struct {
	CleanedBody string `json:"cleaned_body"`
}

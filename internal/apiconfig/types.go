package apiconfig

import (
	"net/http"
	"time"

	"github.com/Waterbootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

type userRequest struct {
	Email string `json:"email"`
}

func (e userRequest) IsValidResponse(writer http.ResponseWriter) bool {
	return true
}

type user struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func fromDatabaseUser(dbUser *database.User) *user {
	return &user{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
	}
}

type chirpRequest struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

type chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func fromDatabaseChirp(dbChirp *database.Chirp) *chirp {
	return &chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	}
}

func fromDatabaseChirps(dbChirp []database.Chirp) []chirp {
	chirps := make([]chirp, len(dbChirp))
	for i, dbChirp := range dbChirp {
		chirps[i] = *fromDatabaseChirp(&dbChirp)
	}
	return chirps
}

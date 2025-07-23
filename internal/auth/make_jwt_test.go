package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestValidateJWT(t *testing.T) {
	userId := uuid.New()
	tokenSecret := "secret"
	token, err := MakeJWT(userId, tokenSecret, time.Hour)

	if err != nil {
		t.Fatal(err)
	}

	testID, err := ValidateJWT(token, tokenSecret)

	if err != nil {
		t.Fatal(err)
	}

	if testID != userId {
		t.Errorf("Expected %s, got %s", userId, testID)
	}
}

package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcryptHashValidation(t *testing.T) {
	hash := "$2a$10$mjPchIHjdc5d2hjxXohZEOyoZgUWGnAii453mgABT0sygBX1Nq4yO"
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte("password123"))
	if err != nil {
		t.Fatalf("Password verification failed: %v", err)
	}
}

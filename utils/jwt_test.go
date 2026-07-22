package utils

import (
	"testing"
	"time"

	"github.com/adsomniac/my-app/models"
)

func TestJWTService_GenerateAndValidateToken(t *testing.T) {
	secret := "test-super-secret-key-32-characters-long"
	jwtSvc := NewJWTService(secret, 1)

	user := &models.User{
		ID:    101,
		Email: "admin@test.com",
		Role:  "admin",
	}

	tokenStr, err := jwtSvc.GenerateToken(user)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	if tokenStr == "" {
		t.Fatal("Expected non-empty token string")
	}

	claims, err := jwtSvc.ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}

	if claims.UserID != user.ID {
		t.Errorf("Expected UserID %d, got %d", user.ID, claims.UserID)
	}

	if claims.Email != user.Email {
		t.Errorf("Expected Email %s, got %s", user.Email, claims.Email)
	}

	if claims.Role != user.Role {
		t.Errorf("Expected Role %s, got %s", user.Role, claims.Role)
	}
}

func TestJWTService_InvalidSecret(t *testing.T) {
	secret1 := "test-super-secret-key-32-characters-1"
	secret2 := "test-super-secret-key-32-characters-2"

	svc1 := NewJWTService(secret1, 1)
	svc2 := NewJWTService(secret2, 1)

	user := &models.User{ID: 1, Email: "user@test.com", Role: "user"}

	tokenStr, _ := svc1.GenerateToken(user)
	_, err := svc2.ValidateToken(tokenStr)
	if err == nil {
		t.Fatal("Expected validation error for mismatched secret key")
	}
}

func TestJWTService_ExpiredToken(t *testing.T) {
	secret := "test-super-secret-key-32-characters-long"
	// Expire in -1 hour (already expired)
	jwtSvc := NewJWTService(secret, -1)

	user := &models.User{ID: 2, Email: "user@test.com", Role: "user"}

	tokenStr, err := jwtSvc.GenerateToken(user)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Small pause to ensure time has passed if needed
	time.Sleep(10 * time.Millisecond)

	_, err = jwtSvc.ValidateToken(tokenStr)
	if err == nil {
		t.Fatal("Expected validation error for expired token")
	}
}

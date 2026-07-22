package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/adsomniac/my-app/models"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`

	jwt.RegisteredClaims
}

type JWTService struct {
	Secret     string
	ExpireHour int
}

func NewJWTService(secret string, expireHour int) *JWTService {
	return &JWTService{
		Secret:     secret,
		ExpireHour: expireHour,
	}
}

func (j *JWTService) GenerateToken(user *models.User) (string, error) {
	if user == nil {
		return "", errors.New("user cannot be nil")
	}

	now := time.Now()
	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(j.ExpireHour) * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func (j *JWTService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	if claims.UserID <= 0 {
		return nil, errors.New("token missing user ID")
	}

	if claims.Role == "" {
		return nil, errors.New("token missing role")
	}

	return claims, nil
}

package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/adsomniac/my-app/utils"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	EmailKey  contextKey = "email"
	RoleKey   contextKey = "role"
)

type AuthMiddlewareProvider struct {
	JWTService *utils.JWTService
}

func NewAuthMiddlewareProvider(jwtSvc *utils.JWTService) *AuthMiddlewareProvider {
	return &AuthMiddlewareProvider{JWTService: jwtSvc}
}

func (m *AuthMiddlewareProvider) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeJSONError(w, http.StatusUnauthorized, "Token otentikasi tidak ditemukan")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			writeJSONError(w, http.StatusUnauthorized, "Format header Authorization harus Bearer <token>")
			return
		}

		tokenString := parts[1]
		claims, err := m.JWTService.ValidateToken(tokenString)
		if err != nil {
			writeJSONError(w, http.StatusUnauthorized, "Token tidak valid atau telah kedaluwarsa")
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, EmailKey, claims.Email)
		ctx = context.WithValue(ctx, RoleKey, claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

var roleHierarchy = map[string]int{
	"user":        1,
	"admin":       2,
	"super_admin": 3,
}

func RequireMinimumRole(minRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole, ok := r.Context().Value(RoleKey).(string)
			if !ok || userRole == "" {
				writeJSONError(w, http.StatusUnauthorized, "Informasi role tidak ditemukan")
				return
			}

			userLevel, userOk := roleHierarchy[userRole]
			minLevel, minOk := roleHierarchy[minRole]

			if !userOk || !minOk || userLevel < minLevel {
				writeJSONError(w, http.StatusForbidden, "Akses ditolak: Peran tidak mencukupi")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func writeJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"message": message,
	})
}

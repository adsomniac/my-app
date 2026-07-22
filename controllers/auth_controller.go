package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/adsomniac/my-app/models"
	"github.com/adsomniac/my-app/services"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONResponse(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"message": "Metode HTTP tidak diizinkan",
		})
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Format request JSON tidak valid",
		})
		return
	}

	if req.Email == "" || req.Password == "" {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Email dan password wajib diisi",
		})
		return
	}

	result, err := c.AuthService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) || errors.Is(err, services.ErrUserInactive) {
			writeJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
				"success": false,
				"message": "Email atau password salah",
			})
			return
		}

		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": "Terjadi kesalahan pada server",
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Login berhasil",
		"data": map[string]interface{}{
			"token": result.Token,
			"user": map[string]interface{}{
				"id":        result.User.ID,
				"email":     result.User.Email,
				"role":      result.User.Role,
				"is_active": result.User.IsActive,
			},
		},
	})
}

func (c *AuthController) Me(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONResponse(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"message": "Metode HTTP tidak diizinkan",
		})
		return
	}

	userID, _ := r.Context().Value(UserIDKey).(int64)
	email, _ := r.Context().Value(EmailKey).(string)
	role, _ := r.Context().Value(RoleKey).(string)

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"id":    userID,
			"email": email,
			"role":  role,
		},
	})
}

func (c *AuthController) Profile(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(UserIDKey).(int64)
	email, _ := r.Context().Value(EmailKey).(string)
	role, _ := r.Context().Value(RoleKey).(string)

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Profil pengguna berhasil diambil",
		"data": map[string]interface{}{
			"id":    userID,
			"email": email,
			"role":  role,
		},
	})
}

func (c *AuthController) AdminTest(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Akses area admin berhasil disetujui",
	})
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

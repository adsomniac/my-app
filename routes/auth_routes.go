package routes

import (
	"net/http"

	"github.com/adsomniac/my-app/controllers"
)

func RegisterAuthRoutes(mux *http.ServeMux, authController *controllers.AuthController, authMiddleware *controllers.AuthMiddlewareProvider) {
	// Public routes
	mux.HandleFunc("POST /api/auth/login", authController.Login)

	// Protected routes (requires valid JWT)
	meHandler := authMiddleware.AuthMiddleware(http.HandlerFunc(authController.Me))
	mux.Handle("GET /api/auth/me", meHandler)

	profileHandler := authMiddleware.AuthMiddleware(http.HandlerFunc(authController.Profile))
	mux.Handle("GET /api/profile", profileHandler)

	// Protected admin routes (requires valid JWT & minimum role "admin")
	adminTestHandler := authMiddleware.AuthMiddleware(
		controllers.RequireMinimumRole("admin")(http.HandlerFunc(authController.AdminTest)),
	)
	mux.Handle("GET /api/admin/test", adminTestHandler)
}

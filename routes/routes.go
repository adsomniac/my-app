package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/adsomniac/my-app/controllers"
)

// SetupRoutes registers all application API routes and static SPA fallback routes
func SetupRoutes(mux *http.ServeMux, authController *controllers.AuthController, authMiddleware *controllers.AuthMiddlewareProvider) {
	// Register Health API Routes
	RegisterHealthRoutes(mux)

	// Register Auth API Routes
	if authController != nil && authMiddleware != nil {
		RegisterAuthRoutes(mux, authController, authMiddleware)
	}

	// Static Files & SPA fallback from frontend/dist
	staticDir := "frontend/dist"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Printf("Warning: '%s' directory does not exist yet. Run 'npm run build' in frontend directory.", staticDir)
		_ = os.MkdirAll(staticDir, 0755)
	}

	spa := controllers.SPAHandler{StaticPath: staticDir, IndexPath: "index.html"}
	mux.Handle("/", spa)
}

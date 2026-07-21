package routes

import (
	"net/http"

	"github.com/adsomniac/my-app/controllers"
)

// RegisterHealthRoutes registers all health-check endpoints
func RegisterHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/health", controllers.HealthHandler)
}

package main

import (
	"log"
	"net/http"

	"github.com/adsomniac/my-app/config"
	"github.com/adsomniac/my-app/database"
	"github.com/adsomniac/my-app/routes"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize Database (optional on startup, logs warning if DATABASE_URL is not set)
	if cfg.DatabaseURL != "" {
		if _, err := database.InitDB(cfg.DatabaseURL); err != nil {
			log.Printf("Database connection warning: %v\n", err)
		}
	}

	mux := http.NewServeMux()

	// Setup modular MVC routes
	routes.SetupRoutes(mux)

	log.Printf("Server starting on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

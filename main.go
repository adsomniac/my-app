package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/adsomniac/my-app/config"
	"github.com/adsomniac/my-app/controllers"
	"github.com/adsomniac/my-app/database"
	"github.com/adsomniac/my-app/repositories"
	"github.com/adsomniac/my-app/routes"
	"github.com/adsomniac/my-app/services"
	"github.com/adsomniac/my-app/utils"
)

func main() {
	cfg := config.LoadConfig()

	var db *sql.DB
	// Initialize Database (optional on startup, logs warning if DATABASE_URL is not set)
	if cfg.DatabaseURL != "" {
		var err error
		db, err = database.InitDB(cfg.DatabaseURL)
		if err != nil {
			log.Printf("Database connection warning: %v\n", err)
		}
	}

	// Initialize Auth Dependencies
	userRepo := repositories.NewUserRepository(db)
	jwtSvc := utils.NewJWTService(cfg.JWTSecret, cfg.JWTExpireHours)
	authSvc := services.NewAuthService(userRepo, jwtSvc)
	authController := controllers.NewAuthController(authSvc)
	authMiddleware := controllers.NewAuthMiddlewareProvider(jwtSvc)

	mux := http.NewServeMux()

	// Setup modular MVC routes
	routes.SetupRoutes(mux, authController, authMiddleware)

	log.Printf("Server starting on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

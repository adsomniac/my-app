package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/adsomniac/my-app/database"
	"github.com/adsomniac/my-app/models"
)

// HealthHandler responds with the status of the Go server and database connection
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbStatus := "disconnected"
	if database.DB != nil {
		if err := database.DB.Ping(); err == nil {
			dbStatus = "connected"
		}
	}

	response := models.StatusResponse{
		Status:   "ok",
		Message:  "Golang Monolith Server (Modular MVC) is active",
		Database: dbStatus,
	}

	json.NewEncoder(w).Encode(response)
}

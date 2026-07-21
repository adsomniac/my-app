package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adsomniac/my-app/db"
)

type StatusResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Database string `json:"database"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbStatus := "disconnected"
	if db.DB != nil {
		if err := db.DB.Ping(); err == nil {
			dbStatus = "connected"
		}
	}

	response := StatusResponse{
		Status:   "ok",
		Message:  "Golang Monolith Server is active",
		Database: dbStatus,
	}

	json.NewEncoder(w).Encode(response)
}

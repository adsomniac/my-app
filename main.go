package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/adsomniac/my-app/config"
	"github.com/adsomniac/my-app/db"
	"github.com/adsomniac/my-app/handlers"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Don't intercept API routes
	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}

	path := filepath.Join(h.staticPath, r.URL.Path)
	fi, err := os.Stat(path)

	// If file doesn't exist or is a directory, serve index.html for SPA routing
	if os.IsNotExist(err) || fi.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	cfg := config.LoadConfig()

	// Initialize Database (optional on startup, logs warning if DATABASE_URL is not set)
	if cfg.DatabaseURL != "" {
		if _, err := db.InitDB(cfg.DatabaseURL); err != nil {
			log.Printf("Database connection warning: %v\n", err)
		}
	}

	mux := http.NewServeMux()

	// API Routes
	mux.HandleFunc("/api/health", handlers.HealthHandler)

	// Static Files & SPA fallback from frontend/dist
	staticDir := "frontend/dist"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Printf("Warning: '%s' directory does not exist yet. Run 'npm run build' in frontend directory.", staticDir)
		_ = os.MkdirAll(staticDir, 0755)
	}

	spa := spaHandler{staticPath: staticDir, indexPath: "index.html"}
	mux.Handle("/", spa)

	log.Printf("Server starting on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

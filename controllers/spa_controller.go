package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// SPAHandler implements http.Handler for serving Single Page Applications
type SPAHandler struct {
	StaticPath string
	IndexPath  string
}

func (h SPAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Don't intercept API routes
	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}

	path := filepath.Join(h.StaticPath, r.URL.Path)
	fi, err := os.Stat(path)

	// If file doesn't exist or is a directory, serve index.html for SPA routing
	if os.IsNotExist(err) || fi.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}

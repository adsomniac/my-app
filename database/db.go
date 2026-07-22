package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

// DB holds the global database instance
var DB *sql.DB

// InitDB initializes the connection to the PostgreSQL database and runs pending migrations
func InitDB(databaseURL string) (*sql.DB, error) {
	if databaseURL == "" {
		log.Println("Warning: DATABASE_URL is empty. Running without database connection.")
		return nil, nil
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database.")

	if err := RunMigrations(db); err != nil {
		log.Printf("Warning: Failed to execute database migrations: %v\n", err)
	}

	DB = db
	return db, nil
}

// RunMigrations automatically executes all .sql migration files in database/migrations
func RunMigrations(db *sql.DB) error {
	migrationsDir := filepath.Join("database", "migrations")
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	var sqlFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}

	sort.Strings(sqlFiles)

	for _, fileName := range sqlFiles {
		filePath := filepath.Join(migrationsDir, fileName)
		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", fileName, err)
		}

		query := string(content)
		if strings.TrimSpace(query) == "" {
			continue
		}

		_, err = db.Exec(query)
		if err != nil {
			// If error is table already exists or index already exists, log it gracefully
			if strings.Contains(err.Error(), "already exists") {
				log.Printf("Migration %s already applied (skipped).\n", fileName)
				continue
			}
			return fmt.Errorf("failed to execute migration %s: %w", fileName, err)
		}

		log.Printf("Successfully executed database migration: %s\n", fileName)
	}

	return nil
}

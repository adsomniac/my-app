package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DB holds the global database instance
var DB *sql.DB

// InitDB initializes the connection to the PostgreSQL database
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
	DB = db
	return db, nil
}

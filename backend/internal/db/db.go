package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Get database URL from environment variable
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// Fallback to localhost for development
		connStr = "host=localhost port=5432 user=postgres password=param1978 dbname=connect4 sslmode=disable"
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	log.Println("Connected to PostgreSQL")
}

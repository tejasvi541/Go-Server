package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Connect initializes the database connection and returns the *sql.DB instance
func Connect() {
	// Open a connection to the PostgreSQL database
	var err error
	DB, err = sql.Open("postgres", "user=postgres dbname=events password=12345678 sslmode=disable")
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	// Set maximum number of idle connections
	DB.SetMaxIdleConns(10)

	// Set maximum number of open connections
	DB.SetMaxOpenConns(100)

	// Verify the connection is valid
	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}
}

// CreateTables creates necessary tables in the database
func CreateTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			date_time TIMESTAMP NOT NULL,
			user_id INT NOT NULL
		);
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}
}

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB


// Connect initializes the database connection and returns the *sql.DB instance
func Connect() {
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "12345678")
	port := getEnv("DB_PORT", "5432")
	dbname := getEnv("DB_NAME", "events")

	// Create connection string
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s port=%s sslmode=disable",
		user,
		dbname,
		password,
		port)
	
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	// Set maximum number of idle connections
	DB.SetMaxIdleConns(10)

	// Set maximum number of open connections
	DB.SetMaxOpenConns(100)
	fmt.Println(connStr, "connStr")
	// Verify the connection is valid
	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}
}

// CreateTables creates necessary tables in the database
func CreateTables() {

	// Create the users table
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		log.Fatal("Error creating tables:", err)
		panic(err)
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			date_time TIMESTAMP NOT NULL,
			user_id INT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users (id)
		);
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal("Error creating tables:", err)
		panic(err)
	}
	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id SERIAL PRIMARY KEY,
			event_id INT NOT NULL,
			user_id INT NOT NULL,
			FOREIGN KEY (event_id) REFERENCES events (id),
			FOREIGN KEY (user_id) REFERENCES users (id)
		);
	`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		log.Fatal("Error creating tables:", err)
		panic(err)
	}
}
// getEnv retrieves the value of an environment variable or returns a default value if the variable is not set
func getEnv(key, defaultValue string) string {
	godotenv.Load(".env")
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
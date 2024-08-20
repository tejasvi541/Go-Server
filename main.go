package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tejasvi541/Go-Server/src/db"
	"github.com/tejasvi541/Go-Server/src/routes"
)

func main() {
	// Connect to the database
	db.Connect()

	// Create tables if they do not exist
	db.CreateTables()

	// Ensure the database connection is closed when the application exits
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()
	// Create a new Gin router
	server := gin.Default()

	// Register routes
	routes.RegisterRoutes(server)
	
	// Start the server on port 8088
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}


package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tejasvi541/Go-Server/src/db"
	"github.com/tejasvi541/Go-Server/src/models"
)

func main() {
	// Connect to the database
	db.Connect()
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	// Create a new Gin router
	server := gin.Default()

	// Route to get all events
	server.GET("/events", func(c *gin.Context) {
		events, err := models.GetAllEvents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, events)
	})

	// Route to create a new event
	server.POST("/events", func(c *gin.Context) {
		var event models.Event
		if err := c.ShouldBindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Set a default UserID for demonstration purposes
		event.UserID = 1

		// Save the event
		if err := event.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, event)
	})

	// Start the server on port 8088
	if err := server.Run(":8088"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

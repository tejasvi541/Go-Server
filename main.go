package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tejasvi541/Go-Server/src/db"
	"github.com/tejasvi541/Go-Server/src/models"
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

	// Route to get all events
	server.GET("/events", getEvents)

	// Route to get sigle event
	server.GET("/events/:id", getEventById)

	// Route to create a new event
	server.POST("/events", createEvent)

	// Start the server on port 8088
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEvents(c *gin.Context) {
		events, err := models.GetAllEvents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, events)
}

func getEventById(c *gin.Context) {
		eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}

		event, err := models.GetEventById(eventId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, event)

}

func createEvent(c *gin.Context) {
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
	}
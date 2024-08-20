package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Route to get all events
	server.GET("/events", getEvents)

	// Route to get sigle event
	server.GET("/events/:id", getEventById)

	// Route to create a new event
	server.POST("/events", createEvent)

	// Route to update an event
	server.PUT("/events/:id", updateEvent)

	// Route to delete an event
	server.DELETE("/events/:id", deleteEvent)

	// Signup route
	server.POST("/signup", signup)
}
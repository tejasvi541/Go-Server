package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tejasvi541/Go-Server/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	// Route to get all events
	server.GET("/events", getEvents)

	// Route to get sigle event
	server.GET("/events/:id", getEventById)

	// Route to create a new event
	server.POST("/events", middleware.Authenticate, createEvent)

	// Route to update an event
	server.PUT("/events/:id", middleware.Authenticate, updateEvent)

	// Route to delete an event
	server.DELETE("/events/:id", middleware.Authenticate,  deleteEvent)

	// Signup route
	server.POST("/signup", signup)

	// Login route
	server.POST("/login", login)

	// Register Event route
	server.POST("/register", middleware.Authenticate, registerEvent)

	// Unregister Event route
	server.POST("/unregister", middleware.Authenticate, unregisterEvent)
}
package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tejasvi541/Go-Server/src/models"
)


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
		event.UserID = c.MustGet("userID").(int64)

		// Save the event
		if err := event.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, event)
}

func updateEvent(c *gin.Context) {
		eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}

		event, err := models.GetEventById(eventId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}

		userId := c.MustGet("userID").(int64)

		if event.UserID != userId {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this event"})
			return
		}

		var updatedEvent models.Event

		if err := c.ShouldBindJSON(&updatedEvent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedEvent.ID = eventId

		if err := updatedEvent.Update(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Event updated successfully"})

		
}

func deleteEvent(c *gin.Context) {
		eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}

		event, err := models.GetEventById(eventId)
		userId := c.MustGet("userID").(int64)

		if event.UserID != userId {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this event"})
			return
		}

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
		
		err = event.DeleteEvent()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
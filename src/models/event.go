package models

import (
	"time"

	"github.com/tejasvi541/Go-Server/src/db"
)

// Event represents an event in the system
type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime"`
	UserID      int       `json:"userId"`
}

// Save inserts a new event into the database
func (e *Event) Save() error {
	if e.DateTime.IsZero() {
		e.DateTime = time.Now() // Set DateTime to current time if it's not set
	}

	query := `
		INSERT INTO events (name, description, location, date_time, user_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&e.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetAllEvents retrieves all events from the database
func GetAllEvents() ([]Event, error) {
	var events []Event
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

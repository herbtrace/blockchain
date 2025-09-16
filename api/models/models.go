package models

import "time"

// EventType represents a generic supply chain event
type EventType struct {
	Type        string `json:"type"`        // e.g., "CollectionEvent"
	Description string `json:"description"` // optional notes
}

// QrCodeData represents the incoming JSON payload
type QrCodeData struct {
	FromID    string    `json:"from_id" binding:"required"`
	ToID      string    `json:"to_id" binding:"required"`
	Crops     string    `json:"crops" binding:"required"`
	FromRole  string    `json:"from_role" binding:"required"`
	ToRole    string    `json:"to_role" binding:"required"`
	StartTime time.Time `json:"start_time"` // optional
	Event     EventType `json:"event" binding:"required"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

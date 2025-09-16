package main

import (
	"qrcode-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health check endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running!"})
	})

	// Main POST endpoint
	r.POST("/collection-event", handlers.CreateCollectionEvent)

	// GET endpoint to query all blockchain events
	r.GET("/collection-events", handlers.GetAllBlockchainEvents)

	// Start server on port 8080
	r.Run(":8080")
}

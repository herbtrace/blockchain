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
	r.POST("/qrcode", handlers.HandleQrCodeData)

	// Start server on port 8080
	r.Run(":8080")
}

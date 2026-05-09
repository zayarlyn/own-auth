package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server on port 8080 (default)
	router.Run()
}

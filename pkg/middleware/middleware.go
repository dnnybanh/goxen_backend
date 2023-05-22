package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Calculate response time
		latency := time.Since(start)

		// Log request details
		log.Printf("[REQUEST] %s | %s | %s",
			c.Request.Method,
			c.Request.URL.Path,
			latency)
	}
}

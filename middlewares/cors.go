package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(c *gin.Context) {
	var allowedOriginStr string
	env := os.Getenv("GO_ENV")

	if env == "production" {
		allowedOriginStr = "https://use-vendex.netlify.app"
	} else {
		allowedOriginStr = "http://localhost:3000"
	}

	origin := c.Request.Header.Get("Origin")

	if origin == allowedOriginStr {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Vary", "Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}

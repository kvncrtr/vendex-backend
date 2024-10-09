package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(c *gin.Context) {
	allowedOriginStr := os.Getenv("CORS_URL")
	c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOriginStr)
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClassAorBMiddleware(c *gin.Context) {
	class, exists := c.Get("class")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Class not found in context"})
		c.Abort()
		return
	}

	classStr, ok := class.(string)
	if !ok || (classStr != "A" && classStr != "B") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()
		return
	}

	c.Next()
}

func ClassAMiddleware(c *gin.Context) {
	class, exists := c.Get("class")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Class not found in context"})
		c.Abort()
		return
	}

	classStr, ok := class.(string)
	if !ok || (classStr != "A") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden, class B and C not permitted!"})
		c.Abort()
		return
	}

	c.Next()
}
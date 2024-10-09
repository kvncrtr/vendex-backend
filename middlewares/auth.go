package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvncrtr/vendex/utils"
)

func AuthenticateEmployee(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized.", "error": "token is empty"})
		return
	}

	employee_id, class, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	c.Set("employee_id", employee_id)
	c.Set("class", class)
	c.Next()
}

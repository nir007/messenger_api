package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleOptions handles OPTEIONS http requests
func HandleOptions(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{"access": true})
		c.Abort()
		return
	}

	c.Next()
}

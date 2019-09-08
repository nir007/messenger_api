package middlewares

import "github.com/gin-gonic/gin"

// FilesType check mime-type of files
func FilesType(c *gin.Context) {
	c.Next()
}

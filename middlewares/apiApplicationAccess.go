package middlewares

import (
	"messenger/drepository"
	"messenger/dto"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

// ApiApplicationAccess check access to api functions
func ApiApplicationAccess(c *gin.Context) {
	secretKey := c.Request.Header["Secret-Key"]
	applicationID := c.Request.Header["Application-Id"]

	if len(secretKey) != 1 || len(applicationID) != 1 {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "undefined Secret-Key or Application-ID headers"},
		)
		c.Abort()
		return
	}

	objID, err := primitive.ObjectIDFromHex(applicationID[0])

	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "invalid application ID"},
		)
		c.Abort()
		return
	}

	findApplication := &dto.FindApplications{
		ID:      objID,
		Secret:  secretKey[0],
		Domains: []string{c.Request.Host},
	}

	app := &drepository.Application{}
	err = app.FindOne(findApplication)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"error": "application not fount"},
		)
		c.Abort()
		return
	}

	c.Next()
}

package middlewares

import (
	"messenger/drepository"
	"messenger/dto"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
)

// SiteApplicationAccess check access to site api functions
func SiteApplicationAccess(c *gin.Context) {
	appObjID, err := primitive.ObjectIDFromHex(c.Param("appID"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid application ID"},
		)
		c.Abort()
		return
	}

	managerID, exists := c.Get("managerID")
	if !exists {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "user is not authorized"},
		)
		c.Abort()
		return
	}

	findApplication := &dto.FindApplications{
		ID:      appObjID,
		Managers:[]string{managerID.(string)},
	}

	app := &drepository.Application{}
	err = app.FindOne(findApplication)

	if err != nil {
		c.JSON(
			http.StatusForbidden,
			gin.H{"error": "application access deny"},
		)
		c.Abort()
		return
	}

	c.Next()
}

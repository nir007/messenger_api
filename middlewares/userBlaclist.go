package middlewares

import (
	"messenger/application"
	"messenger/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Blacklist check access to send a message
func Blacklist(c *gin.Context) {
	message := &application.Message{}
	context := c.Copy()

	err := context.ShouldBind(message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(message.UID2) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "undefined user id"})
		c.Abort()
		return
	}

	findUser := &dto.FindUsers{
		UID:       message.UID2,
		BlackList: []string{message.UID1},
	}

	user := application.User{ApplicationID: message.ApplicationID}
	if err := user.FindOne(findUser); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are in blacklist"})
		c.Abort()
		return
	}

	c.Next()
}

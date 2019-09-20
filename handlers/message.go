package handlers

import (
	"messenger/drepository"
	"messenger/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMessage creates new message
func CreateMessage(c *gin.Context) {
	message := &drepository.Message{}

	err := c.ShouldBind(message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if appID, ok := c.Request.Header["Application-Id"]; ok {
		message.ApplicationID = appID[0]
	}

	_, err = message.Insert()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": message})
}

// FindOneMessage searc message
func FindOneMessage(c *gin.Context) {}

// FindAllMessages search messages
func FindAllMessages(c *gin.Context) {
	find := &dto.FindMessages{}

	err := c.ShouldBind(find)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if dialogID := c.Param("dialogid"); len(dialogID) > 0 {
		find.DialogID = dialogID
		find.ApplicationID = c.Request.Header["Application-Id"][0]
	}

	message := &drepository.Message{ApplicationID: find.ApplicationID}
	messages, total, err := message.Find(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": messages, "total": total})
}

// UpdateMessage changes message
func UpdateMessage(c *gin.Context) {}

// DeleteMessage deletes message
func DeleteMessage(c *gin.Context) {}

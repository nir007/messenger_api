package handlers

import (
	"messenger/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMessage creates new message
func CreateMessage(c *gin.Context) {
	message := &application.Message{}

	err := c.ShouldBind(message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	message.ApplicationID = c.Request.Header["Application-Id"][0]
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
func FindAllMessages(c *gin.Context) {}

// UpdateMessage changes message
func UpdateMessage(c *gin.Context) {}

// DeleteMessage deletes message
func DeleteMessage(c *gin.Context) {}

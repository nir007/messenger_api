package handlers

import (
	"messenger/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateManager creates new manager
func CreateManager(c *gin.Context) {
	manager := &application.Manager{}
	err := c.Bind(manager)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "binding params: " + err.Error(),
		})
		return
	}

	_, err = manager.Insert()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": manager})
}

// FindOneManager search manager
func FindOneManager(c *gin.Context) {}

// FindAllManagers search managers
func FindAllManagers(c *gin.Context) {}

// UpdateManager changes manager
func UpdateManager(c *gin.Context) {}

// DeleteManager removes manager
func DeleteManager(c *gin.Context) {}

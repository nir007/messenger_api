package handlers

import (
	"fmt"
	"messenger/application"
	"messenger/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	manager.Password = ""
	c.JSON(http.StatusOK, gin.H{"result": manager})
}

// FindOneManager search manager
func FindOneManager(c *gin.Context) {
	id, _ := c.Get("managerID")
	objectID, _ := primitive.ObjectIDFromHex(id.(string))

	fmt.Println("id: ", id)

	find := &dto.FindManagers{ID: objectID}

	err := c.ShouldBind(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	manager := &application.Manager{}

	err = manager.FindOne(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": manager})
}

// FindAllManagers search managers
func FindAllManagers(c *gin.Context) {}

// UpdateManager changes manager
func UpdateManager(c *gin.Context) {}

// DeleteManager removes manager
func DeleteManager(c *gin.Context) {}

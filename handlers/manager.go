package handlers

import (
	"messenger/drepository"
	"messenger/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateManager creates new manager
func CreateManager(c *gin.Context) {
	manager := &drepository.Manager{}
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

	find := &dto.FindManagers{ID: objectID}
	err := c.ShouldBind(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	manager := &drepository.Manager{}
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
func UpdateManager(c *gin.Context) {
	updateManager := &dto.UpdateManager{}
	err := c.ShouldBind(updateManager)

	id, _ := c.Get("managerID")
	objectID, _ := primitive.ObjectIDFromHex(id.(string))

	findMenager := &dto.FindManagers{ID: objectID}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	manager := &drepository.Manager{}
	_, err = manager.Update(findMenager, updateManager)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	manager.Password = ""

	c.JSON(http.StatusOK, gin.H{"result": manager})

}

//UpdateManagerAvatar updates manager avatar
func UpdateManagerAvatar(c *gin.Context) {

}

// DeleteManager removes manager
func DeleteManager(c *gin.Context) {}

package handlers

import (
	"messenger/drepository"
	"messenger/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser creates new user
func CreateUser(c *gin.Context) {
	user := &drepository.User{}
	err := c.Bind(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "binding params: " + err.Error(),
		})
		return
	}

	_, err = user.Insert()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": user})
}

// FindOneUser search one user
func FindOneUser(c *gin.Context) {
	userID := c.Param("id")
	appID := c.Param("appID")
	objectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": InvalidIdentifier})
		c.Abort()
		return
	}

	find := &dto.FindUsers{ID: objectID, ApplicationID: appID}
	user := &drepository.User{ApplicationID: appID}
	err = user.FindOne(find)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": FindDbError})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": user})
}

// FindAllUsers find users
func FindAllUsers(c *gin.Context) {
	find := &dto.FindUsers{}
	err := c.ShouldBind(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(find.ApplicationID) == 0 {
		find.ApplicationID = c.Param("appID")
	}

	user := &drepository.User{ApplicationID: find.ApplicationID}
	users, total, err := user.Find(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": users, "total": total})
}

// UpdateUser changes user
func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	appID := c.Param("appID")

	userObjectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": InvalidIdentifier})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": InvalidIdentifier})
		c.Abort()
		return
	}

	findUser := &dto.FindUsers{ID: userObjectID, ApplicationID: appID}
	updateUser := &dto.UpdateUser{}
	err = c.ShouldBind(updateUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": BindingError})
		c.Abort()
		return
	}

	user := &drepository.User{ApplicationID: appID}
	_, err = user.Update(findUser, updateUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": UpdateDbError})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": user})
}

// DeleteUser removes user
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	appID := c.Param("appID")

	objectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": InvalidIdentifier})
		c.Abort()
		return
	}

	user := &drepository.User{ID: objectID, ApplicationID: appID}
	_, err = user.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": DeleteDbError})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": true})
}

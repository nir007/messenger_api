package handlers

import (
	"messenger/application"
	"messenger/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser creates new user
func CreateUser(c *gin.Context) {
	user := &application.User{}
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
func FindOneUser(c *gin.Context) {}

// FindAllUsers search  userss
func FindAllUsers(c *gin.Context) {
	find := &dto.FindUsers{}
	err := c.ShouldBind(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := &application.User{ApplicationID: find.ApplicationID}

	users, total, err := user.Find(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": users, "total": total})
}

// UpdateUser changes user
func UpdateUser(c *gin.Context) {}

// DeleteUser removes user
func DeleteUser(c *gin.Context) {}

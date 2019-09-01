package handlers

import (
	"errors"
	"messenger/application"
	"messenger/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateApp creates new application
func CreateApp(c *gin.Context) {
	app := &application.Application{}
	err := c.Bind(app)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	managerID, ok := c.Get("managerID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("undefined manager id")})
		c.Abort()
		return
	}

	app.Managers = append(app.Managers, managerID.(string))

	_, err = app.Insert()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": app})
}

// FindOneApp handles http request
func FindOneApp(c *gin.Context) {}

// FindAllApp handles http request
func FindAllApp(c *gin.Context) {
	find := &dto.FindApplications{}
	err := c.ShouldBind(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		c.Abort()
		return
	}

	if len(find.ManagerID) == 0 {
		id, _ := c.Get("managerID")
		find.Managers = append(find.Managers, id.(string))
	}

	app := &application.Application{}
	apps, total, err := app.Find(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": apps, "total": total})
}

// UpdateApp changes application
func UpdateApp(c *gin.Context) {}

// UpdateAppSecret changes application secret key
func UpdateAppSecret(c *gin.Context) {
	updateApplicationSecret := &dto.UpdateApplicationSecret{}
	err := c.Bind(updateApplicationSecret)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	application := &application.Application{ID: updateApplicationSecret.ID}
	_, err = application.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": application.Secret})
}

// DeleteApp removes application
func DeleteApp(c *gin.Context) {}

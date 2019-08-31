package application

import (
	"context"
	"errors"
	"messenger/config"
	"messenger/dto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error
var ctx context.Context

func init() {
	conf, err := config.Get("mongodb")

	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(conf["connection"].(string)))

	if err != nil {
		panic(err)
	}
}

// CreateApp creates new application
func CreateApp(c *gin.Context) {
	app := &Application{}
	err := c.Bind(app)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var managerID interface{}
	if managerID, ok := c.Get("managerId"); !ok {
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
	c.ShouldBind(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(find.ManagerID) == 0 {
		id, _ := c.Get("managerId")
		find.Managers = append(find.Managers, id.(string))
	}

	app := &Application{}
	apps, total, err := app.Find(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": apps, "total": total})
}
func UpdateApp(c *gin.Context) {}
func UpdateAppSecret(c *gin.Context) {
	updateApplicationSecret := &dto.UpdateApplicationSecret{}
	err := c.Bind(updateApplicationSecret)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	application := &Application{ID: updateApplicationSecret.ID}
	_, err = application.Update()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": application.Secret})
}
func DeleteApp(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	user := &User{}
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
func FindOneUser(c *gin.Context) {}
func FindAllUsers(c *gin.Context) {
	find := &dto.FindUsers{}
	c.ShouldBind(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := &User{ApplicationID: find.ApplicationID}

	users, total, err := user.Find(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": users, "total": total})
}
func UpdateUser(c *gin.Context) {}
func DeleteUser(c *gin.Context) {}

func CreateManager(c *gin.Context) {
	manager := &Manager{}
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
func FindOneManager(c *gin.Context)  {}
func FindAllManagers(c *gin.Context) {}
func UpdateManager(c *gin.Context)   {}
func DeleteManager(c *gin.Context)   {}

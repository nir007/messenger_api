package application

import (
	"context"
	"messenger/config"
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

// CreateApp creates applications
func CreateApp(c *gin.Context) {
	app := &Application{}
	err := c.Bind(app)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	_, err = app.Insert(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": app,
	})
}
func FindOneApp(c *gin.Context) {}
func FindAllApp(c *gin.Context) {
	managerId := c.Query("manager_id")

	if len(managerId) == 0 {
		id, _ := c.Get("managerId")
		managerId = id.(string)
	}

	app := &Application{Managers: []string{managerId}}
	apps, err := app.Find(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": apps})
}
func UpdateApp(c *gin.Context) {}
func UpdateAppSecret(c *gin.Context) {
	updateApplicationSecret := &UpdateApplicationSecret{}
	err := c.Bind(updateApplicationSecret)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	application := &Application{ID: updateApplicationSecret.ID}
	_, err = application.Update(c)

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
	find := &FindUsers{}
	c.ShouldBind(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := &User{ApplicationID: find.ApplicationID}

	users, total, err := user.Find(c, find)

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

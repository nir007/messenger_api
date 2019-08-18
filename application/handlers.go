package application

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
)

var client *mongo.Client
var err error
var ctx context.Context

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}
}

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

	_, err = app.Insert()

	c.JSON(http.StatusOK, gin.H{
		"result": app,
	})
}

func FindOneApp(c *gin.Context) {}
func FindAllApp(c *gin.Context) {}
func UpdateApp(c *gin.Context) {}
func DeleteApp(c *gin.Context) {}

func CreateUser(c *gin.Context) {}
func FindOneUser(c *gin.Context) {}
func FindAllUsers(c *gin.Context) {}
func UpdateUser(c *gin.Context) {}
func DeleteUser(c *gin.Context) {}

func CreateManager(c *gin.Context) {
	fmt.Println("КАКОГО ХУЯ")

	manager := &Manager{}
	err := c.Bind(manager)

	fmt.Println(manager)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = manager.Insert()

	c.JSON(http.StatusOK, gin.H{
		"result": manager,
	})
}
func FindOneManager(c *gin.Context) {}
func FindAllManagers(c *gin.Context) {}
func UpdateManager(c *gin.Context) {}
func DeleteManager(c *gin.Context) {}

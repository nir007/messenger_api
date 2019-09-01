package routes

import (
	"messenger/handlers"
	"messenger/middlewares"

	"github.com/gin-gonic/gin"
)

//InitAPI adds routes for api
func InitAPI(r *gin.Engine) {
	r.POST("/registration", handlers.CreateManager)
	r.GET("/confirm", handlers.CreateManager)
	r.POST("/login", ginAuthJWT.LoginHandler)

	manage := r.Group("manage")
	manage.Use(ginAuthJWT.MiddlewareFunc())
	//manage.Use(handlers.ApplcationAccess())

	manage.POST("/applications", handlers.CreateApp)
	manage.GET("/applications/:id", handlers.FindOneApp)
	manage.GET("/applications", handlers.FindAllApp)
	manage.PUT("/applications", handlers.UpdateApp)
	manage.PUT("/applications/secret-key", handlers.UpdateAppSecret)
	manage.DELETE("/applications/:id", handlers.DeleteApp)

	manage.POST("/users", handlers.CreateUser)
	manage.GET("/users/:id", handlers.FindOneUser)
	manage.GET("/users", handlers.FindAllUsers)
	manage.PUT("/users", handlers.UpdateUser)
	manage.DELETE("/users/:id", handlers.DeleteUser)

	manage.POST("/messages", middlewares.Blacklist, handlers.CreateMessage)
	manage.GET("/messages/:id", handlers.FindOneMessage)
	manage.GET("/messages", handlers.FindAllMessages)
	manage.PUT("/messages", handlers.UpdateMessage)
	manage.DELETE("/messages/:id", handlers.DeleteMessage)
}

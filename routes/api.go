package routes

import (
	"messenger/application"

	"github.com/gin-gonic/gin"
)

func InitApi(r *gin.Engine) {
	r.POST("/registration", application.CreateManager)
	r.GET("/confirm", application.CreateManager)
	r.POST("/login", ginAuthJWT.LoginHandler)

	manage := r.Group("manage")
	manage.Use(ginAuthJWT.MiddlewareFunc())
	//manage.Use(application.ApplcationAccess())

	manage.POST("/applications", application.CreateApp)
	manage.GET("/applications/:id", application.FindOneApp)
	manage.GET("/applications", application.FindAllApp)
	manage.PUT("/applications", application.UpdateApp)
	manage.PUT("/applications/secret-key", application.UpdateAppSecret)
	manage.DELETE("/applications", application.DeleteApp)

	manage.POST("/users", application.CreateUser)
	manage.GET("/users/:id", application.FindOneUser)
	manage.GET("/users", application.FindAllUsers)
	manage.PUT("/users", application.UpdateUser)
	manage.DELETE("/users", application.DeleteUser)
}

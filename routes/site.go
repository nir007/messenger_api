package routes

import (
	"messenger/handlers"
	"messenger/middlewares"

	"github.com/gin-gonic/gin"
)

//InitSite adds routes for site gui
func InitSite(r *gin.Engine) {
	site := r.Group("site")

	site.POST("/registration", handlers.CreateManager)
	site.GET("/confirm", handlers.CreateManager)
	site.POST("/login", ginAuthJWT.LoginHandler)

	admin := r.Group("admin")
	admin.Use(ginAuthJWT.MiddlewareFunc())
	//manage.Use(handlers.ApplcationAccess())

	admin.POST("/applications", handlers.CreateApp)
	admin.GET("/applications/:id", handlers.FindOneApp)
	admin.GET("/applications", handlers.FindAllApp)
	admin.PUT("/applications", handlers.UpdateApp)
	admin.PUT("/applications/secret-key", handlers.UpdateAppSecret)
	admin.DELETE("/applications/:id", handlers.DeleteApp)

	admin.POST("/users", handlers.CreateUser)
	admin.GET("/users/:id", handlers.FindOneUser)
	admin.GET("/users", handlers.FindAllUsers)
	admin.PUT("/users", handlers.UpdateUser)
	admin.DELETE("/users/:id", handlers.DeleteUser)

	admin.POST("/messages", middlewares.Blacklist, handlers.CreateMessage)
	admin.GET("/messages/:id", handlers.FindOneMessage)
	admin.GET("/messages", handlers.FindAllMessages)
	admin.PUT("/messages", handlers.UpdateMessage)
	admin.DELETE("/messages/:id", handlers.DeleteMessage)

	admin.GET("/managers", handlers.FindOneManager)
}

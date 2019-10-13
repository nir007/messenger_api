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
	admin.PUT("/applications/:id", handlers.UpdateApp)
	admin.DELETE("/applications/:id", handlers.DeleteApp)

	admin.POST("/users", handlers.CreateUser)
	admin.GET("/user/:id/:appID", handlers.FindOneUser)
	admin.GET("/users/:appID", handlers.FindAllUsers)
	admin.PUT("/users/:id/:appID", handlers.UpdateUser)
	admin.DELETE("/users/:id/:appID", handlers.DeleteUser)

	admin.POST("/messages", middlewares.Blacklist, handlers.CreateMessage)
	admin.GET("/messages/:id", handlers.FindOneMessage)
	admin.GET("/messages", handlers.FindAllMessages)
	admin.PUT("/messages", handlers.UpdateMessage)
	admin.DELETE("/messages/:id", handlers.DeleteMessage)

	admin.GET("/managers", handlers.FindOneManager)
	admin.PUT("/managers", handlers.UpdateManager)
	admin.POST("/managers/avatar", handlers.UpdateManagerAvatar)
}

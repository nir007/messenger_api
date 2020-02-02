package routes

import (
	"messenger/handlers"
	"github.com/gin-gonic/gin"
	m "messenger/middlewares"
)

//InitSite adds routes for site gui
func InitSite(r *gin.Engine) {
	site := r.Group("site")

	site.POST("/registration", handlers.CreateManager)
	site.GET("/confirm/email", handlers.ConfirmEmail)
	site.POST("/login", ginAuthJWT.LoginHandler)

	admin := r.Group("admin")
	admin.Use(ginAuthJWT.MiddlewareFunc())

	admin.POST("/applications", handlers.CreateApp)
	admin.GET("/applications/:appID", m.SiteApplicationAccess, handlers.FindOneApp)
	admin.GET("/applications", handlers.FindAllApp)
	admin.PUT("/applications/:appID", m.SiteApplicationAccess, handlers.UpdateApp)
	admin.DELETE("/applications/:appID", m.SiteApplicationAccess, handlers.DeleteApp)

	admin.POST("/users", m.SiteApplicationAccess, handlers.CreateUser)
	admin.GET("/user/:id/:appID", m.SiteApplicationAccess, handlers.FindOneUser)
	admin.GET("/users/:appID", m.SiteApplicationAccess, handlers.FindAllUsers)
	admin.PUT("/users/:id/:appID", m.SiteApplicationAccess, handlers.UpdateUser)
	admin.DELETE("/users/:id/:appID", m.SiteApplicationAccess, handlers.DeleteUser)

	admin.POST("/messages", m.SiteApplicationAccess, m.Blacklist, handlers.CreateMessage)
	admin.GET("/messages/:id", m.SiteApplicationAccess, handlers.FindOneMessage)
	admin.GET("/messages", m.SiteApplicationAccess, handlers.FindAllMessages)
	admin.PUT("/messages", m.SiteApplicationAccess, handlers.UpdateMessage)
	admin.DELETE("/messages/:id", m.SiteApplicationAccess, handlers.DeleteMessage)

	admin.GET("/managers", handlers.FindOneManager)
	admin.PUT("/managers", handlers.UpdateManager)
	admin.POST("/managers/avatar", handlers.UpdateManagerAvatar)
}

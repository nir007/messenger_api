package routes

import (
	"messenger/handlers"
	"messenger/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitREST routes for REST API
func InitREST(r *gin.Engine) {
	v1 := r.Group("v1")
	v1.Use(cors.Default())
	v1.Use(middlewares.ApplicationAccess)

	v1.POST("/messages/", middlewares.Blacklist, handlers.CreateMessage)
	v1.GET("/messages/:dialogid", handlers.FindAllMessages)
	v1.GET("/dialogs/:uid", handlers.FindAllDialogs)
}

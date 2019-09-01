package routes

import (
	"messenger/handlers"
	"messenger/middlewares"

	"github.com/gin-gonic/gin"
)

// InitREST routes for REST API
func InitREST(r *gin.Engine) {
	v1 := r.Group("v1")
	v1.Use(middlewares.ApplicationAccess)
	v1.Use(middlewares.Blacklist)

	v1.POST("/messages/", handlers.CreateMessage)
}

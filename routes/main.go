package routes

import (
	"messenger/config"
	"messenger/middlewares"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginAuthJWT *jwt.GinJWTMiddleware

func init() {
	conf, _ := config.Get("jwt")

	duration := 10   //strconv.ParseInt(conf["timeout"].(string), 10, 64)
	maxRefresh := 10 //strconv.ParseInt(conf["maxRefresh"].(string), 10, 64)

	am := &middlewares.AuthMiddleware{
		Realm:       conf["realm"].(string),
		Key:         []byte(conf["key"].(string)),
		Timeout:     time.Hour * time.Duration(duration),
		MaxRefresh:  time.Hour * time.Duration(maxRefresh),
		IdentityKey: conf["identityKey"].(string),
	}
	ginAuthJWT, _ = am.GetAuthMiddleware()
}

// CORS returns cors policy handler
func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8082", "http://localhost:8081", "http://localhost:9528"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "User-Agent"}
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"}

	return cors.New(config)
}

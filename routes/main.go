package routes

import (
	"strconv"
	"messenger/config"
	"messenger/application"
	"time"
	"github.com/appleboy/gin-jwt"
)

var ginAuthJWT *jwt.GinJWTMiddleware

func init()  {
	duration, _ := strconv.ParseInt(config.Main["jwt"]["Timeout"], 10, 64)
	maxRefresh, _ := strconv.ParseInt(config.Main["jwt"]["MaxRefresh"], 10, 64)

	am := &application.AuthMiddleware{
		Realm: config.Main["jwt"]["Realm"],
		Key: []byte(config.Main["jwt"]["Key"]),
		Timeout: time.Hour * time.Duration(duration),
		MaxRefresh: time.Hour * time.Duration(maxRefresh),
		IdentityKey: config.Main["jwt"]["IdentityKey"],
	}
	ginAuthJWT, _ = am.GetAuthMiddleware()
}

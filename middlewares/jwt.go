package middlewares

import (
	"fmt"
	"log"
	"messenger/application"
	"messenger/dto"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

// AuthMiddleware struct for login managers
type AuthMiddleware struct {
	Realm       string
	Key         []byte
	Timeout     time.Duration
	MaxRefresh  time.Duration
	IdentityKey string
}

// GetAuthMiddleware provides jwt auth handlers
func (a *AuthMiddleware) GetAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       a.Realm,
		Key:         a.Key,
		Timeout:     a.Timeout,
		MaxRefresh:  a.MaxRefresh,
		IdentityKey: a.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(application.Manager); ok {
				return jwt.MapClaims{
					a.IdentityKey: v.ID,
					"email":       v.Email,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			fmt.Println("claims[a.IdentityKey]: ", claims[a.IdentityKey])

			id, _ := primitive.ObjectIDFromHex(claims[a.IdentityKey].(string))

			c.Set("managerID", claims[a.IdentityKey].(string))

			return application.Manager{
				ID:    id,
				Email: claims["email"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginValues login

			if err := c.ShouldBind(&loginValues); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			find := &dto.FindManagers{Email: loginValues.Email}
			manager := application.Manager{}
			_ = manager.FindOne(find)

			if err := bcrypt.CompareHashAndPassword(
				[]byte(manager.Password), []byte(loginValues.Password)); err == nil {
				return manager, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(application.Manager); ok {
				c.Set(a.IdentityKey, v.ID)
				c.Set("email", v.Email)
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			if c.ContentType() != "application/json" {
				c.Redirect(http.StatusMovedPermanently, "/login?m=send_auth_token")
				return
			}

			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: gopa",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return middleware, err
}

package application

import (
	"net/http"
	"fmt"
	"time"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"golang.org/x/crypto/bcrypt"
)

func ApplcationAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := c.Request.Header["Secret"]

		if len(secret) == 0 || secret[0] != "=74G34252H34434DFDGW$%RFEFSDef" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Required [Secret] in request headers",
			})
			c.Abort()
			return
		}
		fmt.Println("Access check")
		// before request
		c.Next()

		// after request
	}
}

type login struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type AuthMiddleware struct {
	Realm        string
	Key          []byte
	Timeout      time.Duration
	MaxRefresh   time.Duration
	IdentityKey  string
}

func (a *AuthMiddleware) Init() {

}

func (a *AuthMiddleware) GetAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       a.Realm,
		Key:         a.Key,
		Timeout:     a.Timeout,
		MaxRefresh:  a.MaxRefresh,
		IdentityKey: a.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(Manager); ok {
				return jwt.MapClaims{
					a.IdentityKey: v.ID,
					"email": v.Email,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return Manager{
				ID: claims[a.IdentityKey].(string),
				Email: claims["email"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginValues login

			if err := c.ShouldBind(&loginValues); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			email := loginValues.Email
			password := loginValues.Password

			manager := Manager{}
			err = manager.FindOne(map[string]interface{}{"email": email})

			if err := bcrypt.CompareHashAndPassword([]byte(manager.Password), []byte(password)); err == nil {
				return manager, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(Manager); ok {
				c.Set(a.IdentityKey, v.ID)
				c.Set("email", v.Email)
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			contentType := c.Request.Header["Content-Type"]
			if len(contentType) == 0 || contentType[0] != "application/json" {
				c.Redirect(http.StatusMovedPermanently, "/login")
				return
			}

			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return middleware, err
}

type managerExists struct {
	Email string `json:"email" binding:"required"`
}

func ManagerExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		manager := &Manager{}
		filter := &managerExists{}
		err := c.Bind(filter)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		err = manager.FindOne(map[string]interface{}{"email": filter.Email})

		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "user already exists",
			})
			c.Abort()
			return
		}

		// before request
		c.Next()

		// after request
	}
}
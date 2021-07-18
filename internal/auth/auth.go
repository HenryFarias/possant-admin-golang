package auth

import (
	jwtgin "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func Init(g *gin.Engine) *gin.RouterGroup {
	jwtMiddleware, err := createMiddleware()
	if err != nil {
		return nil
	}
	g.POST("/login", jwtMiddleware.LoginHandler)
	auth := g.Group("/auth")
	auth.Use(jwtMiddleware.MiddlewareFunc())
	auth.GET("/refresh_token", jwtMiddleware.RefreshHandler)
	return auth
}

func createMiddleware() (*jwtgin.GinJWTMiddleware, error) {
	var identityKey = "id"
	return jwtgin.New(&jwtgin.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwtgin.MapClaims {
			if v, ok := data.(*User); ok {
				return jwtgin.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwtgin.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwtgin.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwtgin.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwtgin.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}

package agent

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func AuthMiddleware(eng *gin.Engine, groups ...*gin.RouterGroup) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "edance",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin#123456") || (userID == "test" && password == "test#123456") {
				return &User{
					UserName:  userID,
					LastName:  "YuHao",
					FirstName: "Wang",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
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
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		panic("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		panic("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	eng.POST("/login", authMiddleware.LoginHandler)
	eng.POST("/logout", authMiddleware.LogoutHandler)
	eng.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(404, gin.H{"status": "404", "message": "Page not found"})
	})
	if groups != nil {
		for _, group := range groups {
			group.GET("/refresh_token", authMiddleware.RefreshHandler)
			// Refresh time can be longer than token timeout
			group.Use(authMiddleware.MiddlewareFunc())
		}
	}
}

func RegisterHandlers(eng *gin.Engine) {
	api := eng.Group("/api")
	AuthMiddleware(eng, api)
	//base
	api.GET("base/test", restTest)

	//song
	api.GET("/song", restListSong)
	api.GET("/song/:song_id", restGetSongById)
	api.POST("/song", restAddSong)
	api.PUT("/song/:song_id", restUpdateSong)
	api.DELETE("/song/:song_id", restDeleteSong)
}

package server

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/leftis/cicada/config"
	"github.com/leftis/cicada/graphql"
	"github.com/leftis/cicada/models"
	"net/http"
	"time"
)

var (
	admin models.Administrator
)

func authenticator(c *gin.Context) (interface{}, error) {
	_ = c.BindJSON(&admin)

	if admin.Authenticate() != nil {
		return admin, nil
	}

	return admin, jwt.ErrFailedAuthentication
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func createJWTMIddleware() jwt.GinJWTMiddleware {
	return jwt.GinJWTMiddleware{
		Realm:         "cicada",
		Key:           []byte(config.App.Secret),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: authenticator,
		Unauthorized:  unauthorized,
	}
}

func AdminRoutes(e *gin.Engine, env string) {
	jwt := createJWTMIddleware()

	admin := e.Group("/admin")
	admin.POST("/login", jwt.LoginHandler)
	admin.POST("/refresh_token", jwt.RefreshHandler)

	admin.Use(jwt.MiddlewareFunc())
	{
		admin.POST("/graph", gin.WrapH(&relay.Handler{Schema: graphql.Schema }))
		admin.GET("/*path", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin.tmpl", gin.H{"env": env})
		})
	}
}

func Init() {
	cd, env := config.App.CurrentDirectory, config.App.Environment
	r := gin.Default()
	r.LoadHTMLGlob(cd + "/templates/*")
	AdminRoutes(r, env)
	r.Run()
}

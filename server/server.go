package server

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/leftis/cicada/config"
	"github.com/leftis/cicada/models"
	"net/http"
	"time"
)

var admin models.Administrator

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


func Init() {
	cd, env := config.App.CurrentDirectory, config.App.Environment

	r := gin.Default()
	r.LoadHTMLGlob(cd + "/templates/*")

	m := &jwt.GinJWTMiddleware{
		Realm:         "cicada",
		Key:           []byte(config.App.Secret),
		Timeout:       time.Minute,
		MaxRefresh:    time.Minute,
		Authenticator: authenticator,
		Unauthorized:  unauthorized,
	}

	admin := r.Group("/admin")
	admin.POST("/login", m.LoginHandler)
	admin.Use(m.MiddlewareFunc())
	{
		admin.POST("/refresh_token", m.RefreshHandler)
		admin.GET("/*path", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin.tmpl", gin.H{"env": env})
		})
	}

	admin.Use()

	r.Run()
}

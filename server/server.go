package server

import (
	"github.com/gin-gonic/gin"
	"github.com/leftis/cicada/config"
	"github.com/leftis/cicada/models"
	"net/http"
)

func handleAdminLogin(c *gin.Context) {
	var admin models.Administrator

	_ = c.BindJSON(&admin)

	if admin.Authenticate() != nil {
		token := admin.GenerateJWTTokenString([]byte(configuration.App.Secret))
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}
}

func Init() {
	cd, env := configuration.App.CurrentDirectory, configuration.App.Environment

	r := gin.Default()
	r.LoadHTMLGlob(cd + "/templates/*")

	admin := r.Group("/admin")
	{
		admin.POST("/login", handleAdminLogin)
		admin.GET("/*path", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin.tmpl", gin.H{"env": env})
		})
	}

	r.Run()
}

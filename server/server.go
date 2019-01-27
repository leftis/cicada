package server

import (
	"github.com/gin-gonic/gin"
	"github.com/leftis/cicada/configuration"
	"net/http"
)

func Init(app configuration.App) {

	r := gin.Default()
	r.LoadHTMLGlob(app.CurrentDirectory + "/templates/*")

	admin := r.Group("/admin")
	{
		admin.GET("/*path", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin.tmpl", gin.H{"env": app.Config.Environment})
		})
	}

	r.Run()
}
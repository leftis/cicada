package server

import (
	"github.com/gin-gonic/gin"
	"github.com/leftis/cicada/configuration"
	"net/http"
)

func Init(app configuration.App) {

	router := gin.Default()
	router.LoadHTMLGlob(app.CurrentDirectory + "/templates/*")

	admin := router.Group("/admin")
	{
		admin.GET("/*path", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin.tmpl", gin.H{})
		})
	}

	router.Run(":8080")
}
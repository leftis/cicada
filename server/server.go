package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/leftis/cicada/configuration"
	"github.com/leftis/cicada/db"
	"github.com/leftis/cicada/models"
	"net/http"
)

var appSecret []byte
var sqlxDB *sqlx.DB

func handleAdminLogin(c *gin.Context) {
	var admin models.Administrator
	//var signedToken string

	_ = c.BindJSON(&admin)

	// Check if the user exists
	if admin.Authenticate(sqlxDB) != nil {
		// if exists create token
		token := admin.GenerateJWTTokenString(appSecret)
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}
}

func Init(app configuration.App) {
	// Grab
	appSecret = []byte(app.Secret)
	rawDb := db.Init(app)
	sqlxDB = sqlx.NewDb(rawDb.Conn, rawDb.Conf.Driver)

	r := gin.Default()
	r.Use(gin.Logger())
	r.LoadHTMLGlob(app.CurrentDirectory + "/templates/*")

	admin := r.Group("/admin")
	{
		admin.POST("/login", handleAdminLogin)
		admin.GET("/*path", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin.tmpl", gin.H{"env": app.Config.Environment})
		})
	}

	r.Run()
}

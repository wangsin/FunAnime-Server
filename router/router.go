package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	host := viper.GetString("mysql_main.host")
	database := viper.GetString("mysql_main.database")
	user := viper.GetString("mysql_main.user")
	password := viper.GetString("mysql_main.password")

	testGroup := r.Group("/v1/test")
	testGroup.Use(func(c *gin.Context) {
		fmt.Println("This Is Middleware Func 1")
	})
	{
		testGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"user-agent": c.GetHeader("user-agent"),
				"develop-environment": viper.GetString("dev.type"),
				"mysql-connection-url": fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database),
			})
		})
	}

	return r
}
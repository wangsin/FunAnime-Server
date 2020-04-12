package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TestController(c *gin.Context) {
	host := viper.GetString("mysql_main.host")
	database := viper.GetString("mysql_main.database")
	user := viper.GetString("mysql_main.user")
	password := viper.GetString("mysql_main.password")

	c.JSON(200, gin.H{
		"user-agent":           c.GetHeader("user-agent"),
		"develop-environment":  viper.GetString("dev.type"),
		"mysql-connection-url": fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database),
	})
}

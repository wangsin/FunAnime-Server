package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)
import _ "github.com/go-sql-driver/mysql"

func GetDBConnection() (*gorm.DB, error) {
	host := viper.GetString("mysql_main.host")
	database := viper.GetString("mysql_main.database")
	user := viper.GetString("mysql_main.user")
	password := viper.GetString("mysql_main.password")
	DB, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database))
	return DB, err
}

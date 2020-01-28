package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
)

var (
	DB  *gorm.DB
)

func Database() {
	host := viper.GetString("mysql_main.host")
	database := viper.GetString("mysql_main.database")
	user := viper.GetString("mysql_main.user")
	password := viper.GetString("mysql_main.password")
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database))
	if err != nil {
		DB = &gorm.DB{}
		DB.Error = err
		return
	}

	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	//todo 设置连接池
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

func migration() {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{})
}

func Redis() {

}

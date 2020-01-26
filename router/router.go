package router

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/controller"
	"sinblog.cn/FunAnime-Server/middleware/token"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	testGroup := r.Group("/v1/test")
	{
		testGroup.GET("/ping", controller.TestController)
	}

	userGroup := r.Group("/v1/user")
	{
		userGroup.POST("/register", )
		userGroup.POST("/login", controller.UserLogin)
		userAuthGroup := userGroup.Group("")
		userAuthGroup.Use(token.UserAuth())
		{
			userAuthGroup.GET("/info", )
			userAuthGroup.POST("/logout", )
		}
	}

	return r
}
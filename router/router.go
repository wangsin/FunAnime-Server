package router

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/controller"
	"sinblog.cn/FunAnime-Server/middleware/token"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1Group := r.Group("/v1")
	testGroup := v1Group.Group("/test")
	{
		testGroup.GET("/ping", controller.TestController)
	}

	userGroup := v1Group.Group("/user")
	{
		userGroup.POST("/register", controller.UserRegister)
		userGroup.POST("/smsCode", )
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
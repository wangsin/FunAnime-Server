package router

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/controller"
	"sinblog.cn/FunAnime-Server/middleware/token"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1Group := r.Group("/v1")
	{
		testGroup := v1Group.Group("/test")
		{
			testGroup.GET("/ping", controller.TestController)
		}

		userGroup := v1Group.Group("/user")
		{
			userGroup.POST("/smsCode", controller.UserSendSmsCode) // 发送验证码
			userGroup.POST("/register", controller.UserRegister)   // 注册
			userGroup.POST("/login", controller.UserLogin)         // 登陆
			userAuthGroup := userGroup.Group("")
			userAuthGroup.Use(token.UserAuth())
			{
				userAuthGroup.PUT("/supplement", controller.SuppleUserInfo) // 消息完善
				userAuthGroup.GET("/info", controller.GetUserInfo)          // 获取用户信息
				userAuthGroup.POST("/logout", controller.UserLogOut)        // 注销
			}
		}
	}

	return r
}

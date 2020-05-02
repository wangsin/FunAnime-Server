package router

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/controller"
	"sinblog.cn/FunAnime-Server/middleware/common"
	"sinblog.cn/FunAnime-Server/middleware/user"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	outerGroup := r.Group("/funanime/server/api/outer")
	outerGroup.Use(common.Cors())
	{
		basicGroup := outerGroup.Group("/basic")
		{
			basicGroup.GET("/config")
		}

		userGroup := outerGroup.Group("/user")
		{
			userGroup.POST("/smsCode", controller.UserSendSmsCode) // 发送验证码
			userGroup.POST("/register", controller.UserRegister)   // 注册
			userGroup.POST("/login", controller.UserLogin)         // 登陆
			userAuthGroup := userGroup.Group("/operate")
			userAuthGroup.Use(user.UserAuth())
			{
				userAuthGroup.PUT("/supple", controller.SuppleUserInfo) // 消息完善
				userAuthGroup.GET("/info", controller.GetUserInfo)          // 获取用户信息
				userAuthGroup.POST("/logout", controller.UserLogOut)        // 注销
			}
		}

		videoGroup := outerGroup.Group("/video")
		{
			// 视频基本接口
			videoGroup.GET("/detail/:id")
			videoGroup.GET("/list")
			videoOperateGroup := videoGroup.Group("/operate")
			videoOperateGroup.Use(user.UserAuth())
			{
				videoOperateGroup.POST("/collect")
				videoOperateGroup.POST("/share")
			}

			// 视频管理后台接口
			videoManageGroup := videoGroup.Group("/manage")
			videoManageGroup.Use(user.UserAuth())
			{
				videoManageGroup.POST("/hide")
				videoManageGroup.POST("/upload")
				videoManageGroup.PUT("/update")
				videoManageGroup.DELETE("/remove")
			}

			// 视频弹幕逻辑使用GO IM架构，此服务用于数据持久化

			// 视频评论接口
			videoCommentGroup := videoGroup.Group("/comment")
			{
				videoCommentGroup.Group("/list")
				videoCommentGroup.Use(user.UserAuth())
				{
					videoCommentGroup.POST("/like")
					videoCommentGroup.POST("/new")
				}
			}
		}
	}

	return r
}

package router

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/controller"
	"sinblog.cn/FunAnime-Server/middleware/user"
)

func OuterRouter(r *gin.Engine) {
	outerGroup := r.Group("/funanime/server/api/outer")
	{
		basicGroup := outerGroup.Group("/basic")
		{
			basicGroup.GET("/config", controller.GetBasicConfig)
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
			videoGroup.GET("/detail/:id", controller.GetVideoDetailForOuter)
			videoGroup.GET("/list", controller.GetVideoListForOuter)
			videoOperateGroup := videoGroup.Group("/operate")
			videoOperateGroup.Use(user.UserAuth())
			{
				videoOperateGroup.POST("/collect", controller.CreateCollection)
				videoOperateGroup.POST("/unCollect", controller.RemoveCollection)
				videoOperateGroup.GET("/collectList", controller.GetCollectList)
			}

			// 视频管理后台接口
			videoManageGroup := videoGroup.Group("/manage")
			videoManageGroup.Use(user.UserAuth())
			{
				videoManageGroup.GET("/uploadSign", controller.GetVideoUploadSign)
				videoManageGroup.GET("/list", controller.GetManageVideoList)
				videoManageGroup.POST("/hide", controller.HideVideo)
				videoManageGroup.POST("/upload", controller.UploadVideo)
				videoManageGroup.POST("/update", controller.UpdateVideoInfo)
				videoManageGroup.POST("/remove", controller.RemoveVideo)
				videoManageGroup.POST("/reSubmit", controller.ReSubmitVideo)
				videoManageGroup.POST("/show", controller.ShowVideo)
			}

			// 弹幕发射逻辑使用websocket，本接口用户DB数据初始化
			videoBarrageGroup := videoGroup.Group("/barrage")
			{
				videoBarrageGroup.GET("/list/:id", controller.GetBarrageList)
			}

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
}

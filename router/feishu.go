package router

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/controller"
)

func FeishuRouter(r *gin.Engine) {
	feishuGroup := r.Group("/feiShu/callback")
	{
		feishuGroup.POST("/event", controller.FeishuCallbackFunc)
	}
}
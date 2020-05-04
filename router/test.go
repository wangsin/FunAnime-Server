package router

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/controller"
)

func TestRouter(r *gin.Engine) {
	testRouter := r.Group("/funanime/server/api/7dsi19/test")
	{
		testRouter.POST("/main/config", controller.SetBasicConfig)
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/middleware/common"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(common.Cors())

	OuterRouter(r)
	InnerRouter(r)
	TestRouter(r)

	return r
}

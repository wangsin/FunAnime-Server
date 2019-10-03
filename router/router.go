package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	testGroup := r.Group("/v1/test")
	testGroup.Use(func(c *gin.Context) {
		fmt.Println("This Is Middleware Func 1")
	})
	{
		testGroup.GET("/ping", controller.TestController)
	}

	return r
}
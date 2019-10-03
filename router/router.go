package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sinblog.cn/FunAnime-Server/controller"
)

type RouterArr struct {
	Routers []*gin.Engine
}

var TotalRouter *RouterArr

func (tr *RouterArr) TestRouter() {
	testR := gin.New()
	testR.GET("/test/ping", controller.TestController)

	tr.Routers = append(tr.Routers, testR)
}

func (tr *RouterArr) PingRouter() {
	pingR := gin.New()
	pingR.GET("/test2/ping", controller.TestController)

	tr.Routers = append(tr.Routers, pingR)
}

func Init() {
	routerArr := make([]*gin.Engine, 0)
	TotalRouter = &RouterArr{
		Routers: routerArr,
	}

	TotalRouter.TestRouter()
	TotalRouter.PingRouter()
}

func Run() {
	for _, v := range TotalRouter.Routers {
		err := v.Run(":8080")
		if err != nil {
			panic(fmt.Errorf("Router Run Error, Message: %s\n", err.Error()))
		}
	}
}
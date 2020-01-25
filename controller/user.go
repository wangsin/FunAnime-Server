package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
	"sinblog.cn/FunAnime-Server/util/common"
	"sinblog.cn/FunAnime-Server/util/errno"
)

func UserLogin(ctx *gin.Context) {
	loginRequest := user.RequestInfo{}
	err := loginRequest.BindRequest(ctx)
	fmt.Println(ctx.Params)
	if err != nil {
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}
	flag := loginRequest.CheckRequest()
	if !flag {
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}
	// todo: 返回格式记得改成json

	//service
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 0,
		"message": "登录成功了吧。。",
	})
	return
}

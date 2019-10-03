package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sinblog.cn/FunAnime-Server/vo/request/user"
)

func UserLogin(ctx *gin.Context) {
	loginRequest := user.RequestInfo{}
	err := loginRequest.BindRequest(ctx)
	fmt.Println(ctx.Params)
	if err != nil {
		// TODO:合并处理
		ctx.JSON(http.StatusOK, gin.H{
			"statusCode": -200000,
			"message":    err.Error(),
		})
		ctx.Abort()
		return
	}
	flag := loginRequest.CheckRequest()
	if !flag {
		ctx.JSON(http.StatusOK, gin.H{
			"statusCode": -200000,
			"message":    ctx.Params,
		})
		ctx.Abort()
		return
	}

	//service
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 0,
		"message": "登录成功了吧。。",
	})
	return
}

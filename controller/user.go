package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
	serviceUser "sinblog.cn/FunAnime-Server/service/user"
	"sinblog.cn/FunAnime-Server/util/common"
	"sinblog.cn/FunAnime-Server/util/errno"
)

func UserLogin(ctx *gin.Context) {
	loginRequest := user.LoginRequestInfo{}
	err := loginRequest.BindRequest(ctx)
	if err != nil {
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}
	flag := loginRequest.CheckRequest()
	if !flag {
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}

	common.EchoSuccessJson(ctx, map[string]interface{}{})
	return
}

func UserRegister(ctx *gin.Context) {
	registerRequest := user.RegisterRequestInfo{}
	err := registerRequest.BindRequest(ctx)
	if err != nil {
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}

	errNo := serviceUser.RegisterUser(registerRequest)

	common.EchoJson(ctx, http.StatusOK, errNo, nil)
	return
}

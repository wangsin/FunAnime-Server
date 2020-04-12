package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/serializable/request/user"
	respUser "sinblog.cn/FunAnime-Server/serializable/response/user"
	serviceUser "sinblog.cn/FunAnime-Server/service/user"
	"sinblog.cn/FunAnime-Server/util/common"
	"sinblog.cn/FunAnime-Server/util/errno"
)

func UserSendSmsCode(ctx *gin.Context) {
	sendSmsRequest := user.SendSmsRequest{}
	err := sendSmsRequest.BindRequest(ctx)
	if err != nil {
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}

	err = serviceUser.SendSmsCode(&sendSmsRequest)
	if err != nil {
		common.EchoFailedJson(ctx, errno.SmsSendFailed)
		return
	}

	common.EchoSuccessJson(ctx, map[string]interface{}{})
	return
}

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

	token, errNo := serviceUser.LoginUser(&loginRequest)
	if errNo != errno.Success {
		common.EchoFailedJson(ctx, errNo)
		return
	}
	common.EchoSuccessJson(ctx, map[string]interface{}{"token": token})
	return
}

func UserRegister(ctx *gin.Context) {
	registerRequest := user.RegisterRequestInfo{}
	err := registerRequest.BindRequest(ctx)
	if err != nil {
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}

	errNo := serviceUser.RegisterUser(&registerRequest)

	common.EchoJson(ctx, http.StatusOK, errNo, nil)
	return
}

func SuppleUserInfo(ctx *gin.Context) {

}

func GetUserInfo(ctx *gin.Context) {
	uInfo, ok := ctx.Get("userInfo")
	if !ok {
		common.EchoJson(ctx, http.StatusOK, errno.Uncertified, nil)
		return
	}

	userInfo, ok := uInfo.(*token.UserInfo)
	if !ok {
		common.EchoFailedJson(ctx, errno.UnknownError)
		return
	}

	modelUser, errNo := serviceUser.GetUserInfo(userInfo)
	if errNo != errno.Success {
		common.EchoFailedJson(ctx, errNo)
		return
	}

	resp := respUser.BuildResponse(modelUser)
	common.EchoBaseJson(ctx, http.StatusOK, errNo, resp)
	return
}

func UserLogOut(ctx *gin.Context) {
	// todo 重写登陆方式 借助Redis实现
	//if !serviceUser.Logout(ctx) {
	//	common.EchoFailedJson(ctx, errno.UnknownError)
	//	return
	//}

	common.EchoBaseJson(ctx, http.StatusOK, errno.Success, nil)
	return
}

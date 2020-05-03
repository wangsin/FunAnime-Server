package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	reqUser "sinblog.cn/FunAnime-Server/serializable/request/user"
	respUser "sinblog.cn/FunAnime-Server/serializable/response/user"
	serviceUser "sinblog.cn/FunAnime-Server/service/user"
	"sinblog.cn/FunAnime-Server/util/common"
	"sinblog.cn/FunAnime-Server/util/errno"
	"sinblog.cn/FunAnime-Server/util/logger"
)

func UserSendSmsCode(ctx *gin.Context) {
	sendSmsRequest := reqUser.SendSmsRequest{}
	err := sendSmsRequest.BindRequest(ctx)
	if err != nil {
		logger.Error("bind_request_failed_at_UserSendSmsCode", logger.Fields{"err": err.Error()})
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
	loginRequest := reqUser.LoginRequestInfo{}
	err := loginRequest.BindRequest(ctx)
	if err != nil {
		logger.Error("params_error_at_login", logger.Fields{"err": err, "request": loginRequest})
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}
	flag := loginRequest.CheckRequest()
	if !flag {
		logger.Error("request_not_fit_at_login", logger.Fields{"request": loginRequest})
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}

	token, userInfo, errNo := serviceUser.LoginUser(&loginRequest)
	if errNo != errno.Success {
		common.EchoFailedJson(ctx, errNo)
		return
	}

	resp := respUser.BuildResponse(userInfo, token)
	common.EchoBaseJson(ctx, http.StatusOK, errNo, resp)
	return
}

func UserRegister(ctx *gin.Context) {
	registerRequest := reqUser.RegisterRequestInfo{}
	err := registerRequest.BindRequest(ctx)
	if err != nil {
		logger.Error("bind_request_failed_at_UserRegister", logger.Fields{"err": err.Error()})
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
	bu := new(reqUser.BasicUser)
	if err := bu.GetUserInfo(ctx); err != nil {
		common.EchoFailedJson(ctx, errno.Uncertified)
		return
	}

	modelUser, errNo := serviceUser.GetUserInfo(bu)
	if errNo != errno.Success {
		common.EchoFailedJson(ctx, errNo)
		return
	}

	resp := respUser.BuildResponse(modelUser, "")
	common.EchoBaseJson(ctx, http.StatusOK, errNo, resp)
	return
}

func UserLogOut(ctx *gin.Context) {
	bu := new(reqUser.BasicUser)
	if err := bu.GetUserInfo(ctx); err != nil {
		common.EchoFailedJson(ctx, errno.Uncertified)
		return
	}

	err := serviceUser.Logout(bu)
	if err != nil {
		common.EchoFailedJson(ctx, errno.Uncertified)
		return
	}

	common.EchoBaseJson(ctx, http.StatusOK, errno.Success, nil)
	return
}

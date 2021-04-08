package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sinblog.cn/FunAnime-Server/util/errno"
)

type RespFormat struct {
	ErrNo  int64  `json:"errno"`
	ErrMsg string `json:"errmsg"`
	Data   gin.H  `json:"data"`
}

func Echo(ctx *gin.Context, code int, data gin.H) {
	ctx.JSON(code, data)
}

func EchoJson(ctx *gin.Context, code int, errNo int64, data gin.H) {
	errMsg := "未知错误"
	if msg, ok := errno.ErrmsgMap[errNo]; ok {
		errMsg = msg
	}
	ctx.JSON(code, RespFormat{
		ErrNo:  errNo,
		ErrMsg: errMsg,
		Data:   data,
	})
}

func EchoBaseJson(ctx *gin.Context, code int, errNo int64, data interface{}) {
	errMsg := "未知错误"
	if msg, ok := errno.ErrmsgMap[errNo]; ok {
		errMsg = msg
	}
	ctx.JSON(code, map[string]interface{}{
		"errno":  errNo,
		"errmsg": errMsg,
		"data":   data,
	})
}

// 失败返回 需Abort
func EchoFailedJson(ctx *gin.Context, statusCode int64) {
	EchoJson(ctx, http.StatusOK, statusCode, map[string]interface{}{})
	//ctx.AbortWithStatus(http.StatusBadRequest)
}

// 成功返回
func EchoSuccessJson(ctx *gin.Context, data gin.H) {
	EchoJson(ctx, http.StatusOK, errno.Success, data)
}

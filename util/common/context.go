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

// 失败返回 需Abort
func EchoFailedJson(ctx *gin.Context, statusCode int64) {
	EchoJson(ctx, http.StatusBadRequest, statusCode, map[string]interface{}{})
	ctx.AbortWithStatus(http.StatusBadRequest)
}

// 成功返回
func EchoSuccessJson(ctx *gin.Context, data gin.H) {
	EchoJson(ctx, http.StatusOK, errno.Success, data)
}

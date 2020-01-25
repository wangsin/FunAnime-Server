package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sinblog.cn/FunAnime-Server/util/errno"
)

type RespFormat struct {
	StatusCode int64       `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func EchoJson(ctx *gin.Context, code int, format RespFormat) {
	ctx.JSON(code, format)
}

// 失败返回 需Abort
func EchoFailedJson(ctx *gin.Context, statusCode int64) {
	resp := RespFormat{}
	resp.StatusCode = statusCode
	if msg, ok := errno.ErrmsgMap[statusCode]; !ok {
		resp.Data = "未知错误"
	} else {
		resp.Data = msg
	}
	EchoJson(ctx, http.StatusBadRequest, resp)
	ctx.AbortWithStatus(http.StatusBadRequest)
}

// 成功返回
func EchoSuccessJson(ctx *gin.Context, format RespFormat) {
	EchoJson(ctx, http.StatusOK, format)
}

// 内部错误
func EchoInternalErrorJson(ctx *gin.Context, format RespFormat) {
	EchoJson(ctx, http.StatusInternalServerError, format)
	ctx.Abort()
}

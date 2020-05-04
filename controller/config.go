package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	request "sinblog.cn/FunAnime-Server/serializable/request"
	"sinblog.cn/FunAnime-Server/serializable/response"
	"sinblog.cn/FunAnime-Server/service/config"
	"sinblog.cn/FunAnime-Server/util/common"
	"sinblog.cn/FunAnime-Server/util/errno"
	"sinblog.cn/FunAnime-Server/util/logger"
)

func GetBasicConfig(ctx *gin.Context) {
	// service：获取首页轮播图图片 + 页顶路由 + 搜索框文本
	head, router, search, err := config.GetBasicConfig(ctx)
	if err != errno.Success {
		logger.Error("get_basic_config_error", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, err)
		return
	}

	// response：格式化数据
	common.EchoBaseJson(ctx, http.StatusOK, errno.Success, response.BasicConfigResp{
		CarouselImg: head,
		HeadRouter:  router,
		SearchArea:  search,
	})
	return
}

func SetBasicConfig(ctx *gin.Context) {
	// 设置基础配置
	r := request.BasicConfig{}
	err := r.GetRequest(ctx)
	if err != nil {
		logger.Error("get_request_data_failed_at_SetBasicConfig", logger.Fields{"err": err})
		common.EchoFailedJson(ctx, errno.ParamsError)
		return
	}

	errNo := config.SetBasicConfig(ctx, &r)
	if errNo != errno.Success {
		logger.Error("get_basic_config_error", logger.Fields{"err": errNo})
		common.EchoFailedJson(ctx, errNo)
		return
	}

	common.EchoSuccessJson(ctx, map[string]interface{}{})
}

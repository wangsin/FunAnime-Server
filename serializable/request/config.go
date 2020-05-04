package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	serviceStruct "sinblog.cn/FunAnime-Server/service/struct"
)

type BasicConfig struct {
	CarouselImg []*serviceStruct.CarouselInfo `json:"carousel_img"`
	HeadRouter  []*serviceStruct.BasicRouter  `json:"head_router"`
	SearchArea  string                        `json:"search_area"`
	ConfigType  int                           `json:"config_type"`
}

func (bcr *BasicConfig) GetRequest(ctx *gin.Context) error {
	if err := ctx.Bind(bcr); err != nil {
		return err
	}

	if bcr.ConfigType < 0 || len(bcr.CarouselImg) <= 0 || len(bcr.HeadRouter) <= 0 {
		return errors.New("params_error")
	}
	return nil
}

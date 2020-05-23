package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mime/multipart"
	"net/http"
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

type ImgUpload struct {
	File        multipart.File
	FileHandler *multipart.FileHeader
}

func (formatData *ImgUpload) GetRequestData(r *http.Request) error {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if file == nil || handler == nil || err != nil {
		return err
	}

	formatData.File = file
	formatData.FileHandler = handler
	return binding.Default(r.Method, r.Header.Get("Content-Type")).Bind(r, formatData)
}
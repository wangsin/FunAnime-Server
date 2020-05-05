package response

import (
	serviceStruct "sinblog.cn/FunAnime-Server/service/struct"
)

type BasicConfigResp struct {
	CarouselImg []*serviceStruct.CarouselInfo `json:"carousel_img"`
	HeadRouter  []*serviceStruct.BasicRouter  `json:"head_router"`
	SearchArea  string                        `json:"search_area"`
}

type PageData struct {
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Count int64 `json:"count"`
}

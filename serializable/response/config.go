package response

import (
	serviceStruct "sinblog.cn/FunAnime-Server/service/struct"
)

type BasicConfigResp struct {
	CarouselImg []*serviceStruct.CarouselInfo `json:"carousel_img"`
	HeadRouter  []*serviceStruct.BasicRouter  `json:"head_router"`
	SearchArea  string                        `json:"search_area"`
}

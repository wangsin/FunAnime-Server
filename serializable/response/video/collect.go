package responseVideo

import (
	"sinblog.cn/FunAnime-Server/model"
	"sinblog.cn/FunAnime-Server/serializable/response"
	"sinblog.cn/FunAnime-Server/util/common"
)

type CollectListResponse struct {
	CollectList []*CollectElement `json:"collect_list"`
	PageData    response.PageData `json:"page_data"`
}

type CollectElement struct {
	Id        int64  `json:"id"`
	VideoName string `json:"video_name"`
	VideoId   int64  `json:"video_id"`
	VideoPic  string `json:"video_pic"`
}

func FormatCollectListResponse(list []*model.FaCollection, page, size int, count int64) *CollectListResponse {
	resp := new(CollectListResponse)
	respList := make([]*CollectElement, len(list))

	for index, collection := range list {
		element := &CollectElement{
			Id:        collection.Id,
			VideoName: collection.VideoName,
			VideoId:   collection.VideoId,
			VideoPic:  common.BuildImageLink(collection.VideoCoverImg),
		}
		respList[index] = element
	}

	resp = &CollectListResponse{
		CollectList: respList,
		PageData: response.PageData{
			Page:  page,
			Size:  size,
			Count: count,
		},
	}
	return resp
}

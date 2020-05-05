package responseVideo

import (
	"sinblog.cn/FunAnime-Server/model"
	"sinblog.cn/FunAnime-Server/serializable/response"
	"sinblog.cn/FunAnime-Server/util/common"
	"sinblog.cn/FunAnime-Server/util/consts"
	utilMath "sinblog.cn/FunAnime-Server/util/math"
)

type VideoInfo struct {
	TrueImg string `json:"true_img"`
	Title   string `json:"title"`
	Volume  string `json:"volume"`
	Date    string `json:"date"`
	VideoId int64  `json:"video_id"`
}

type VideoListResponse struct {
	VideoList []*VideoInfo      `json:"video_list"`
	PageData  response.PageData `json:"page_data"`
}

func BuildVideoListResponse(videoList []*model.FaVideo, page int, size int, count int64) VideoListResponse {
	list := make([]*VideoInfo, len(videoList))
	for i, video := range videoList {
		list[i] = &VideoInfo{
			TrueImg: common.BuildImageLink(video.CoverImg),
			Title:   video.VideoName,
			Volume:  utilMath.GetHumanFormatNumber(video.Pv),
			Date:    video.PassTime.Format(consts.TimeFormatYMDHM),
			VideoId: video.Id,
		}
	}

	pageData := response.PageData{
		Page:  page,
		Size:  size,
		Count: count,
	}

	return VideoListResponse{
		VideoList: list,
		PageData: pageData,
	}
}

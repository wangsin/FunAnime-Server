package serviceVideo

import (
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/model"
	requestVideo "sinblog.cn/FunAnime-Server/serializable/request/video"
	"sinblog.cn/FunAnime-Server/util/errno"
	"sinblog.cn/FunAnime-Server/util/logger"
)

func GetVideoList(req *requestVideo.GetVideoListForOuter) ([]*model.FaVideo, int64, int64) {
	whereMap := map[string]interface{}{
		"status=?": model.FaVideoNormal,
	}

	if req.Category != 0 {
		whereMap["category_top_level=?"] = req.Category
	}

	videoList, count, err := model.GetVideoList(whereMap, "", req.Page, req.Size, "create_time desc")
	if err != nil {
		logger.Error("get_video_list_failed", logger.Fields{"err":err})
		return nil, 0, errno.DBOpError
	}

	return videoList, count, errno.Success
}

/*
	获取详情页所有信息
	返回
		视频表信息
		是否收藏
		上传用户信息
 */
func GetVideoDetail(id int64, userInfo *token.UserInfo) (*model.FaVideo, *model.User, bool, int64) {
	// 视频详情
	videoDetail, err := model.GetVideoById(id)
	if err != nil {
		logger.Error("get_video_by_id_failed_at_GetVideoDetail", logger.Fields{"err":err, "id":id})
		return nil, nil, false, errno.NotFound
	}

	if videoDetail.Status != model.FaVideoNormal {
		logger.Error("get_video_by_id_failed_at_GetVideoDetail", logger.Fields{"err":err, "id":id})
		return nil, nil, false, errno.NotFound
	}

	collected := false
	if userInfo != nil {
		whereMap := map[string]interface{}{
			"video_id": videoDetail.Id,
			"status": model.FaCollectionNormalStatus,
		}

		whereMap["user_id"] = userInfo.Id
		// 是否收藏
		collection, _, err := model.GetCollectionByWhereMap(whereMap, "", 1, 1, "")
		if err != nil {
			logger.Error("get_collection_by_where_map_failed", logger.Fields{"err":err})
			return nil, nil, false, errno.DBOpError
		}

		if len(collection) > 0 {
			collected = true
		}
	}

	uploaderInfo, err := model.QueryUserWithId(videoDetail.Creator)
	if err != nil {
		logger.Error("get_user_by_id_failed", logger.Fields{"err":err})
		return nil, nil, false, errno.DBOpError
	}

	return videoDetail, uploaderInfo, collected, errno.Success
}
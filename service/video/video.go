package serviceVideo

import (
	"fmt"
	"sinblog.cn/FunAnime-Server/middleware/token"
	"sinblog.cn/FunAnime-Server/model"
	requestVideo "sinblog.cn/FunAnime-Server/serializable/request/video"
	"sinblog.cn/FunAnime-Server/util/errno"
	"sinblog.cn/FunAnime-Server/util/logger"
	"strings"
	"time"
)

func GetVideoList(req *requestVideo.GetVideoListForOuter) ([]*model.FaVideo, int64, int64) {
	whereText := make([]string, 0)
	whereMap := map[string]interface{}{
		"status=?": model.FaVideoNormal,
	}

	if req.Category != 0 {
		whereMap["category_top_level=?"] = req.Category
	}

	if req.Title != "" {
		whereText = append(whereText, fmt.Sprintf("video_name LIKE '%%%s%%'", req.Title))
	}

	videoList, count, err := model.GetVideoList(whereMap, strings.Join(whereText, " and "), req.Page, req.Size, "create_time desc")
	if err != nil {
		logger.Error("get_video_list_failed", logger.Fields{"err": err})
		return nil, 0, errno.DBOpError
	}

	return videoList, count, errno.Success
}

func GetFixUserVideoList(userInfo *token.UserInfo, page, size int) ([]*model.FaVideo, int64, int64) {
	videoList, count, err := model.GetVideoList(map[string]interface{}{
		"creator=?": userInfo.UserId,
		"status>?":  -1,
	}, "", page, size, "create_time desc")
	if err != nil {
		logger.Error("get_video_list_failed", logger.Fields{"err": err})
		return nil, 0, errno.DBOpError
	}

	return videoList, count, errno.Success
}

func GetVideoDetail(id int64, userInfo *token.UserInfo) (*model.FaVideo, *model.User, bool, int64) {
	// 视频详情
	videoDetail, err := model.GetVideoById(id)
	if err != nil {
		logger.Error("get_video_by_id_failed_at_GetVideoDetail", logger.Fields{"err": err, "id": id})
		return nil, nil, false, errno.NotFound
	}

	if videoDetail.Status != model.FaVideoNormal {
		logger.Error("get_video_by_id_failed_at_GetVideoDetail", logger.Fields{"err": err, "id": id})
		return nil, nil, false, errno.NotFound
	}

	collected := false
	if userInfo != nil {
		whereMap := map[string]interface{}{
			"video_id": videoDetail.Id,
			"status":   model.FaCollectionNormalStatus,
		}

		whereMap["user_id"] = userInfo.Id
		// 是否收藏
		collection, _, err := model.GetCollectionByWhereMap(whereMap, "", 1, 1, "")
		if err != nil {
			logger.Error("get_collection_by_where_map_failed", logger.Fields{"err": err})
			return nil, nil, false, errno.DBOpError
		}

		if len(collection) > 0 {
			collected = true
		}
	}

	uploaderInfo, err := model.QueryUserWithId(videoDetail.Creator)
	if err != nil {
		logger.Error("get_user_by_id_failed", logger.Fields{"err": err})
		return nil, nil, false, errno.DBOpError
	}

	// 调用一次PV+1
	go asyncAddPv(videoDetail.Id, videoDetail.Pv)

	return videoDetail, uploaderInfo, collected, errno.Success
}

func asyncAddPv(videoId int64, pv int64) {
	db, err := model.GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err": err})
		return
	}

	tx := db.Begin()
	err = model.UpdateVideoWithTrans(tx, map[string]interface{}{
		"id": videoId,
	}, map[string]interface{}{
		"pv": pv + 1,
	}, 1)
	if err != nil {
		logger.Error("create_video_failed", logger.Fields{"err": err})
		tx.Rollback()
		return
	}
	tx.Commit()
}

func UploadVideo(req *requestVideo.UploadRequest) int64 {
	videoInfo := &model.FaVideo{
		VideoName:             req.Name,
		VideoRemoteId:         req.RemoteId,
		VideoDesc:             req.Desc,
		CategoryTopLevel:      req.CategoryTop,
		CategoryTopLevelDesc:  req.CategoryTopDesc,
		CategoryNextLevel:     req.CategoryNext,
		CategoryNextLevelDesc: req.CategoryNextDesc,
		CoverImg:              req.CoverImg,
		Creator:               req.UserInfo.UserId,
		Status:                model.FaVideoAuditing,
		PassTime:              time.Now(),
		CreateTime:            time.Now(),
		ModifyTime:            time.Now(),
	}

	db, err := model.GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err": err})
		return errno.DBOpError
	}

	tx := db.Begin()
	err = model.CreateVideoWithTrans(tx, videoInfo)
	if err != nil {
		logger.Error("create_video_failed", logger.Fields{"err": err})
		tx.Rollback()
		return errno.DBOpError
	}
	tx.Commit()

	return errno.Success
}

func GetBarrageList(id int64) ([]*model.FaBarrage, int64) {
	barrageList, _, err := model.GetBarrageList(map[string]interface{}{
		"video_id=?": id,
		"status=?":   1,
	})
	if err != nil {
		logger.Error("get_barrage_list_failed", logger.Fields{"err": err})
		return nil, errno.DBOpError
	}

	return barrageList, errno.Success
}

func VideoStatusTrans(req *requestVideo.CreateCollection, status int) int64 {
	db, err := model.GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err": err})
		return errno.DBOpError
	}

	tx := db.Begin()
	err = model.UpdateVideoWithTrans(tx,
		map[string]interface{}{
			"id": req.VideoId,
		},
		map[string]interface{}{
			"status":      status,
			"modify_time": time.Now(),
		}, 1)
	if err != nil {
		logger.Error("trans_video_failed", logger.Fields{"err": err})
		tx.Rollback()
		return errno.DBOpError
	}
	tx.Commit()

	return errno.Success
}

func UpdateVideoInfo(req *requestVideo.UpdateVideoInfo) int64 {
	db, err := model.GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err": err})
		return errno.DBOpError
	}

	tx := db.Begin()
	err = model.UpdateVideoWithTrans(tx,
		map[string]interface{}{
			"id":     req.VideoId,
			"status": model.FaVideoNormal,
		},
		map[string]interface{}{
			"video_name":  req.VideoName,
			"video_desc":  req.VideoDesc,
			"status":      model.FaVideoAuditing,
			"modify_time": time.Now(),
		}, 1)
	if err != nil {
		logger.Error("trans_video_failed", logger.Fields{"err": err})
		tx.Rollback()
		return errno.DBOpError
	}
	tx.Commit()

	return errno.Success
}

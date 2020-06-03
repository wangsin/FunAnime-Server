package serviceVideo

import (
	"sinblog.cn/FunAnime-Server/model"
	requestVideo "sinblog.cn/FunAnime-Server/serializable/request/video"
	"sinblog.cn/FunAnime-Server/util/errno"
	"sinblog.cn/FunAnime-Server/util/logger"
	"strings"
	"time"
)

func GetUserCollection(req *requestVideo.GetVideoListForCollection) ([]*model.FaCollection, int64, int64) {
	whereMap, whereText := getUserCollectionWhereMap(req)
	list, count, err := model.GetCollectionByWhereMap(whereMap, strings.Join(whereText, " and "), req.Page, req.Size, "create_time desc")
	if err != nil {
		logger.Warn("get_collect_list_failed", logger.Fields{"err": err, "req": req})
		return nil, 0, errno.DBOpError
	}

	return list, count, errno.Success
}

func getUserCollectionWhereMap(req *requestVideo.GetVideoListForCollection) (map[string]interface{}, []string) {
	whereMap := make(map[string]interface{})
	whereText := make([]string, 0)
	if req.Name != "" {
		whereMap["video_name LIKE '%?%'"] = req.Name
	}
	whereMap["user_id=?"] = req.UserInfo.UserId
	whereMap["status<>?"] = model.FaCollectionDeleteStatus
	return whereMap, whereText
}

func CreateUserCollection(req *requestVideo.CreateCollection) int64 {
	videoInfo, err := model.GetVideoById(req.VideoId)
	if err != nil {
		logger.Warn("get_video_by_id_failed", logger.Fields{"err":err})
		return errno.DBOpError
	}

	instance := &model.FaCollection{
		VideoId:       videoInfo.Id,
		UserId:        req.UserInfo.UserId,
		VideoCoverImg: videoInfo.CoverImg,
		VideoName:     videoInfo.VideoName,
		Status:        model.FaCollectionNormalStatus,
		CreateTime:    time.Now(),
		ModifyTime:    time.Now(),
	}

	db, err := model.GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err":err})
		return errno.DBOpError
	}

	tx := db.Begin()
	err = model.CreateCollectionByInstance(tx, instance)
	if err != nil {
		tx.Rollback()
		logger.Error("create_collection_failed", logger.Fields{"err":err})
		return errno.DBOpError
	}

	tx.Commit()
	return errno.Success
}

func RemoveUserCollection(req *requestVideo.RemoveCollection) int64 {
	db, err := model.GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err":err})
		return errno.DBOpError
	}

	tx := db.Begin()
	err = model.UpdateCollectionByMap(tx, map[string]interface{}{
		"video_id": req.VideoId,
	}, map[string]interface{}{
		"status": model.FaCollectionDeleteStatus,
	})
	if err != nil {
		tx.Rollback()
		logger.Error("create_collection_failed", logger.Fields{"err":err})
		return errno.DBOpError
	}

	tx.Commit()
	return errno.Success
}
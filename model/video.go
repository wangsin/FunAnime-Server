package model

import (
	"github.com/jinzhu/gorm"
	"sinblog.cn/FunAnime-Server/util/logger"
	"time"
)

const FaVideoTableName = "fa_video"

//-1已删除，1审核中，2审核失败，3隐藏，4正常
const (
	FaVideoDeleter = iota - 1
	_
	FaVideoAuditing
	FaVideoNotFit
	FaVideoHide
	FaVideoNormal
)


type FaVideo struct {
	Id                    int64     `json:"id" gorm:"column:id"`
	VideoName             string    `json:"video_name" gorm:"column:video_name"`
	VideoRemoteId         string    `json:"video_remote_id" gorm:"column:video_remote_id"`
	VideoDesc             string    `json:"video_desc" gorm:"column:video_desc"`
	CategoryTopLevel      int64     `json:"category_top_level" gorm:"column:category_top_level"`
	CategoryTopLevelDesc  string    `json:"category_top_level_desc" gorm:"column:category_top_level_desc"`
	CategoryNextLevel     int64     `json:"category_next_level" gorm:"column:category_next_level"`
	CategoryNextLevelDesc string    `json:"category_next_level_desc" gorm:"column:category_next_level_desc"`
	CoverImg              string    `json:"cover_img" gorm:"column:cover_img"`
	Creator               int64     `json:"creator" gorm:"column:creator"`
	Pv                    int64     `json:"pv" gorm:"column:pv"`
	Uv                    int64     `json:"uv" gorm:"column:uv"`
	Status                int       `json:"status" gorm:"column:status"`
	PassTime              time.Time `json:"pass_time" gorm:"column:pass_time"`
	CreateTime            time.Time `json:"create_time" gorm:"column:create_time"`
	ModifyTime            time.Time `json:"modify_time" gorm:"column:modify_time"`
}

func (fv *FaVideo) TableName() string {
	return FaVideoTableName
}

func GetVideoList(whereMap map[string]interface{}, whereText string, page, size int, order string) ([]*FaVideo, int64, error) {
	db, err := GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err": err})
		return nil, 0, err
	}

	for s, i := range whereMap {
		db = db.Where(s, i)
	}

	videoList := make([]*FaVideo, 0)
	var count int64
	db = db.Debug().Table(FaVideoTableName).Where(whereText)
	db.Count(&count)
	err = db.Offset((page - 1) * size).Limit(size).Order(order).Find(&videoList).Error
	return videoList, count, err
}

func GetVideoById(videoId int64) (*FaVideo, error) {
	db, err := GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err": err})
		return nil, err
	}

	videoInfo := new(FaVideo)
	err = db.Debug().Table(FaVideoTableName).Where("id=?", videoId).Find(&videoInfo).Error
	return videoInfo, err
}

func CreateVideoWithTrans(tx *gorm.DB, instance *FaVideo) error {
	return tx.Debug().Table(FaVideoTableName).Create(instance).Error
}

func UpdateVideoWithTrans(tx *gorm.DB, whereMap, updateMap map[string]interface{}, limit int) error {
	return tx.Debug().Table(FaVideoTableName).Where(whereMap).Limit(limit).Update(updateMap).Error
}
package model

import (
	"sinblog.cn/FunAnime-Server/util/logger"
	"time"
)

const FaBarrageTableName = "fa_barrage"

type FaBarrage struct {
	Id           int64     `json:"id" gorm:"column:id"`
	VideoId      int64     `json:"video_id" gorm:"column:video_id"`
	Creator      int64     `json:"creator" gorm:"column:creator"`
	BarrageText  string    `json:"barrage_text" gorm:"column:barrage_text"`
	BarrageColor string    `json:"barrage_color" gorm:"column:barrage_color"`
	Status       int       `json:"status" gorm:"column:status"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
	ModifyTime   time.Time `json:"modify_time" gorm:"column:modify_time"`
}

func CreateBarrage(instance *FaBarrage) error {
	db, err := GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err": err})
		return err
	}

	return db.Debug().Table(FaBarrageTableName).Create(&instance).Error
}

func GetBarrageList(whereMap map[string]interface{}) ([]*FaBarrage, int64, error) {
	db, err := GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err": err})
		return nil, 0, err
	}

	var count int64
	var barrageList []*FaBarrage

	for s, i := range whereMap {
		db = db.Where(s, i)
	}

	err = db.Debug().Table(FaBarrageTableName).Count(&count).Find(&barrageList).Error
	return barrageList, count, err
}
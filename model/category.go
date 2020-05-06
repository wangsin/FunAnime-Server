package model

import (
	"sinblog.cn/FunAnime-Server/util/logger"
	"time"
)

const FaCategoryTableName = "fa_category"

type FaCategory struct {
	Id               int64     `json:"id" gorm:"column:id"`
	CategoryName     string    `json:"category_name" gorm:"column:category_name"`
	ParentCategoryId int64     `json:"parent_category_id" gorm:"column:parent_category_id"`
	Creator          int64     `json:"creator" gorm:"column:creator"`
	CreateTime       time.Time `json:"create_time" gorm:"column:create_time"`
	ModifyTime       time.Time `json:"modify_time" gorm:"column:modify_time"`
}

func (fcObj *FaCategory) TableName() string {
	return FaCategoryTableName
}

func GetFaCategoryList(whereMap map[string]interface{}) ([]*FaCategory, error) {
	db, err := GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed", logger.Fields{"err":err})
		return nil, err
	}

	list := make([]*FaCategory, 0)

	for s, i := range whereMap {
		db = db.Where(s, i)
	}

	err = db.Debug().Table(FaCategoryTableName).Find(&list).Error
	return list, err
}
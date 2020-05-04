package model

import (
	"github.com/jinzhu/gorm"
	"sinblog.cn/FunAnime-Server/util/logger"
	"time"
)

const FaConfigTableName = "fa_config"

const (
	FaConfigStatusDeleted = iota - 1
	FaConfigStatusValid
)

type FaConfig struct {
	Id         int64     `json:"id" gorm:"column:id"`
	ConfigType int       `json:"config_type" gorm:"column:config_type"`
	Status     int       `json:"status" gorm:"column:status"`
	ConfigData string    `json:"config_data" gorm:"column:config_data"`
	ConfigUser int64     `json:"config_user" gorm:"column:config_user"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	ModifyTime time.Time `json:"modify_time" gorm:"column:modify_time"`
}

func (fc *FaConfig) TableName() string {
	return FaConfigTableName
}

func GetBasicConfigList(whereMap map[string]interface{}) ([]*FaConfig, error) {
	db, err := GetDatabaseConnection()
	if err != nil {
		logger.Error("get_db_conn_failed_at_GetBasicConfigList", logger.Fields{"err": err})
		return nil, err
	}

	result := make([]*FaConfig, 0)

	for s, i := range whereMap {
		db = db.Where(s, i)
	}

	err = db.Debug().Table(FaConfigTableName).Where("status=?", FaConfigStatusValid).Find(&result).Error
	return result, err
}

func CreateConfigWithTrans(tx *gorm.DB, instance *FaConfig) error {
	return tx.Debug().Table(FaConfigTableName).Create(&instance).Error
}

func UpdateConfigWithTrans(tx *gorm.DB, whereMap, updateMap map[string]interface{}, limit int) error {
	return tx.Debug().Table(FaConfigTableName).Where(whereMap).Limit(limit).Update(updateMap).Error
}
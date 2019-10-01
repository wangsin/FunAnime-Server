package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Log struct {
	gorm.Model
	Id        int64 `gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	Level     int
	ForType   int
	LogTime   time.Time
	Errno     int64
	Message   string
	ExtraInfo string
}

func SaveLogInfo(instance *Log) (int64, error) {
	err := DB.Model(&Log{}).Save(instance).Error
	if err != nil {
		return -1, err
	}
	return instance.Id, nil
}

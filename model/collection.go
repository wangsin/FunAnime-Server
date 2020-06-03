package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const FaCollectionTableName = "fa_collection"

const FaCollectionNormalStatus = 1
const FaCollectionDeleteStatus = -1

type FaCollection struct {
	Id            int64     `json:"id" gorm:"column:id"`
	CollectionId  int64     `json:"collection_id" gorm:"column:collection_id"`
	VideoId       int64     `json:"video_id" gorm:"column:video_id"`
	UserId        int64     `json:"user_id" gorm:"column:user_id"`
	VideoCoverImg string    `json:"video_cover_img" gorm:"column:video_cover_img"`
	VideoName     string    `json:"video_name" gorm:"column:video_name"`
	Status        int       `json:"status" gorm:"column:status"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
	ModifyTime    time.Time `json:"modify_time" gorm:"column:modify_time"`
}

func (fc *FaCollection) TableName() string {
	return FaCollectionTableName
}

func GetCollectionByWhereMap(whereMap map[string]interface{}, whereText string, page, size int, order string) ([]*FaCollection, int64, error) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return nil, 0, err
	}

	for s, i := range whereMap {
		db = db.Where(s, i)
	}

	videoList := make([]*FaCollection, 0)
	var count int64
	db = db.Debug().Table(FaCollectionTableName).Where(whereText)
	db.Count(&count)
	err = db.Offset((page - 1) * size).Limit(size).Order(order).Find(&videoList).Error
	return videoList, count, err
}

func CreateCollectionByInstance(tx *gorm.DB, instance *FaCollection) error {
	return tx.Debug().Table(FaCollectionTableName).Create(instance).Error
}

func UpdateCollectionById(tx *gorm.DB, id int64, updateMap map[string]interface{}) error {
	return tx.Debug().Table(FaCollectionTableName).Where(map[string]interface{}{"id": id}).Update(updateMap).Error
}

func UpdateCollectionByMap(tx *gorm.DB, whereMap, updateMap map[string]interface{}) error {
	return tx.Debug().Table(FaCollectionTableName).Where(whereMap).Update(updateMap).Error
}
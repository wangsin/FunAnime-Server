package model

import (
	"time"
)

const FaUserCollectionTableName = "fa_user_collection"

type FaUserCollection struct {
	Id             int64     `json:"id" gorm:"column:id"`
	UserId         int64     `json:"user_id" gorm:"column:user_id"`
	CollectionName string    `json:"collection_name" gorm:"column:collection_name"`
	CoverImg       string    `json:"cover_img" gorm:"column:cover_img"`
	IsShare        int       `json:"is_share" gorm:"column:is_share"`
	Status         int       `json:"status" gorm:"column:status"`
	CreateTime     time.Time `json:"create_time" gorm:"column:create_time"`
	ModifyTime     time.Time `json:"modify_time" gorm:"column:modify_time"`
}

func (fuc *FaUserCollection) TableName() string {
	return FaUserCollectionTableName
}
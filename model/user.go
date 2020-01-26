package model

import (
	"time"
)

const UserTableName = "fa_user"

const (
	UserDeleted = iota - 1
	UserAvailable
	UserNotActive
	UserBanned
)

type User struct {
	Id                int64     `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Username          string    `gorm:"column:user_name;UNIQUE"`
	Nickname          string    `gorm:"column:nick_name"`
	Password          string    `gorm:"column:password"`
	Phone             string    `gorm:"column:phone;UNIQUE"`
	Sex               int8      `gorm:"column:sex"`
	DefaultCollection int64     `gorm:"column:default_collection_id;UNIQUE"`
	CollectionId      string    `gorm:"column:collection_id"`
	HistoryId         int64     `gorm:"column:history_id;UNIQUE"`
	Level             int       `gorm:"column:level"`
	Mail              string    `gorm:"column:mail;UNIQUE"`
	Birthday          time.Time `gorm:"column:birthday"`
	Avatar            string    `gorm:"column:avatar;size:1000"`
	Status            int       `gorm:"column:status"`
	CreateTime        time.Time `gorm:"column:create_time"`
	ModifyTime        time.Time `gorm:"column:modify_time"`
}

func (u User) TableName() string {
	return UserTableName
}

func CreateUserWithInstance(u *User) (int64, error) {
	if DB == nil || DB.Error != nil {
		return 0, DB.Error
	}
	DB.Debug().Table(UserTableName).Create(u)
	return u.Id, DB.Error
}

func QueryUserWithWhereMap(where, whereText map[string]interface{}) ([]*User, int64, error) {
	if DB == nil || DB.Error != nil {
		return nil, 0, DB.Error
	}
	var count int64
	var userList []*User
	DB.Debug().Table(UserTableName).Where(where)
	for wKey, wText := range whereText {
		DB.Where(wKey, wText)
	}
	DB.Count(&count)
	DB.Find(userList)
	return userList, count, DB.Error
}

func QueryUserWithId(userId int64) (*User, error) {
	if DB == nil || DB.Error != nil {
		return nil, DB.Error
	}
	var userInfo *User
	DB.Debug().Table(UserTableName).Where("id = ?", userId).Find(userInfo)
	return userInfo, DB.Error
}

func UpdateUserWithId(userId int64, updateMap map[string]interface{}) error {
	if DB == nil || DB.Error != nil {
		return DB.Error
	}
	return DB.Debug().Table(UserTableName).Where("id = ?", userId).Update(updateMap).Error
}

func UpdateUserWithMap(updateMap, where, whereText map[string]interface{}) (int64, error) {
	if DB == nil || DB.Error != nil {
		return 0, DB.Error
	}
	for wKey, wText := range whereText {
		DB.Where(wKey, wText)
	}
	db := DB.Debug().Table(UserTableName).Where(where).Update(updateMap)
	return db.RowsAffected, db.Error
}

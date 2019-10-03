package model

import (
	"time"
)

const UserTableName = "fa_user"

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
	UnEditFlag        int       `gorm:"column:username_edit_flag"`
	Birthday          time.Time `gorm:"column:birthday"`
	Mail              string    `gorm:"column:mail;UNIQUE"`
}

func (u User) TableName() string {
	return UserTableName
}

func CreateUserWithInstance(u *User) (int64, error) {
	db, err := GetDBConnection()
	if err != nil || db == nil {
		// todo：添加日志功能后记得添加日志
		return 0, err
	}
	db.Debug().Table(UserTableName).Create(u)
	return u.Id, db.Error
}

func QueryUserWithWhereMap(where map[string]interface{}, whereText []string) ([]*User, int64, error) {
	db, err := GetDBConnection()
	if err != nil || db == nil {
		// todo：添加日志功能后记得添加日志
		return nil, 0, err
	}
	var count int64
	var userList []*User
	db.Debug().Table(UserTableName).Where(where)
	for _, wText := range whereText {
		db.Where(wText)
	}
	db.Count(&count)
	db.Find(userList)
	return userList, count, db.Error
}

func QueryUserWithId(userId int64) (*User, error) {
	db, err := GetDBConnection()
	if err != nil || db == nil {
		// todo：添加日志功能后记得添加日志
		return nil, err
	}
	var userInfo *User
	db.Debug().Table(UserTableName).Where("id = ?", userId).Find(userInfo)
	return userInfo, db.Error
}
package models

import "gorm.io/gorm"

type (
	// 定义原始的数据库表字段
	UserInfoModel struct {
		gorm.Model
		Name  string `json:"name"`
		Sex   string `json:"sex"`
		Phone int    `json:"phone"`
		City  string `json:"city"`
	}
	// 处理返回的字段
	TransformedUserInfo struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
)


type UserDB struct {
	Db *gorm.DB
}

func NewUsersDB(db *gorm.DB) *UserDB {
	return &UserDB{
		db,
	}
}
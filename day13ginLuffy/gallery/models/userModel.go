package models

import "gorm.io/gorm"

// 用户 -> restful http接口 ——> 序列化 --> 入库
// 数据库 --> 反序列化 -> restful http接口  ——-> 用户
type (
	// 定义用户原始的数据库字段
	UserInfoModel struct {
		gorm.Model
		Name string `json:"name"`
		Sex string `json:"sex"`
		Phone int `json:"phone"`
		City string `json:"city"`
	}

	// 处理返回的字段
	TransformedUserInfo struct {
		ID uint `json:"id"`
		Name string `json:"name"`
		Phone int `json:"phone"`
	}
)

// 为model创造构造函数，方便其他地方调用
type UserDB struct {
	Db *gorm.DB
}


func NewUsersDB(db *gorm.DB) *UserDB{
	return &UserDB{
		db,
	}
}
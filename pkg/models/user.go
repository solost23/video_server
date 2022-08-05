package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string `gorm:"comment: 用户名" json:"user_name"`
	Password     string `gorm:"comment: 用户密码" json:"password"`
	Nickname     string `gorm:"comment: 昵称" json:"nickname"`
	Role         string `gorm:"comment: 用户角色;type:enum('ADMIN','USER');default:USER" json:"role"`
	Avatar       string `gorm:"comment: 用户头像" json:"avatar"`
	Introduce    string `gorm:"comment: 用户介绍" json:"introduce"`
	FansCount    int64  `gorm:"comment: 用户粉丝数;default:0"`
	CommentCount int64  `gorm:"comment: 用户评论数;default:0"`
	// DeleteStatus string `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
}

func (u *User) TableName() string {
	return "users"
}

func (t *User) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *User) Delete(db *gorm.DB, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Error
}

func (t *User) Updates(db *gorm.DB, value interface{}, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Updates(value).Error
}

func (t *User) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (user *User, err error) {
	err = db.Model(&t).Where(query, args...).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (t *User) WhereAll(db *gorm.DB, query interface{}, args ...interface{}) (categories []*User, err error) {
	err = db.Model(&t).Where(query, args...).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (t *User) PageListOrder(db *gorm.DB, order string, params *ListPageInput, conditions interface{}, args ...interface{}) (users []*User, total int64, err error) {
	if order == "" {
		order = "created_at DESC"
	}
	offset := (params.Page - 1) * params.Size

	err = db.Model(&t).Where(conditions, args...).Offset(offset).Limit(params.Size).Order(order).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Model(&t).Where(conditions, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

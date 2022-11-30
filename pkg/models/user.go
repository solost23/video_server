package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName      string    `json:"userName" gorm:"column:user_name;type:varchar(100);comment: 用户名"`
	Password      string    `json:"password" gorm:"column:password;type:varchar(300);comment: 用户密码"`
	Nickname      string    `json:"nickname" gorm:"column:nickname;type:varchar(100);comment: 昵称"`
	Role          string    `json:"role" gorm:"column:role;comment: 用户角色;type:enum('ADMIN','USER');default:USER;comment: 用户角色"`
	Avatar        string    `json:"avatar" gorm:"column:avatar;type:text;comment: 用户头像"`
	Introduce     string    `json:"introduce" gorm:"column:introduce;type:varchar(300);comment: 用户介绍"`
	FansCount     int64     `json:"fansCount" gorm:"column:fans_count;type:bigint unsigned;comment: 用户粉丝数;default:0"`
	CommentCount  int64     `json:"commentCount" gorm:"column:comment_count;type:bigint unsigned;comment: 用户评论数;default:0"`
	LastLoginTime time.Time `json:"lastLoginTime" gorm:"column:last_login_time;type:datetime;comment: 用户上次登录时间"`
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

func (t *User) WhereAll(db *gorm.DB, query interface{}, args ...interface{}) (users []*User, err error) {
	err = db.Model(&t).Where(query, args...).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
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

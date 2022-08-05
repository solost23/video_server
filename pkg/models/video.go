package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserID       uint   `gorm:"comment: 用户 ID"`
	CategoryID   uint   `gorm:"comment: 视频分类 ID"`
	Title        string `gorm:"comment: 视频标题" json:"title" form:"title"`
	Introduce    string `gorm:"comment: 视频介绍" json:"introduce" form:"introduce"`
	ImageUrl     string `gorm:"comment: 视频封面oss地址" json:"image_url" form:"image_url"`
	VideoUrl     string `gorm:"comment: 视频流oss地址"`
	ThumbCount   int64  `gorm:"comment: 点赞数;default:0"`
	CommentCount int64  `gorm:"comment: 评论数;default:0"`
	//DeleteStatus string `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`  // 定时任务清除 oss 中的视频 if 数据库中无视频信息
}

func (v *Video) TableName() string {
	return "videos"
}

func (t *Video) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *Video) Delete(db *gorm.DB, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Error
}

func (t *Video) Updates(db *gorm.DB, value interface{}, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Updates(value).Error
}

func (t *Video) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (category *Category, err error) {
	err = db.Model(&t).Where(query, args...).First(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (t *Video) WhereAll(db *gorm.DB, query interface{}, args ...interface{}) (categories []*Category, err error) {
	err = db.Model(&t).Where(query, args...).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (t *Video) PageListOrder(db *gorm.DB, order string, params *ListPageInput, conditions interface{}, args ...interface{}) (categories []*Category, total int64, err error) {
	if order == "" {
		order = "created_at DESC"
	}
	offset := (params.Page - 1) * params.Size

	err = db.Model(&t).Where(conditions, args...).Offset(offset).Limit(params.Size).Order(order).Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Model(&t).Where(conditions, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return categories, total, nil
}

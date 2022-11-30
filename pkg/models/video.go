package models

import (
	"gorm.io/gorm"
)

type Video struct {
	CreatorBase
	UserId       uint   `json:"userId" gorm:"column:user_id;type:bigint unsigned;comment: 用户 ID"`
	CategoryId   uint   `json:"categoryId" gorm:"column:category_id;type:bigint unsigned;comment: 视频分类 ID"`
	Title        string `json:"title" gorm:"column:title;type:varchar(100);comment: 视频标题"`
	Introduce    string `json:"introduce" gorm:"column:introduce;type:varchar(500);comment: 视频介绍"`
	ImageUrl     string `json:"imageUrl" gorm:"column:image_url;type:text;comment: 视频封面oss地址"`
	VideoUrl     string `json:"videoUrl" gorm:"column:video_url;type:text;comment: 视频流oss地址"`
	ThumbCount   int64  `json:"thumbCount" gorm:"column:thumb_count;type:bigint unsigned;comment: 点赞数;default:0"`
	CommentCount int64  `json:"commentCount" gorm:"column:comment_count;type:bigint unsigned;comment: 评论数;default:0"`
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

func (t *Video) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (video *Video, err error) {
	err = db.Model(&t).Where(query, args...).First(&video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (t *Video) WhereAll(db *gorm.DB, query interface{}, args ...interface{}) (videos []*Video, err error) {
	err = db.Model(&t).Where(query, args...).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (t *Video) PageListOrder(db *gorm.DB, order string, params *ListPageInput, conditions interface{}, args ...interface{}) (videos []*Video, total int64, err error) {
	if order == "" {
		order = "created_at DESC"
	}
	offset := (params.Page - 1) * params.Size

	err = db.Model(&t).Where(conditions, args...).Offset(offset).Limit(params.Size).Order(order).Find(&videos).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Model(&t).Where(conditions, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return videos, total, nil
}

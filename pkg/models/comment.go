package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId   uint   `gorm:"comment: 用户 ID"`
	VideoID  uint   `gorm:"comment: 视频 ID"`
	Content  string `gorm:"comment: 评论内容" json:"content"`
	ParentID uint   `gorm:"comment: 父评论 ID" json:"parentId"`
	ISThumb  string `gorm:"comment: 评论-点赞区别;type:enum('ISTHUMB','ISCOMMENT');default:ISTHUMB" json:"isThumb"`
	// DeleteStatus string `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
}

func (c *Comment) TableName() string {
	return "comments"
}

func (t *Comment) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *Comment) Delete(db *gorm.DB, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Error
}

func (t *Comment) Updates(db *gorm.DB, value interface{}, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Updates(value).Error
}

func (t *Comment) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (category *Category, err error) {
	err = db.Model(&t).Where(query, args...).First(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (t *Comment) WhereAll(db *gorm.DB, query interface{}, args ...interface{}) (categories []*Category, err error) {
	err = db.Model(&t).Where(query, args...).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (t *Comment) PageListOrder(db *gorm.DB, order string, params *ListPageInput, conditions interface{}, args ...interface{}) (categories []*Category, total int64, err error) {
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

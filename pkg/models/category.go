package models

import (
	"gorm.io/gorm"
)

type Category struct {
	CreatorBase
	UserID    uint   `json:"userId" gorm:"column:user_id;comment: 用户 ID"`
	Title     string `json:"title" gorm:"column:title;comment: 分类标题"`
	Introduce string `json:"introduce" gorm:"column:introduce;comment: 分类介绍"`
}

func (t *Category) TableName() string {
	return "categories"
}

func (t *Category) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *Category) Delete(db *gorm.DB, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Error
}

func (t *Category) Updates(db *gorm.DB, value interface{}, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Updates(value).Error
}

func (t *Category) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (category *Category, err error) {
	err = db.Model(&t).Where(query, args...).First(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (t *Category) WhereAll(db *gorm.DB, query interface{}, args ...interface{}) (categories []*Category, err error) {
	err = db.Model(&t).Where(query, args...).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (t *Category) PageListOrder(db *gorm.DB, order string, params *ListPageInput, conditions interface{}, args ...interface{}) (categories []*Category, total int64, err error) {
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

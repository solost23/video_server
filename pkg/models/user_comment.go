package models

import "gorm.io/gorm"

type UserComment struct {
	gorm.Model
	UserId  uint `gorm:"comment: 用户 ID"`
	Comment uint `gorm:"comment: 评论 ID"`
}

func (t *UserComment) Insert(db *gorm.DB) (err error) {
	return db.Model(&t).Create(&t).Error
}

func (t *UserComment) Delete(db *gorm.DB, conditions interface{}, args ...interface{}) (err error) {
	return db.Model(&t).Where(conditions, args...).Error
}

func (t *UserComment) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (userComment *UserComment, err error) {
	err = db.Model(&t).Where(query, args...).First(&userComment).Error
	if err != nil {
		return nil, err
	}
	return userComment, nil
}

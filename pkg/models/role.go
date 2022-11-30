package models

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type CasbinRule struct {
	gormadapter.CasbinRule
}

func (c *CasbinRule) TableName() string {
	return "casbin_rule"
}

func (t *CasbinRule) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *CasbinRule) Delete(db *gorm.DB, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Error
}

func (t *CasbinRule) Updates(db *gorm.DB, value interface{}, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Updates(value).Error
}

func (t *CasbinRule) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (casbinModel *CasbinRule, err error) {
	err = db.Model(&t).Where(query, args...).First(&casbinModel).Error
	if err != nil {
		return nil, err
	}
	return casbinModel, nil
}

func (t *CasbinRule) WhereAll(db *gorm.DB, query interface{}, args ...interface{}) (casbinModels []*CasbinRule, err error) {
	err = db.Model(&t).Where(query, args...).Find(&casbinModels).Error
	if err != nil {
		return nil, err
	}
	return casbinModels, nil
}

package models

import (
	"gorm.io/gorm"
)

type CasbinModel struct {
	PType    string `gorm:"column:p_type;default:p"`
	RoleName string `gorm:"column:v0" json:"role_name"`
	Path     string `gorm:"column:v1" json:"path"`
	Method   string `gorm:"column:v2" json:"method"`
}

func (c *CasbinModel) TableName() string {
	return "casbin_rule"
}

func (t *CasbinModel) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *CasbinModel) Delete(db *gorm.DB, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Error
}

func (t *CasbinModel) Updates(db *gorm.DB, value interface{}, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Updates(value).Error
}

func (t *CasbinModel) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (casbinModel *CasbinModel, err error) {
	err = db.Model(&t).Where(query, args...).First(&casbinModel).Error
	if err != nil {
		return nil, err
	}
	return casbinModel, nil
}

func (t *CasbinModel) WhereAll(db *gorm.DB, query interface{}, args ...interface{}) (casbinModels []*CasbinModel, err error) {
	err = db.Model(&t).Where(query, args...).Find(&casbinModels).Error
	if err != nil {
		return nil, err
	}
	return casbinModels, nil
}

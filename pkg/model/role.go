package model

import (
	"gorm.io/gorm"
)

type CasbinModel struct {
	conn     *gorm.DB `gorm:"_"`
	Ptype    string   `gorm:"column:p_type;default:p"`
	RoleName string   `gorm:"column:v0" json:"role_name"`
	Path     string   `gorm:"column:v1" json:"path"`
	Method   string   `gorm:"column:v2" json:"method"`
}

func NewCasbinModel(conn *gorm.DB) *CasbinModel {
	return &CasbinModel{
		conn: conn,
	}
}

func (c *CasbinModel) TableName() string {
	return "casbin_rule"
}

func (c *CasbinModel) Connection() *gorm.DB {
	return c.conn.Table(c.TableName())
}

func (c *CasbinModel) Create(data *CasbinModel) (err error) {
	err = c.Connection().Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CasbinModel) Delete(data *CasbinModel) (casbinModel *CasbinModel, err error) {
	err = c.Connection().Where("v0=? AND v1=? AND v2=?", data.RoleName, data.Path, data.Method).Delete(&casbinModel).Error
	if err != nil {
		return casbinModel, err
	}
	return casbinModel, nil
}

func (c *CasbinModel) FindByRoleName(roleName string) (casbinModels []*CasbinModel, err error) {
	err = DBCasbin.Table(c.TableName()).Where("v0=?", roleName).Find(&casbinModels).Error
	if err != nil {
		return casbinModels, err
	}
	return casbinModels, nil
}

func (c *CasbinModel) Find() (casbinModels []*CasbinModel, err error) {
	err = DBCasbin.Table(c.TableName()).Find(&casbinModels).Error
	if err != nil {
		return casbinModels, err
	}
	return casbinModels, nil
}

func (c *CasbinModel) FindByRolePathMethod(roleName, path, method string) (casbinModel *CasbinModel, err error) {
	err = DBCasbin.Table(c.TableName()).Where("v0 = ? AND v1 = ? AND v2 = ?", roleName, path, method).First(&casbinModel).Error
	if err != nil {
		return casbinModel, err
	}
	return casbinModel, nil
}

package models

import (
	"gorm.io/gorm"
)

var (
	dbConn   *gorm.DB
	dbCasbin *gorm.DB
	err      error
)

func init() {
	dbConn, err = NewMysqlClient(false)
	if err != nil {
		panic(err)
	}
	dbCasbin, err = NewMysqlClient(true)
	// 生成表
	err = autoMigrate(dbConn, &User{}, &Category{}, &CasbinModel{}, &Comment{}, &UserComment{}, &Video{})
	if err != nil {
		panic(err)
	}
	err = autoMigrate(dbCasbin, &CasbinModel{})
	if err != nil {
		panic(err)
	}
}

func autoMigrate(conn *gorm.DB, dst ...interface{}) error {
	return conn.AutoMigrate(dst...)
}

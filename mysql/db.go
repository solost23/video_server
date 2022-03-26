package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetMysqlConn() *gorm.DB {
	return dbConn
}

func GetCasbinConn() *gorm.DB {
	return casbinConn
}

var dbConn *gorm.DB
var casbinConn *gorm.DB

func init() {
	dsn := "root:123@tcp(localhost:3306)/video_server?charset=utf8mb4"
	casbinDsn := "root:123@tcp(localhost:3306)/casbin?charset-utf8mb4"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err.Error())
	}
	casbConn, err := gorm.Open(mysql.Open(casbinDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err.Error())
	}
	dbConn = conn
	casbinConn = casbConn
}

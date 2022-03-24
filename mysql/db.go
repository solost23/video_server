package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetMysqlConn() *gorm.DB {
	return dbConn
}

var dbConn *gorm.DB

func init() {
	dsn := "root:123@tcp(localhost:3306)/video_server?charset=utf8mb4"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	dbConn = conn
}

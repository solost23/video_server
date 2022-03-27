package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"video_server/config"
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
	mysqlClient := config.NewConnections()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", mysqlClient.UserName, mysqlClient.Password, mysqlClient.Host, mysqlClient.Port, mysqlClient.DB, mysqlClient.Charset)
	casbinDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", mysqlClient.UserName, mysqlClient.Password, mysqlClient.Host, mysqlClient.Port, mysqlClient.CasbinDB, mysqlClient.Charset)
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

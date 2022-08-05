package models

import (
	"log"
	"os"
	"strconv"
	"time"
	"video_server/config"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/solost23/tools/mysql"
)

// new mysql conn // flagDb 为false 连接默认库，flagDb为true连接casbin
func NewMysqlClient(flagDb bool) (db *gorm.DB, err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // Log Level
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound错误
			Colorful:                  true,        // 禁用彩色打印
		})
	mysqlClient := config.NewConnections()
	port, err := strconv.Atoi(mysqlClient.Port)
	if err != nil {
		return nil, err
	}
	var dbName = mysqlClient.DB
	if flagDb {
		dbName = mysqlClient.CasbinDB
	}
	mysqlConfig := &mysql.Config{
		UserName: mysqlClient.UserName,
		Password: mysqlClient.Password,
		Host:     mysqlClient.Host,
		Port:     port,
		DB:       dbName,
		Charset:  mysqlClient.Charset,
		Logger:   newLogger,
	}
	db, err = mysql.NewMysqlConnect(mysqlConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}

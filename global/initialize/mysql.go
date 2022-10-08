package initialize

import (
	"github.com/solost23/tools/mysql"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"video_server/global"
)

// new mysql conn // flagDb 为false 连接默认库，flagDb为true连接casbin
func InitMysql(flag bool) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // Log Level
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound错误
			Colorful:                  true,        // 禁用彩色打印
		})
	var dbName = global.ServerConfig.MysqlConfig.DB
	if flag {
		dbName = global.ServerConfig.MysqlConfig.CasbinDB
	}
	addr := strings.Split(global.ServerConfig.MysqlConfig.Addr, ":")
	host := addr[0]
	port := 3306
	var err error
	if len(addr) > 1 {
		port, err = strconv.Atoi(addr[1])
		if err != nil {
			panic(err)
		}
	}
	mysqlConfig := &mysql.Config{
		UserName: global.ServerConfig.MysqlConfig.User,
		Password: global.ServerConfig.MysqlConfig.Password,
		Host:     host,
		Port:     port,
		DB:       dbName,
		Charset:  global.ServerConfig.MysqlConfig.Charset,
		Logger:   newLogger,
	}
	if flag {
		global.CasbinDB, err = mysql.NewMysqlConnect(mysqlConfig)
	} else {
		global.DB, err = mysql.NewMysqlConnect(mysqlConfig)
	}
	if err != nil {
		panic(err)
	}
}

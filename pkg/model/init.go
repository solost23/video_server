package model

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
	"video_server/mysql"
)

var dbConn *gorm.DB

func init() {
	dbConn = mysql.GetMysqlConn()
	dbConn.AutoMigrate(&User{})
	dbConn.AutoMigrate(&Comment{})
	dbConn.AutoMigrate(&Video{})
	dbConn.AutoMigrate(&Class{})
}

const (
	SECRET    = "TY"
	ROLEADMIN = "ADMIN"
	ROLEUSER  = "USER"
)

// 生成唯一UUID
func NewUUID() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Println(err.Error())
		return NewUUID()
	}
	return uuid.String()
}

// 返回当前时间
func GetCurrentTime() int64 {
	return time.Now().Unix()
}

// 给数据库中用户密码进行加密
func NewMd5(str string, salt ...interface{}) string {
	if len(salt) > 0 {
		slice := make([]string, len(salt)+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

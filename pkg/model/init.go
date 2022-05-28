package model

import (
	"crypto/md5"
	"fmt"
	"log"
	"strings"
	"time"
	"video_server/config"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"video_server/mysql"
)

var dbConn *gorm.DB
var DBCasbin *gorm.DB

func init() {
	dbConn = mysql.GetMysqlConn()
	DBCasbin = mysql.GetCasbinConn()
	dbConn.AutoMigrate(&User{})
	dbConn.AutoMigrate(&Comment{})
	dbConn.AutoMigrate(&Video{})
	dbConn.AutoMigrate(&Category{})
	DBCasbin.AutoMigrate(&CasbinModel{})
}

const (
	// 管理员
	ROLEADMIN = "ADMIN"
	ROLEUSER  = "USER"

	// 删除状态
	DELETENORMAL = "DELETE_STATUS_NORMAL"
	DELETEDEL    = "DELETE_STATUS_DEL"

	// 评论类型
	ISTHUMB   = "ISTHUMB"
	ISCOMMENT = "ISCOMMENT"
)

var (
	FilePath = config.Video_path
	SECRET   = config.Md5
)

type PageInfo struct {
	IsNotPage  bool  `json:"-"`
	Page       int32 `json:"page"`
	PageSize   int32 `json:"pageSize"`
	TotalCount int32 `json:"totalCount"`
}

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

// 生成模糊匹配字符串
func LikeFilter(value interface{}) string {
	return fmt.Sprintf("%%%v%%", value)
}

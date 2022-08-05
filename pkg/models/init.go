package models

import (
	"fmt"
	"video_server/config"

	"gorm.io/gorm"
)

var dbConn *gorm.DB
var DBCasbin *gorm.DB

func init() {
	var err error
	dbConn, err = NewMysqlClient(false)
	if err != nil {
		panic(err)
	}
	DBCasbin, err = NewMysqlClient(true)
	if err != nil {
		panic(err)
	}
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
	FilePath   = config.Video_path
	SECRET     = config.Md5
	TimeFormat = "2006-01-02 15:04:05"
)

type PageForm struct {
	UsePage bool `json:"-"`
	Page    int  `json:"page"`
	Size    int  `json:"pageSize"`
}

type PageList struct {
	Size    int   `json:"size"`
	Pages   int64 `json:"pages"`
	Total   int64 `json:"total"`
	Current int   `json:"current"`
}

// 生成模糊匹配字符串
func LikeFilter(value interface{}) string {
	return fmt.Sprintf("%%%v%%", value)
}

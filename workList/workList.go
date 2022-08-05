package workList

import (
	"video_server/pkg/models"

	"gorm.io/gorm"
)

type WorkList struct {
	Conn       *gorm.DB
	ConnCasbin *gorm.DB
}

// 单例模式
func (w *WorkList) GetMysqlConn() (db *gorm.DB) {
	if w.Conn != nil {
		return w.Conn
	}
	var err error
	w.Conn, err = models.NewMysqlClient(false)
	if err != nil {
		panic(err)
	}
	return w.Conn
}

// 单例模式
func (w *WorkList) GetMysqlConnCasbin() (db *gorm.DB) {
	if w.ConnCasbin != nil {
		return w.ConnCasbin
	}
	var err error
	w.ConnCasbin, err = models.NewMysqlClient(true)
	if err != nil {
		panic(err)
	}
	return w.ConnCasbin
}

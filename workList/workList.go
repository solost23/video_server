package workList

import (
	"video_server/pkg/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WorkList struct {
	Conn       *gorm.DB
	ConnCasbin *gorm.DB
	Ctx        *gin.Context
}

func (w *WorkList) Init(ctx *gin.Context) {
	w.Ctx = ctx
}

// 单例模式
func (w *WorkList) GetMysqlConn() (db *gorm.DB) {
	if w.Conn != nil {
		return w.Conn
	}
	var err error
	w.Conn, err = model.NewMysqlClient(false)
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
	w.ConnCasbin, err = model.NewMysqlClient(true)
	if err != nil {
		panic(err)
	}
	return w.ConnCasbin
}

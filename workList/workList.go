package workList

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"video_server/mysql"
)

type WorkList struct {
	Conn *gorm.DB
	Ctx  *gin.Context
}

func NewWorkList(conn *gorm.DB) *WorkList {
	return &WorkList{
		Conn: conn,
	}
}

func (w *WorkList) Init(ctx *gin.Context) {
	w.Ctx = ctx
}

func (w *WorkList) GetMysqlConn() *gorm.DB {
	return mysql.GetMysqlConn()
}

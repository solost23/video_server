package workList

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"video_server/mysql"
)

const response_body = "RESPONSE_BODY"

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

func (w *WorkList) GetMysqlConnCasbin() *gorm.DB {
	return mysql.GetCasbinConn()
}

type ErrCode int

type ApiResponse struct {
	// in: body
	Code    ErrCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceId interface{} `json:"traceId"`
}

// 封装返回
func (w WorkList) RenderJSON(code ErrCode, message string, data interface{}) {
	res := ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	bodyBytes, _ := json.Marshal(res)
	w.Ctx.Set(response_body, string(bodyBytes))
	w.Ctx.JSON(http.StatusOK, res)
}

func (w *WorkList) Render(data interface{}, err error) {
	if err != nil {
		w.RenderJSON(http.StatusOK, err.Error(), data)
	} else {
		w.RenderJSON(http.StatusOK, "success", data)
	}
}

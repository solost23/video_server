package delete

import (
	"errors"
	"github.com/gin-gonic/gin"
	"video_server/pkg/model"
	"video_server/workList"
)

type Action struct {
	workList.WorkList
}

func NewActionWithCtx(ctx *gin.Context) *Action {
	r := &Action{}
	r.Init(ctx)
	return r
}

func (a *Action) Deal(request *Request) (resp Response, err error) {
	if request.ID == "" {
		err = errors.New("request.ID not empty")
		return resp, err
	}
	_, err = model.NewComment(a.GetMysqlConn()).Delete(request.ID)
	if err != nil {
		return resp, err
	}
	return resp, err
}

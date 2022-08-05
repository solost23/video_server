package delete

import (
	"errors"
	"video_server/pkg/models"
	"video_server/workList"

	"github.com/gin-gonic/gin"
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
	_, err = models.NewComment(a.GetMysqlConn()).Delete(request.ID)
	if err != nil {
		return resp, err
	}
	return resp, err
}

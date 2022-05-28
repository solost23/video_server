package detail

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
	// 检查id是否为空
	if request.ID == "" {
		err = errors.New("request.Id not empty")
		return resp, err
	}
	user, err := model.NewUser(a.GetMysqlConn()).FindByID(request.ID)
	if err != nil {
		return resp, err
	}
	// 封装数据 返回
	resp = a.buildResponse(user)
	return resp, err
}

func (a *Action) buildResponse(user *model.User) (resp Response) {
	resp = Response{
		ID:           user.ID,
		UserName:     user.UserName,
		Nickname:     user.Nickname,
		Role:         user.Role,
		Avatar:       user.Avatar,
		Introduce:    user.Introduce,
		FansCount:    user.FansCount,
		CommentCount: user.CommentCount,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
	}
	return resp
}

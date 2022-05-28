package update

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
	if request.UserName == "" {
		err = errors.New("request.UserName not empty")
		return resp, err
	}
	if request.Password == "" {
		err = errors.New("request.Password not empty")
		return resp, err
	}
	// 检查用户是否存在，若用户存在，则更新用户信息
	user, err := model.NewUser(a.GetMysqlConn()).FindByID(request.ID)
	if err != nil {
		return resp, err
	}
	// 删除
	if err := user.Update(a.buildRequest(request)); err != nil {
		return resp, err
	}
	return resp, err
}

func (a *Action) buildRequest(request *Request) (user *model.User) {
	user = &model.User{
		ID:         request.ID,
		UserName:   request.UserName,
		Password:   request.Password,
		Nickname:   request.Nickname,
		Avatar:     request.Avatar,
		Introduce:  request.Introduce,
		UpdateTime: model.GetCurrentTime(),
	}
	return user
}

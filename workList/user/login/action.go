package login

import (
	"errors"
	"github.com/gin-gonic/gin"
	"video_server/pkg/middleware"
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
	if request.UserName == "" {
		err = errors.New("request.UserName not empty")
		return resp, err
	}
	if request.Password == "" {
		err = errors.New("request.Password not empty")
		return resp, err
	}
	// 查看有没有用户，如果没有直接报错
	// 若有，检查账户密码，若有一个为错，则返回
	user, err := model.NewUser(a.GetMysqlConn()).FindBYUserName(request.UserName)
	if err != nil {
		return resp, err
	}
	if request.UserName != user.UserName || model.NewMd5(request.Password, model.SECRET) != user.Password {
		err = errors.New("UserName or Password err")
		return resp, err
	}
	tokenStr, err := middleware.CreateToken(user.UserName, user.Role)
	if err != nil {
		return resp, err
	}
	resp = Response{TokenStr: tokenStr}
	return resp, err
}

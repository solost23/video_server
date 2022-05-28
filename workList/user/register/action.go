package register

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	// 校验参数
	if request.UserName == "" {
		err = errors.New("request.UserName not empty")
		return resp, err
	}
	if request.Password == "" {
		err = errors.New("request.Password not empty")
		return resp, err
	}
	if request.Role != model.ROLEADMIN && request.Role != model.ROLEUSER {
		err = errors.New("request.Role param err")
	}
	// 检查当前用户是否存在，若存在，返回错误
	// 若不存在，创建
	user, err := model.NewUser(a.GetMysqlConn()).FindBYUserName(request.UserName)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, err
		}
	}
	// 说明用户已存在
	if user.ID != "" {
		err = errors.New("user already exist")
		return resp, err
	}
	if err = model.NewUser(a.GetMysqlConn()).Create(a.buildRequest(request)); err != nil {
		return resp, err
	}
	return resp, err
}

func (a *Action) buildRequest(request *Request) (user *model.User) {
	user = &model.User{
		ID:           model.NewUUID(),
		UserName:     request.UserName,
		Password:     model.NewMd5(request.Password, model.SECRET),
		Nickname:     request.Nickname,
		Role:         request.Role,
		Avatar:       request.Avatar,
		Introduce:    request.Introduce,
		FansCount:    0,
		CommentCount: 0,
		CreateTime:   model.GetCurrentTime(),
		UpdateTime:   model.GetCurrentTime(),
	}
	return user
}



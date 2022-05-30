package create

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
	// 先查询用户是否存在，再查询用户下此分类是否存在
	if request.UserId == "" {
		err = errors.New("request.UserId not empty")
		return resp, err
	}
	if request.Title == "" {
		err = errors.New("request.CategoryTitle not empty")
		return resp, err
	}
	_, err = model.NewUser(a.GetMysqlConn()).FindByID(request.UserId)
	if err != nil {
		return resp, err
	}
	// 查询此用户下此分类是否存在，若不存在，则创建
	category, err := model.NewCategory(a.GetMysqlConn()).FindByUserIDClassTitle(request.UserId, request.Title)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, err
		}
	}
	if category.ID != "" {
		return resp, errors.New("用户下此分类已存在")
	}
	if err = model.NewCategory(a.GetMysqlConn()).Create(a.buildRequest(request)); err != nil {
		return resp, err
	}
	return resp, err
}

func (a *Action) buildRequest(request *Request) (category *model.Category) {
	category = &model.Category{
		ID:         model.NewUUID(),
		UserID:     request.UserId,
		Title:      request.Title,
		Introduce:  request.Introduce,
		CreateTime: model.GetCurrentTime(),
		UpdateTime: model.GetCurrentTime(),
	}
	return category
}

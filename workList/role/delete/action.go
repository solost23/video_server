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
	if request.RoleName == "" || request.Path == "" || request.Method == "" {
		err = errors.New("request.RoleName or request.Path or request.Method not empty")
		return resp, err
	}
	// 先查询本条数据是否存在，若存在，删除
	if err = model.NewCasbinModel(a.GetMysqlConnCasbin()).FindByRolePathMethod(request.RoleName, request.Path, request.Method); err != nil {
		return resp, err
	}
	if err = model.NewCasbinModel(a.GetMysqlConnCasbin()).Delete(a.buildRequest(request)); err != nil {
		return resp, err
	}
	return resp, err
}

func (a *Action) buildRequest(request *Request) (casbinModel *model.CasbinModel) {
	casbinModel = &model.CasbinModel{
		RoleName: request.RoleName,
		Path:     request.Path,
		Method:   request.Method,
	}
	return casbinModel
}

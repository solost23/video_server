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
	if request.RoleName == "" || request.Path == "" || request.Method == "" {
		err = errors.New("request.RoleName or request.Path or request.Method not empty")
		return resp, err
	}
	// 先查询本条数据是否存在，若存在，删除
	_, err = models.NewCasbinModel(a.GetMysqlConnCasbin()).FindByRolePathMethod(request.RoleName, request.Path, request.Method)
	if err != nil {
		return resp, err
	}
	_, err = models.NewCasbinModel(a.GetMysqlConnCasbin()).Delete(a.buildRequest(request))
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (a *Action) buildRequest(request *Request) (casbinModel *models.CasbinModel) {
	casbinModel = &models.CasbinModel{
		RoleName: request.RoleName,
		Path:     request.Path,
		Method:   request.Method,
	}
	return casbinModel
}

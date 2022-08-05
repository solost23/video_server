package list

import (
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
	// 通过三个字段筛选内容并分页
	resp.List = []RoleInfo{}
	// 筛选
	casbinModels, total, err := a.FindByFilter(request)
	if err != nil {
		return resp, err
	}
	// 组装数据，返回
	resp.PageInfo = models.PageInfo{
		Page:       request.PageInfo.Page,
		PageSize:   request.PageInfo.PageSize,
		TotalCount: int32(total),
	}
	for _, casbinModel := range casbinModels {
		roleInfo := RoleInfo{
			RoleName: casbinModel.RoleName,
			Path:     casbinModel.Path,
			Method:   casbinModel.Method,
		}
		resp.List = append(resp.List, roleInfo)
	}

	return resp, err
}

func (a *Action) FindByFilter(request *Request) (casbinModels []*models.CasbinModel, total int64, err error) {
	tx := models.NewCasbinModel(a.GetMysqlConnCasbin()).Connection().Select("*")
	if request.Filter.RoleName != "" {
		tx.Where("v0 LIKE ?", models.LikeFilter(request.Filter.RoleName))
	}
	if request.Filter.Path != "" {
		tx.Where("v1 LIKE ?", models.LikeFilter(request.Filter.Path))
	}
	if request.Filter.Method != "" {
		tx.Where("v2 LIKE ?", models.LikeFilter(request.Filter.Method))
	}
	tx.Count(&total)
	// 数据分页
	if request.PageInfo == nil {
		request.PageInfo = &models.PageInfo{
			Page:     1,
			PageSize: 10,
		}
	}
	err = tx.Offset(int((request.PageInfo.Page - 1) * request.PageInfo.PageSize)).Limit(int(request.PageInfo.PageSize)).Find(&casbinModels).Error
	return casbinModels, total, err
}

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
	resp.List = []CategoryInfo{}
	// 筛选
	categorys, total, err := a.FindByFilter(request)
	if err != nil {
		return resp, err
	}
	// 组装数据，返回
	resp.PageInfo = models.PageInfo{
		Page:       request.PageInfo.Page,
		PageSize:   request.PageInfo.PageSize,
		TotalCount: int32(total),
	}
	for _, category := range categorys {
		categoryInfo := CategoryInfo{
			ID:         category.ID,
			UserID:     category.UserID,
			Title:      category.Title,
			Introduce:  category.Introduce,
			CreateTime: category.CreateTime,
			UpdateTime: category.UpdateTime,
		}
		resp.List = append(resp.List, categoryInfo)
	}
	return resp, err
}

func (a *Action) FindByFilter(request *Request) (categorys []*models.Category, total int64, err error) {
	tx := models.NewCategory(a.GetMysqlConn()).Connection().Select("*")
	if request.Filter.UserID != "" {
		tx.Where("user_id = ?", request.Filter.UserID)
	}
	tx.Count(&total)
	// 数据分页
	if request.PageInfo == nil {
		request.PageInfo = &models.PageInfo{
			Page:     1,
			PageSize: 10,
		}
	}
	err = tx.Offset(int((request.PageInfo.Page - 1) * request.PageInfo.PageSize)).Limit(int(request.PageInfo.PageSize)).Find(&categorys).Error
	return categorys, total, err
}

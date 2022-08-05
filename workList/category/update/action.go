package update

import (
	"errors"
	"video_server/pkg/models"
	"video_server/workList"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	// 校验用户是否存在，校验分类是否存在
	if request.UserId == "" {
		err = errors.New("request.UserId not empty")
		return resp, err
	}
	if request.CategoryId == "" {
		err = errors.New("request.CategoryId not empty")
	}
	_, err = models.NewUser(a.GetMysqlConn()).FindByID(request.UserId)
	if err != nil {
		return resp, err
	}
	category, err := models.NewCategory(a.GetMysqlConn()).FindByID(request.CategoryId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, err
		}
	}
	if category.ID == "" {
		return resp, errors.New("用户下此分类不存在")
	}
	if err = models.NewCategory(a.GetMysqlConn()).Update(a.buildRequest(request, category.ID)); err != nil {
		return resp, err
	}
	return resp, err
}

func (a *Action) buildRequest(request *Request, categoryID string) (category *models.Category) {
	category = &models.Category{
		ID:         categoryID,
		UserID:     request.UserId,
		Title:      request.Title,
		Introduce:  request.Introduce,
		UpdateTime: models.GetCurrentTime(),
	}
	return category
}

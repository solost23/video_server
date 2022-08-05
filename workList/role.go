package workList

import (
	"errors"
	"strings"
	"video_server/forms"
	"video_server/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleService struct {
	WorkList
}

func (w *RoleService) Insert(c *gin.Context, params *forms.RoleInsertForm) (err error) {
	// 查询数据是否存在
	db := w.GetMysqlConn()

	query := []string{"role_name = ?", "path = ?", "method = ?"}
	args := []interface{}{params.RoleName, params.Path, params.Method}
	_, err = (&models.CasbinModel{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err == nil {
		return errors.New("权限存在")
	}
	err = (&models.CasbinModel{
		RoleName: params.RoleName,
		Path:     params.Path,
		Method:   params.Method,
	}).Insert(db)
	if err != nil {
		return err
	}
	return nil
}

func (w *RoleService) Delete(c *gin.Context, params *forms.RoleInsertForm) (err error) {
	db := w.GetMysqlConn()

	query := []string{"role_name = ?", "path = ?", "method = ?"}
	args := []interface{}{params.RoleName, params.Path, params.Method}
	_, err = (&models.CasbinModel{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	err = (&models.CasbinModel{}).Delete(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	return nil
}

func (w *RoleService) List(c *gin.Context, params *forms.RoleListForm) (response *forms.RoleListResponse, err error) {
	db := w.GetMysqlConn()

	query := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)
	if params.RoleName != "" {
		query = append(query, "role_name LIKE ?")
		args = append(args, models.LikeFilter(params.RoleName))
	}
	if params.Path != "" {
		query = append(query, "path LIKE ?")
		args = append(args, models.LikeFilter(params.Page))
	}
	if params.Method != "" {
		query = append(query, "method LIKE ?")
		args = append(args, models.LikeFilter(params.Method))
	}
	casbinModels, err := (&models.CasbinModel{}).WhereAll(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	records := make([]forms.RoleInsertForm, 0, len(casbinModels))
	for _, casbinModel := range casbinModels {
		records = append(records, forms.RoleInsertForm{
			RoleName: casbinModel.RoleName,
			Method:   casbinModel.Method,
			Path:     casbinModel.Path,
		})
	}
	response = &forms.RoleListResponse{
		Records: records,
	}
	return response, nil
}

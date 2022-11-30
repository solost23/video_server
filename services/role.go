package services

import (
	"errors"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"video_server/forms"
	"video_server/global"
	"video_server/pkg/models"
)

func (s *Service) InsertRole(c *gin.Context, params *forms.RoleInsertForm) (err error) {
	// 查询数据是否存在
	db := global.DB

	query := []string{"v0 = ?", "v1 = ?", "v2 = ?"}
	args := []interface{}{params.RoleName, params.Path, params.Method}
	_, err = (&models.CasbinRule{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err == nil {
		return errors.New("权限存在")
	}
	err = (&models.CasbinRule{
		CasbinRule: gormadapter.CasbinRule{
			Ptype: "p",
			V0:    params.RoleName,
			V1:    params.Path,
			V2:    params.Method,
		},
	}).Insert(db)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteRole(c *gin.Context, params *forms.RoleInsertForm) (err error) {
	db := global.DB

	query := []string{"v0 = ?", "v1 = ?", "v2 = ?"}
	args := []interface{}{params.RoleName, params.Path, params.Method}
	_, err = (&models.CasbinRule{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	err = (&models.CasbinRule{}).Delete(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ListRole(c *gin.Context, params *forms.RoleListForm) (response *forms.RoleListResponse, err error) {
	db := global.DB

	query := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)
	if params.RoleName != "" {
		query = append(query, "v0 LIKE ?")
		args = append(args, models.LikeFilter(params.RoleName))
	}
	if params.Path != "" {
		query = append(query, "v1 LIKE ?")
		args = append(args, models.LikeFilter(params.Page))
	}
	if params.Method != "" {
		query = append(query, "v2 LIKE ?")
		args = append(args, models.LikeFilter(params.Method))
	}
	casbinModels, err := (&models.CasbinRule{}).WhereAll(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	records := make([]forms.RoleInsertForm, 0, len(casbinModels))
	for _, casbinModel := range casbinModels {
		records = append(records, forms.RoleInsertForm{
			RoleName: casbinModel.V0,
			Method:   casbinModel.V1,
			Path:     casbinModel.V2,
		})
	}
	response = &forms.RoleListResponse{
		Records: records,
	}
	return response, nil
}

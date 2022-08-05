package workList

import (
	"math"
	"strings"
	"video_server/forms"
	"video_server/pkg/models"
	"video_server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CategoryService struct {
	WorkList
}

func (w *CategoryService) Insert(c *gin.Context, params *forms.CategoryInsertForm) (err error) {
	db := w.GetMysqlConn()
	// 查询用户是否存在，若存在，则新建分类
	query := []string{"id = ?"}
	args := []interface{}{params.UserId}
	_, err = (&models.User{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	// 增加分类
	err = (&models.Category{
		UserID:    params.UserId,
		Title:     params.Title,
		Introduce: params.Introduce,
	}).Insert(db)
	if err != nil {
		return err
	}
	return nil
}

func (w *CategoryService) List(c *gin.Context, params *forms.CategoryListForm) (response *forms.CategoryListResponse, err error) {
	db := w.GetMysqlConn()
	query := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)
	if params.UserID > 0 {
		query = append(query, "user_id = ?")
		args = append(args, params.UserID)
	}
	if params.Title != "" {
		query = append(query, "title LIKE ?")
		args = append(args, models.LikeFilter(params.Title))
	}
	if params.Introduce != "" {
		query = append(query, "introduce LIKE ?")
		args = append(args, models.LikeFilter(params.Introduce))
	}
	categories, total, err := (&models.Category{}).PageListOrder(db, "", &models.ListPageInput{
		Page: params.Page,
		Size: params.Size,
	}, strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	// 封装数据
	records := make([]forms.CategoryListRecord, 0, len(categories))
	for _, category := range categories {
		records = append(records, forms.CategoryListRecord{
			Title:     category.Title,
			Introduce: category.Introduce,
		})
	}
	response = &forms.CategoryListResponse{
		Records: records,
		PageList: &utils.PageList{
			Size:    params.Size,
			Pages:   int64(math.Ceil(float64(total) / float64(params.Size))),
			Total:   total,
			Current: params.Page,
		},
	}
	return response, nil
}

func (w *CategoryService) Update(c *gin.Context, id uint, params *forms.CategoryUpdateForm) (err error) {
	// base logic: 查看用户是否存在，若存在，更新数据，否则错误
	db := w.GetMysqlConn()

	query := []string{"id = ?"}
	args := []interface{}{id}

	_, err = (&models.Category{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	value := map[string]interface{}{
		"title":     params.Title,
		"introduce": params.Introduce,
	}
	err = (&models.Category{}).Updates(db, value, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	return nil
}

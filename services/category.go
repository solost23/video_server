package services

import (
	"math"
	"strconv"
	"strings"
	"video_server/forms"
	"video_server/global"
	"video_server/pkg/constants"
	"video_server/pkg/models"
	"video_server/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (s *Service) InsertCategory(c *gin.Context, params *forms.CategoryInsertForm) (err error) {
	db := global.DB
	user := utils.GetUser(c)

	introduce := ""
	if params.Introduce != nil {
		introduce = *params.Introduce
	}
	category := &models.Category{
		CreatorBase: models.CreatorBase{
			CreatorId: user.ID,
		},
		Title:     *params.Title,
		Introduce: introduce,
	}
	if err = models.GInsert(db, category); err != nil {
		return err
	}
	z := NewZinc()
	err = z.InsertDocument(c, constants.ZINCINDEXCATEGORY, strconv.Itoa(int(user.ID)), map[string]interface{}{
		"title":     category.Title,
		"introduce": category.Introduce,
	})
	if err != nil {
		return
	}
	return nil
}

func (s *Service) ListCategory(c *gin.Context, params *forms.CategoryListForm) (response *forms.CategoryListResponse, err error) {
	db := global.DB

	query := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)
	if params.UserID != nil && *params.UserID > 0 {
		query = append(query, "creator_id = ?")
		args = append(args, params.UserID)
	}
	if params.Title != nil && *params.Title != "" {
		query = append(query, "title LIKE ?")
		args = append(args, models.LikeFilter(params.Title))
	}
	if params.Introduce != nil && *params.Introduce != "" {
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
	records := make([]*forms.CategoryListRecord, 0, len(categories))
	for _, category := range categories {
		id := category.ID
		title := category.Title
		introduce := category.Introduce
		records = append(records, &forms.CategoryListRecord{
			Id:        &id,
			Title:     &title,
			Introduce: &introduce,
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

func (s *Service) UpdateCategory(c *gin.Context, id uint, params *forms.CategoryUpdateForm) (err error) {
	// base logic: 查看用户是否存在，若存在，更新数据，否则错误
	db := global.DB
	user := utils.GetUser(c)

	query := []string{"id = ?", "creator_id = ?"}
	args := []interface{}{id, user.ID}

	sqlCategory, err := (&models.Category{}).WhereOne(db, strings.Join(query, " AND "), args...)
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
	if err != nil {
		return err
	}
	z := NewZinc()
	err = z.DeleteDocument(c, constants.ZINCINDEXCATEGORY, strconv.Itoa(int(user.ID)))
	if err != nil {
		return err
	}
	err = z.InsertDocument(c, constants.ZINCINDEXCATEGORY, strconv.Itoa(int(user.ID)), map[string]interface{}{
		"title":     sqlCategory.Title,
		"introduce": sqlCategory.Introduce,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SearchCategory(c *gin.Context, params *forms.SearchForm) (*forms.CategoryListResponse, error) {
	// 直接搜索
	z := NewZinc()
	from := int32((params.Page - 1) * params.Size)
	size := from + int32(params.Size) - 1
	searchResults, total, err := z.SearchDocument(c, constants.ZINCINDEXCATEGORY, *params.Keyword, from, size)
	if err != nil {
		return nil, err
	}
	categoryIds := make([]string, 0, len(searchResults))
	for _, searchResult := range searchResults {
		categoryIds = append(categoryIds, *searchResult.Id)
	}
	sqlCategories, err := models.GWhereAllSelectOrder(global.DB, &models.Category{}, "*", "id DESC", "id IN ?", categoryIds)
	if err != nil {
		return nil, err
	}

	records := make([]*forms.CategoryListRecord, 0, len(searchResults))
	for _, sqlCategory := range sqlCategories {
		id := sqlCategory.ID
		title := sqlCategory.Title
		introduce := sqlCategory.Introduce
		records = append(records, &forms.CategoryListRecord{
			Id:        &id,
			Title:     &title,
			Introduce: &introduce,
		})
	}
	result := &forms.CategoryListResponse{
		Records: records,
		PageList: &utils.PageList{
			Size:    params.Size,
			Pages:   int64(math.Ceil(float64(total) / float64(params.Size))),
			Total:   total,
			Current: params.Page,
		},
	}
	return result, nil
}

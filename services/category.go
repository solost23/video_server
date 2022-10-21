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

	// 查询用户是否存在，若存在，则新建分类
	query := []string{"id = ?"}
	args := []interface{}{user.ID}
	_, err = (&models.User{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	// 增加分类
	category := &models.Category{
		UserID:    user.ID,
		Title:     params.Title,
		Introduce: params.Introduce,
	}
	err = category.Insert(db)
	if err != nil {
		return err
	}
	z := &Zinc{Username: global.ServerConfig.ZincConfig.Username, Password: global.ServerConfig.ZincConfig.Password}
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
			Id:        category.ID,
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

func (s *Service) UpdateCategory(c *gin.Context, id uint, params *forms.CategoryUpdateForm) (err error) {
	// base logic: 查看用户是否存在，若存在，更新数据，否则错误
	db := global.DB

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
	user, err := (&models.Category{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	z := &Zinc{Username: global.ServerConfig.ZincConfig.Username, Password: global.ServerConfig.ZincConfig.Password}
	err = z.DeleteDocument(c, constants.ZINCINDEXCATEGORY, strconv.Itoa(int(user.ID)))
	if err != nil {
		return err
	}
	err = z.InsertDocument(c, constants.ZINCINDEXCATEGORY, strconv.Itoa(int(user.ID)), map[string]interface{}{
		"title":     user.Title,
		"introduce": user.Introduce,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SearchCategory(c *gin.Context, params *forms.SearchForm) (*forms.CategoryListResponse, error) {
	// 直接搜索
	z := &Zinc{Username: global.ServerConfig.ZincConfig.Username, Password: global.ServerConfig.ZincConfig.Password}
	from := int32((params.Page - 1) * params.Size)
	size := from + int32(params.Size) - 1
	searchResults, total, err := z.SearchDocument(c, constants.ZINCINDEXCATEGORY, params.Keyword, from, size)
	if err != nil {
		return nil, err
	}
	records := make([]forms.CategoryListRecord, 0, len(searchResults))
	for _, searchResult := range searchResults {
		id, _ := strconv.Atoi(*searchResult.Id)
		records = append(records, forms.CategoryListRecord{
			Id:        uint(id),
			Title:     searchResult.Source["title"].(string),
			Introduce: searchResult.Source["introduce"].(string),
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

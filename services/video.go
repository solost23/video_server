package services

import (
	"errors"
	"math"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
	"video_server/forms"
	"video_server/global"
	"video_server/pkg/constants"
	"video_server/pkg/models"
	"video_server/pkg/utils"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func (s *Service) VideoList(c *gin.Context, params *forms.VideoListForm) (response *forms.VideoListResponse, err error) {
	db := global.DB

	categoryIds := make([]uint, 0)
	if params.CategoryId > 0 {
		categoryIds = append(categoryIds, params.CategoryId)
	}
	if params.CategoryName != "" {
		query := []string{"title LIKE ?"}
		args := []interface{}{models.LikeFilter(params.CategoryName)}
		categories, err := (&models.Category{}).WhereAll(db, strings.Join(query, " AND "), args...)
		if err != nil {
			return nil, err
		}
		for _, category := range categories {
			categoryIds = append(categoryIds, category.ID)
		}
	}
	userIds := make([]uint, 0)
	if params.UserName != "" {
		query := []string{"user_name LIKE ?"}
		args := []interface{}{models.LikeFilter(params.UserName)}
		users, err := (&models.User{}).WhereAll(db, strings.Join(query, " AND "), args...)
		if err != nil {
			return nil, err
		}
		for _, user := range users {
			userIds = append(userIds, user.ID)
		}
	}

	query := make([]string, 0, 4)
	args := make([]interface{}, 0, 4)
	if len(categoryIds) > 0 {
		query = append(query, "category_id IN ?")
		args = append(args, categoryIds)
	}
	if len(userIds) > 0 {
		query = append(query, "creator_id IN ?")
		args = append(args, userIds)
	}
	if params.VideoTitle != "" {
		query = append(query, "title LIKE ?")
		args = append(args, models.LikeFilter(params.VideoTitle))
	}
	if params.Introduce != "" {
		query = append(query, "introduce LIKE ?")
		args = append(args, models.LikeFilter(params.Introduce))
	}
	videos, total, err := (&models.Video{}).PageListOrder(db, "", &models.ListPageInput{Page: params.Page, Size: params.Size}, strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	videoIds := make([]uint, 0, len(videos))
	categoryIds = make([]uint, 0, len(videos))
	userIds = make([]uint, 0, len(videos))
	for _, video := range videos {
		videoIds = append(videoIds, video.ID)
		categoryIds = append(categoryIds, video.CategoryId)
		userIds = append(userIds, video.UserId)
	}
	// 求出所有所属人和所属类别
	userNameMaps := make(map[uint]string)
	categoryNameMaps := make(map[uint]string)
	query = []string{"id IN ?"}
	args = []interface{}{userIds}
	users, err := (&models.User{}).WhereAll(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	query = []string{"id IN ?"}
	args = []interface{}{categoryIds}
	categories, err := (&models.Category{}).WhereAll(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		userNameMaps[user.ID] = user.Username
	}
	for _, category := range categories {
		categoryNameMaps[category.ID] = category.Title
	}
	// 求出视频点赞数和评论数
	thumbCountMap := make(map[uint]uint)
	commentCountMap := make(map[uint]uint)
	query = []string{"video_id IN ?", "is_thumb = ?"}
	args = []interface{}{videoIds, "ISTHUMB"}
	thumbs, err := (&models.Comment{}).WhereCountGroup(db, "", strings.Join(query, " AND "))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	query = []string{"video_id IN ?", "is_thumb = ?"}
	args = []interface{}{videoIds, "ISCOMMENT"}
	comments, err := (&models.Comment{}).WhereCountGroup(db, "", strings.Join(query, " AND "))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	for _, thumb := range thumbs {
		thumbCountMap[thumb.VideoId] = thumb.CommentCount
	}
	for _, comment := range comments {
		commentCountMap[comment.VideoId] = comment.CommentCount
	}
	// 封装数据返回
	records := make([]forms.VideoListRecord, 0, len(videos))
	for _, video := range videos {
		records = append(records, forms.VideoListRecord{
			ID:           video.ID,
			UserID:       video.UserId,
			UserName:     userNameMaps[video.UserId],
			CategoryID:   video.CategoryId,
			CategoryName: categoryNameMaps[video.CategoryId],
			Title:        video.Title,
			Introduce:    video.Introduce,
			ImageUrl:     utils.FulfillImageOSSPrefix(video.ImageUrl),
			VideoUrl:     utils.FulfillImageOSSPrefix(video.VideoUrl),
			ThumbCount:   int64(thumbCountMap[video.ID]),
			CommentCount: int64(commentCountMap[video.ID]),
			CreatedAt:    video.CreatedAt.Format(constants.DateTime),
			UpdatedAt:    video.UpdatedAt.Format(constants.DateTime),
		})
	}
	response = &forms.VideoListResponse{
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

func (s *Service) SearchVideo(c *gin.Context, params *forms.SearchForm) (*forms.VideoListResponse, error) {
	db := global.DB

	keyword := "*"
	if params.Keyword != nil && *params.Keyword != "" {
		keyword = *params.Keyword
	}
	z := NewZinc()
	from := int32((params.Page - 1) * params.Size)
	size := from + int32(params.Size) - 1
	searchResults, total, err := z.SearchDocument(c, constants.ZINCINDEXVIDEO, keyword, from, size)
	if err != nil {
		return nil, err
	}
	userIds := make([]uint, 0, len(searchResults))
	videoIds := make([]uint, 0, len(searchResults))
	categoryIds := make([]uint, 0, len(searchResults))
	for _, searchResult := range searchResults {
		userId := uint(searchResult.Source["user_id"].(float64))
		videoId, _ := strconv.Atoi(*searchResult.Id)
		categoryId := uint(searchResult.Source["category_id"].(float64))
		userIds = append(userIds, userId)
		videoIds = append(videoIds, uint(videoId))
		categoryIds = append(categoryIds, categoryId)
	}
	users, err := (&models.User{}).WhereAll(db, "id IN ?", userIds)
	if err != nil {
		return nil, err
	}
	userIdToUsernameMaps := make(map[uint]string, len(users))
	for _, user := range users {
		userIdToUsernameMaps[user.ID] = user.Username
	}
	categories, err := (&models.Category{}).WhereAll(db, "id IN ?", categoryIds)
	if err != nil {
		return nil, err
	}
	categoryIdToNameMaps := make(map[uint]string, len(categories))
	for _, category := range categories {
		categoryIdToNameMaps[category.ID] = category.Title
	}
	videos, err := (&models.Video{}).WhereAll(db, "id IN ?", videoIds)
	if err != nil {
		return nil, err
	}
	videoIdToVideoInfoMaps := make(map[uint]struct {
		ImageUrl     string
		VideoUrl     string
		ThumbCount   int64
		CommentCount int64
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}, len(videos))
	for _, video := range videos {
		videoIdToVideoInfoMaps[video.ID] = struct {
			ImageUrl     string
			VideoUrl     string
			ThumbCount   int64
			CommentCount int64
			CreatedAt    time.Time
			UpdatedAt    time.Time
		}{ImageUrl: video.ImageUrl, VideoUrl: video.VideoUrl, ThumbCount: video.ThumbCount, CommentCount: video.CommentCount, CreatedAt: video.CreatedAt, UpdatedAt: video.UpdatedAt}
	}

	records := make([]forms.VideoListRecord, 0, len(searchResults))
	for _, searchResult := range searchResults {
		videoId, _ := strconv.Atoi(*searchResult.Id)
		userId := uint(searchResult.Source["user_id"].(float64))
		categoryId := uint(searchResult.Source["category_id"].(float64))
		records = append(records, forms.VideoListRecord{
			ID:           uint(videoId),
			UserID:       userId,
			UserName:     userIdToUsernameMaps[userId],
			CategoryID:   categoryId,
			CategoryName: categoryIdToNameMaps[categoryId],
			Title:        searchResult.Source["title"].(string),
			Introduce:    searchResult.Source["introduce"].(string),
			ImageUrl:     utils.FulfillImageOSSPrefix(videoIdToVideoInfoMaps[uint(videoId)].ImageUrl),
			VideoUrl:     utils.FulfillImageOSSPrefix(videoIdToVideoInfoMaps[uint(videoId)].VideoUrl),
			ThumbCount:   videoIdToVideoInfoMaps[uint(videoId)].ThumbCount,
			CommentCount: videoIdToVideoInfoMaps[uint(videoId)].CommentCount,
			CreatedAt:    videoIdToVideoInfoMaps[uint(videoId)].CreatedAt.Format(constants.DateTime),
			UpdatedAt:    videoIdToVideoInfoMaps[uint(videoId)].UpdatedAt.Format(constants.DateTime),
		})
	}
	result := &forms.VideoListResponse{
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

func (s *Service) VideoDetail(c *gin.Context, id uint) (response *forms.VideoListRecord, err error) {
	db := global.DB

	query := []string{"id = ?"}
	args := []interface{}{id}
	video, err := (&models.Video{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	// 统计点赞数和评论数
	query = []string{"video_id = ?", "is_thumb = ?"}
	args = []interface{}{video.ID, "ISTHUMB"}
	thumbs, err := (&models.Comment{}).WhereCountGroup(db, "", strings.Join(query, " AND "))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	query = []string{"video_id = ?", "is_thumb = ?"}
	args = []interface{}{video.ID, "ISCOMMENT"}
	comments, err := (&models.Comment{}).WhereCountGroup(db, "", strings.Join(query, " AND "))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	var thumbCount, commentCount int64
	if len(thumbs) > 0 {
		thumbCount = int64(thumbs[0].CommentCount)
	}
	if len(comments) > 0 {
		commentCount = int64(comments[0].CommentCount)
	}
	// 封装数据返回
	response = &forms.VideoListRecord{
		ID:           video.ID,
		UserID:       video.UserId,
		CategoryID:   video.CategoryId,
		Title:        video.Title,
		Introduce:    video.Introduce,
		ImageUrl:     utils.FulfillImageOSSPrefix(video.ImageUrl),
		VideoUrl:     utils.FulfillImageOSSPrefix(video.VideoUrl),
		ThumbCount:   thumbCount,
		CommentCount: commentCount,
		CreatedAt:    video.CreatedAt.Format(constants.DateTime),
		UpdatedAt:    video.UpdatedAt.Format(constants.DateTime),
	}
	return response, nil
}

func (s *Service) VideoDelete(c *gin.Context, id uint) (err error) {
	// base logic: 删视频，删评论
	db := global.DB
	tx := db.Begin()

	query := []string{"id = ?"}
	args := []interface{}{id}
	_, err = (&models.Video{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	err = (&models.Video{}).Delete(tx, strings.Join(query, " AND "), args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = []string{"video_id = ?"}
	err = (&models.Comment{}).Delete(tx, strings.Join(query, " AND "), args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	_ = tx.Commit().Error
	z := NewZinc()
	err = z.DeleteDocument(c, constants.ZINCINDEXVIDEO, strconv.Itoa(int(id)))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) VideoInsert(c *gin.Context, params *forms.VideoInsertForm) (id uint, err error) {
	db := global.DB
	user := utils.GetUser(c)

	// base logic: 查看分类是否存在，若存在，则创建视频
	query := []string{"id = ?"}
	args := []interface{}{params.CategoryId}
	_, err = (&models.Category{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return 0, err
	}
	video := &models.Video{
		UserId:       user.ID,
		CategoryId:   params.CategoryId,
		Title:        params.Title,
		Introduce:    params.Introduce,
		ImageUrl:     utils.TrimDomainPrefix(params.ImageUrl),
		VideoUrl:     utils.TrimDomainPrefix(params.VideoUrl),
		ThumbCount:   0,
		CommentCount: 0,
	}
	err = video.Insert(db)
	if err != nil {
		return 0, err
	}
	z := NewZinc()
	err = z.InsertDocument(c, constants.ZINCINDEXVIDEO, strconv.Itoa(int(video.ID)), map[string]interface{}{
		"user_id":     video.UserId,
		"category_id": video.CategoryId,
		"title":       video.Title,
		"introduce":   video.Introduce,
	})
	if err != nil {
		return 0, err
	}
	return video.ID, nil
}

func (s *Service) VideoUploadImg(c *gin.Context, file *multipart.FileHeader) (result string, err error) {
	user := utils.GetUser(c)

	folder := "video.server.videos.img"

	url, err := UploadImg(user.ID, folder, file.Filename, file, "image")
	if err != nil {
		return "", err
	}

	return utils.FulfillImageOSSPrefix(utils.TrimDomainPrefix(url)), nil
}

func (s *Service) VideoUploadVid(c *gin.Context, file *multipart.FileHeader) (result string, err error) {
	user := utils.GetUser(c)

	folder := "video.server.videos.vid"

	url, err := UploadVid(user.ID, folder, file.Filename, file, "video")
	if err != nil {
		return "", err
	}
	return utils.FulfillImageOSSPrefix(utils.TrimDomainPrefix(url)), nil
}

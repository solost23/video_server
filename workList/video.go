package workList

import (
	"errors"
	"math"
	"mime/multipart"
	"strings"
	"video_server/forms"
	"video_server/pkg/constants"
	"video_server/pkg/models"
	"video_server/pkg/utils"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type VideoService struct {
	WorkList
}

func (w *VideoService) VideoList(c *gin.Context, params *forms.VideoListForm) (response *forms.VideoListResponse, err error) {
	db := w.GetMysqlConn()

	categoryIds := make([]uint, 0)
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
		query = append(query, "user_id IN ?")
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
		userNameMaps[user.ID] = user.UserName
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
			ImageUrl:     video.ImageUrl,
			VideoUrl:     video.VideoUrl,
			ThumbCount:   int64(thumbCountMap[video.ID]),
			CommentCount: int64(commentCountMap[video.ID]),
			CreatedAt:    video.CreatedAt.Format(constants.TimeFormat),
			UpdatedAt:    video.UpdatedAt.Format(constants.TimeFormat),
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

func (w *VideoService) VideoDetail(c *gin.Context, id uint) (response *forms.VideoListRecord, err error) {
	db := w.GetMysqlConn()

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
		ImageUrl:     video.ImageUrl,
		VideoUrl:     video.VideoUrl,
		ThumbCount:   thumbCount,
		CommentCount: commentCount,
		CreatedAt:    video.CreatedAt.Format(constants.TimeFormat),
		UpdatedAt:    video.UpdatedAt.Format(constants.TimeFormat),
	}
	return response, nil
}

func (w *VideoService) VideoDelete(c *gin.Context, id uint) (err error) {
	// base logic: 删视频，删评论
	db := w.GetMysqlConn()
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
	return nil
}

func (w *VideoService) VideoInsert(c *gin.Context, params *forms.VideoInsertForm) (err error) {
	db := w.GetMysqlConn()
	user := utils.GetUser(c)

	// base logic: 查看分类是否存在，若存在，则创建视频
	query := []string{"id = ?"}
	args := []interface{}{params.CategoryId}
	_, err = (&models.Category{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	err = (&models.Video{
		UserId:       user.ID,
		CategoryId:   params.CategoryId,
		Title:        params.Title,
		Introduce:    params.Introduce,
		ImageUrl:     params.ImageUrl,
		VideoUrl:     params.VideoUrl,
		ThumbCount:   0,
		CommentCount: 0,
	}).Insert(db)
	if err != nil {
		return err
	}
	return nil
}

func (w *VideoService) VideoUploadImg(c *gin.Context, file *multipart.FileHeader) (result string, err error) {
	user := utils.GetUser(c)

	result, err = UploadImg(user, "img", file)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (w *VideoService) VideoUploadVid(c *gin.Context, file *multipart.FileHeader) (result string, err error) {
	user := utils.GetUser(c)

	result, err = UploadVid(user, "vid", file)
	if err != nil {
		return "", err
	}

	return result, nil
}

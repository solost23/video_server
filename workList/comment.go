package workList

import (
	"errors"
	"math"
	"strings"
	"video_server/forms"
	"video_server/pkg/models"
	"video_server/pkg/utils"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type CommentService struct {
	WorkList
}

func (w *CommentService) CommentInsert(c *gin.Context, params *forms.CommentCreateForm) (err error) {
	// 直接创建,关系表也要创建
	db := w.GetMysqlConn()
	tx := db.Begin()
	user := utils.GetUser(c)

	// base logic: 查看video_id是否存在，查询是否有父评论，若无，直接存储，若有，parent_id = 父评论 ID 存储
	query := []string{"id = ?"}
	args := []interface{}{params.VideoID}
	_, err = (&models.Video{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	query = append(query, "is_thumb = ?")
	args = []interface{}{params.ParentID, params.ISThumb}
	comment, err := (&models.Comment{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	var commentData *models.Comment
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在
		commentData = &models.Comment{
			UserId:   user.ID,
			VideoId:  params.VideoID,
			Content:  params.Content,
			ParentId: 0,
			ISThumb:  params.ISThumb,
		}
		err = commentData.Insert(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// 存在
		commentData = &models.Comment{
			UserId:   user.ID,
			VideoId:  params.VideoID,
			Content:  params.Content,
			ParentId: comment.ID,
			ISThumb:  params.ISThumb,
		}
		err = commentData.Insert(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// 存储关系
	err = (&models.UserComment{
		UserId:  user.ID,
		Comment: commentData.ID,
	}).Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (w *CommentService) CommentDelete(c *gin.Context, id uint) (err error) {
	// base logic: 查看关系表中有无此用户和评论对应关系，如果有，那么去评论表删除数据
	db := w.GetMysqlConn()
	tx := db.Begin()
	user := utils.GetUser(c)

	query := []string{"user_id = ?", "comment_id = ?"}
	args := []interface{}{user.ID, id}
	_, err = (&models.UserComment{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	err = (&models.Comment{}).Delete(tx, strings.Join(query, " AND "), args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = (&models.UserComment{}).Delete(tx, strings.Join(query, " AND "), args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (w *CommentService) CommentList(c *gin.Context, params *forms.CommentListForm) (response *forms.CommentListResponse, err error) {
	db := w.GetMysqlConn()

	query := make([]string, 0, 1)
	args := make([]interface{}, 0, 1)
	if params.VideoId > 0 {
		query = append(query, "video_id = ?")
		args = append(args, params.VideoId)
	}
	comments, total, err := (&models.Comment{}).PageListOrder(db, "", &models.ListPageInput{Page: params.Page, Size: params.Size}, strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	// 封装数据，返回
	records := make([]forms.CommentListRecord, 0, len(comments))
	for _, comment := range comments {
		records = append(records, forms.CommentListRecord{
			Id:          comment.ID,
			Content:     comment.Content,
			ParentId:    comment.ParentId,
			ISThumb:     comment.ISThumb,
			CreatedAt:   comment.CreatedAt.Format(models.TimeFormat),
			UpdatedTime: comment.UpdatedAt.Format(models.TimeFormat),
		})
	}

	response = &forms.CommentListResponse{
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

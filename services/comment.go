package services

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"video/forms"
	"video/global"
	"video/pkg/constants"
	"video/pkg/models"
	"video/pkg/utils"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func (s *Service) CommentInsert(c *gin.Context, params *forms.CommentCreateForm) (err error) {
	// 直接创建,关系表也要创建
	db := global.DB
	tx := db.Begin()
	user := utils.GetUser(c)

	// base logic: 查看video_id是否存在，查询是否有父评论，若无，直接存储，若有，parent_id = 父评论 ID 存储
	query := []string{"id = ?"}
	args := []interface{}{params.VideoID}
	_, err = (&models.Video{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("视频:[%d]未找到, 参数错误", params.VideoID))
	}
	query = append(query, "is_thumb = ?")
	args = []interface{}{params.ParentID, params.ISThumb}
	sqlComment, err := (&models.Comment{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	var commentData *models.Comment
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在
		commentData = &models.Comment{
			CreatorBase: models.CreatorBase{
				CreatorId: user.ID,
			},
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
			CreatorBase: models.CreatorBase{
				CreatorId: user.ID,
			},
			VideoId:  params.VideoID,
			Content:  params.Content,
			ParentId: sqlComment.ID,
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
		CreatorBase: models.CreatorBase{
			CreatorId: user.ID,
		},
		CommentId: commentData.ID,
	}).Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) CommentDelete(c *gin.Context, id uint) (err error) {
	// base logic: 查看关系表中有无此用户和评论对应关系，如果有，那么去评论表删除数据
	db := global.DB
	tx := db.Begin()
	user := utils.GetUser(c)

	_, err = (&models.UserComment{}).WhereOne(db, "creator_id = ? AND comment_id = ?", user.ID, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("用户:[%d]下评论:[%d]未找到，参数错误", user.ID, id))
	}
	err = (&models.Comment{}).Delete(tx, "id = ? AND creator_id = ?", id, user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = (&models.UserComment{}).Delete(tx, "creator_id = ? AND comment_id = ?", user.ID, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) CommentList(c *gin.Context, params *forms.CommentListForm) (response *forms.CommentListResponse, err error) {
	db := global.DB

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
			CreatedAt:   comment.CreatedAt.Format(constants.DateTime),
			UpdatedTime: comment.UpdatedAt.Format(constants.DateTime),
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

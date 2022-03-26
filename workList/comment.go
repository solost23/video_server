package workList

import (
	"errors"

	"gorm.io/gorm"

	"video_server/pkg/model"
)

func (w *WorkList) CreateComment(comment *model.Comment) error {
	// 查看video_id是否存在
	// 如果评论类型不在规定的两个里面，返回错误
	// 查询是否有此父评论id，若不存在，则parent_id = "0"，
	// 若存在，parent_id=查到的父评论存储评论, 存储
	videoID := w.ctx.Param("video_id")
	if comment.ISThumb != model.ISTHUMB && comment.ISThumb != model.ISCOMMENT {
		return errors.New("评论类型不存在")
	}
	var video = new(model.Video)
	if err := video.FindByVideoID(videoID, model.DELETENORMAL); err != nil {
		return err
	}
	var tmpComment = new(model.Comment)
	if err := tmpComment.FindByID(comment.ID); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if tmpComment.ID == "" {
		comment.ParentID = "0"
	}
	comment.VideoID = videoID
	if err := comment.Create(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteComment(comment *model.Comment) error {
	// 直接删除
	videoID := w.ctx.Param("video_id")
	commentID := w.ctx.Param("comment_id")
	if err := comment.DeleteByVideoIDAndCommentID(videoID, commentID); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetCommentByVideoID(comment *model.Comment) (comments []*model.Comment, err error) {
	// 直接通过video_id获取
	videoID := w.ctx.Param("video_id")
	comments, err = comment.FindByVideoID(videoID)
	if err != nil {
		return comments, err
	}
	return comments, nil
}

package create

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"video_server/pkg/model"
	"video_server/workList"
)

type Action struct {
	workList.WorkList
}

func NewActionWithCtx(ctx *gin.Context) *Action {
	r := &Action{}
	r.Init(ctx)
	return r
}

func (a *Action) Deal(request *Request) (resp Response, err error) {
	// 查看video_id是否存在
	// 如果评论类型不在规定的俩个里面，返回错误
	// 查询是否有此父评论，若不存在，则parent_id = "0"
	// 若存在，parent_id = 查到的父评论，存储
	if request.VideoID == "" {
		err = errors.New("request.VideoID not empty")
		return resp, err
	}
	if request.Content == "" {
		err = errors.New("request.Content not empty")
		return resp, err
	}
	if request.ISThumb != model.ISTHUMB && request.ISThumb != model.ISCOMMENT {
		err = errors.New("request.ISThumb error")
	}
	_, err = model.NewVideo(a.GetMysqlConn()).FindByVideoID(request.VideoID, model.DELETENORMAL)
	if err != nil {
		return resp, err
	}
	// 查找父评论
	parentComment, err := model.NewComment(a.GetMysqlConn()).FindByID(request.ParentID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, err
		}
	}
	if parentComment.ID == "" {
		request.ParentID = "0"
	}
	if err := model.NewComment(a.GetMysqlConn()).Create(a.buildRequest(request)); err != nil {
		return resp, err
	}
	return resp, err
}

func (a *Action) buildRequest(request *Request) (comment *model.Comment) {
	comment = &model.Comment{
		ID:         model.NewUUID(),
		VideoID:    request.VideoID,
		Content:    request.Content,
		ParentID:   request.ParentID,
		ISThumb:    request.ISThumb,
		CreateTime: model.GetCurrentTime(),
		UpdateTime: model.GetCurrentTime(),
	}
	return comment
}

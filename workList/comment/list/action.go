package list

import (
	"errors"
	"github.com/gin-gonic/gin"
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
	resp.List = []CommentInfo{}
	if request.VideoID == "" {
		err = errors.New("request.VideoID not empty")
		return resp, err
	}
	var comments []*model.Comment
	var total int64
	tx := model.NewComment(a.GetMysqlConn()).Connection().Where("video_id = ?", request.VideoID)
	if request.PageInfo == nil {
		request.PageInfo = &model.PageInfo{
			Page:     1,
			PageSize: 10,
		}
	}
	tx.Count(&total)
	err = tx.Offset(int((request.PageInfo.Page - 1) * request.PageInfo.PageSize)).Limit(int(request.PageInfo.PageSize)).Find(&comments).Error
	if err != nil {
		return resp, err
	}
	// 封装数据
	resp.PageInfo = model.PageInfo{
		Page:       request.PageInfo.Page,
		PageSize:   request.PageInfo.PageSize,
		TotalCount: int32(total),
	}
	for _, comment := range comments {
		commentInfo := CommentInfo{
			ID:         comment.ID,
			Content:    comment.Content,
			ParentID:   comment.ParentID,
			ISThumb:    comment.ISThumb,
			CreateTime: comment.CreateTime,
			UpdateTime: comment.UpdateTime,
		}
		resp.List = append(resp.List, commentInfo)
	}

	return resp, err
}

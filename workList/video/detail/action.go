package detail

import (
	"errors"
	"video_server/pkg/models"
	"video_server/workList"

	"github.com/gin-gonic/gin"
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
	if request.ID == "" {
		err = errors.New("request.ID not empty")
		return resp, err
	}
	data, err := models.NewVideo(a.GetMysqlConn()).FindByVideoID(request.ID, models.DELETENORMAL)
	if err != nil {
		return resp, err
	}
	// 封装数据，返回
	return a.buildResponse(data), err
}

func (a *Action) buildResponse(data *models.Video) (resp Response) {
	resp = Response{
		ID:           data.ID,
		UserID:       data.UserID,
		ClassID:      data.ClassID,
		Title:        data.Title,
		Introduce:    data.Introduce,
		ImageUrl:     data.ImageUrl,
		VideoUrl:     data.VideoUrl,
		ThumbCount:   data.ThumbCount,
		CommentCount: data.CommentCount,
		DeleteStatus: data.DeleteStatus,
		CreateTime:   data.CreateTime,
		UpdateTime:   data.UpdateTime,
	}
	return resp
}

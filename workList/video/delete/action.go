package delete

import (
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
	// 先查找视频信息,找不到则报错
	// 给视频的delete_status打上标记就可以
	// 注意:删视频的时候，视频下面的评论也要删除
	_, err = models.NewUser(a.GetMysqlConn()).FindByID(request.UserID)
	if err != nil {
		return resp, err
	}
	// 直接通过id, 分类, 视频id查询就可以
	if _, err = models.NewVideo(a.GetMysqlConn()).FindByUserIDANDClassIDAndID(request.UserID, request.ClassID, request.VideoID, models.DELETENORMAL); err != nil {
		return resp, err
	}
	// 删视频
	if _, err = models.NewVideo(a.GetMysqlConn()).Delete(request.VideoID); err != nil {
		return resp, err
	}
	// 删除视频下的评论
	comments, err := models.NewComment(a.GetMysqlConn()).FindByVideoID(request.VideoID)
	if err != nil {
		return resp, err
	}
	if len(comments) <= 0 {
		return resp, nil
	}
	if _, err = models.NewComment(a.GetMysqlConn()).DeleteByVideoID(request.VideoID); err != nil {
		return resp, err
	}
	return resp, err
}

package delete

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
	// 查看用户是否存在，若不存在，则返回错误
	// 若存在，则按照id删除
	// 注意:删除用户的时候，用户下的分类，分类下的视频，视频下的评论，都要删除
	if request.ID == "" {
		err = errors.New("request.ID not empty")
	}
	user, err := model.NewUser(a.GetMysqlConn()).FindByID(request.ID)
	if err != nil {
		return resp, err
	}
	// 删除用户下的分类
	categorys, err := model.NewCategory(a.GetMysqlConn()).FindByUserID(request.ID)
	if err != nil {
		return resp, err
	}
	if len(categorys) <= 0 {
		return resp, nil
	}
	var videos []*model.Video
	for _, category := range categorys {
		tmpVideos, err := model.NewVideo(a.GetMysqlConn()).FindByClassID(category.ID, model.DELETENORMAL)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = nil
			}
		}
		videos = append(videos, tmpVideos...)
	}
	// 删除分类
	_, err = model.NewCategory(a.GetMysqlConn()).DeleteByUserID(user.ID)
	if err != nil {
		return resp, err
	}
	// 通过视频查找评论
	for _, video := range videos {
		comments, err := model.NewComment(a.GetMysqlConn()).FindByVideoID(video.ID)
		if err != nil {
			return resp, err
		}
		if len(comments) <= 0 {
			return resp, err
		}
		// 删除评论
		_, err = model.NewComment(a.GetMysqlConn()).DeleteByVideoID(video.ID)
		if err != nil {
			return resp, err
		}
	}
	// 删除视频
	for _, video := range videos {
		_, err = model.NewVideo(a.GetMysqlConn()).Delete(video.ID)
		if err != nil {
			return resp, err
		}
	}
	return resp, err
}

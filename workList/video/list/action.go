package list

import (
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
	resp.List = []VideoInfo{}
	// 筛选
	videos, total, err := a.FindByFilter(request)
	if err != nil {
		return resp, err
	}
	// 组装数据，返回
	resp.PageInfo = model.PageInfo{
		Page:       request.PageInfo.Page,
		PageSize:   request.PageInfo.PageSize,
		TotalCount: int32(total),
	}
	for _, video := range videos {
		videoInfo := VideoInfo{
			ID:           video.ID,
			UserID:       video.UserID,
			ClassID:      video.ClassID,
			Title:        video.Title,
			Introduce:    video.Introduce,
			ImageUrl:     video.ImageUrl,
			VideoUrl:     video.VideoUrl,
			ThumbCount:   video.ThumbCount,
			CommentCount: video.CommentCount,
			DeleteStatus: video.DeleteStatus,
			CreateTime:   video.CreateTime,
			UpdateTime:   video.UpdateTime,
		}
		resp.List = append(resp.List, videoInfo)
	}
	return resp, err
}

func (a *Action) FindByFilter(request *Request) (videos []model.Video, total int64, err error) {
	// 通过用户名查找出用户id
	// 通过分类名查找出分类id
	// 通过视频标题找出视频内容
	tx := model.NewVideo(a.GetMysqlConn()).Connection().Select("*")
	if request.Filter.UserName != "" {
		user, err := model.NewUser(a.GetMysqlConn()).FindBYUserName(request.Filter.UserName)
		if err != nil {
			return videos, total, err
		}
		tx.Where("user_id = ?", user.ID)
	}
	if request.Filter.CategoryName != "" {
		category, err := model.NewCategory(a.GetMysqlConn()).FindByClassName(request.Filter.CategoryName)
		if err != nil {
			return videos, total, err
		}
		tx.Where("class_id = ?", category.ID)
	}
	if request.Filter.VideoTitle != "" {
		tx.Where("title = ?", request.Filter.VideoTitle)
	}
	tx.Count(&total)
	// 数据分页
	if request.PageInfo == nil {
		request.PageInfo = &model.PageInfo{
			Page:     1,
			PageSize: 10,
		}
	}
	err = tx.Offset(int((request.PageInfo.Page - 1) * request.PageInfo.PageSize)).Limit(int(request.PageInfo.PageSize)).Find(&videos).Error
	return videos, total, err
}

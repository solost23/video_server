package list

import (
	"github.com/gin-gonic/gin"

	"video_server/pkg/models"
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
	resp.PageInfo = models.PageInfo{
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

func (a *Action) FindByFilter(request *Request) (videos []models.Video, total int64, err error) {
	// 通过用户名查找出用户id
	// 通过分类名查找出分类id
	// 通过视频标题找出视频内容
	tx := models.NewVideo(a.GetMysqlConn()).Connection().Select("*")
	if request.Filter.UserName != "" {
		var users []*models.User
		err = models.NewUser(a.GetMysqlConn()).Connection().Where("user_name LIKE ?", request.Filter.UserName).First(&users).Error
		if err != nil {
			return videos, total, err
		}
		// 保存用户id列表
		var userIds []string
		for _, user := range users {
			userIds = append(userIds, user.ID)
		}
		tx.Where("user_id in ?", userIds)
	}
	if request.Filter.CategoryName != "" {
		var categorys []*models.Category
		err = models.NewCategory(a.GetMysqlConn()).Connection().Where("title = ?", request.Filter.CategoryName).First(&categorys).Error
		if err != nil {
			return videos, total, err
		}
		// 保存视频分类id列表
		var categoryIds []string
		for _, category := range categorys {
			categoryIds = append(categoryIds, category.ID)
		}
		tx.Where("class_id in ?", categoryIds)
	}
	if request.Filter.VideoTitle != "" {
		tx.Where("title LIKE ?", models.LikeFilter(request.Filter.VideoTitle))
	}
	tx.Where("delete_status = ?", models.DELETENORMAL)
	tx.Count(&total)
	// 数据分页
	if request.PageInfo == nil {
		request.PageInfo = &models.PageInfo{
			Page:     1,
			PageSize: 10,
		}
	}
	err = tx.Offset(int((request.PageInfo.Page - 1) * request.PageInfo.PageSize)).Limit(int(request.PageInfo.PageSize)).Find(&videos).Error
	return videos, total, err
}

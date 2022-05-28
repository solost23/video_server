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
	resp.List = []UserInfo{}
	// 筛选
	users, total, err := a.FindByFilter(request)
	if err != nil {
		return resp, err
	}
	// 组装数据，返回
	resp.PageInfo = model.PageInfo{
		Page:       request.PageInfo.Page,
		PageSize:   request.PageInfo.PageSize,
		TotalCount: int32(total),
	}
	for _, user := range users {
		userInfo := UserInfo{
			ID:           user.ID,
			UserName:     user.UserName,
			Nickname:     user.Nickname,
			Role:         user.Role,
			Avatar:       user.Avatar,
			Introduce:    user.Introduce,
			FansCount:    user.FansCount,
			CommentCount: user.CommentCount,
			CreateTime:   user.CreateTime,
			UpdateTime:   user.UpdateTime,
		}
		resp.List = append(resp.List, userInfo)
	}

	return resp, err
}

func (a *Action) FindByFilter(request *Request) (users []model.User, total int64, err error) {
	//ID       string `json:"id"`
	//UserName string `json:"userName"`
	//Role     string `json:"role"`
	tx := model.NewUser(a.GetMysqlConn()).Connection().Select("*")
	if request.Filter.ID != "" {
		tx.Where("id LIKE ?", model.LikeFilter(request.Filter.ID))
	}
	if request.Filter.UserName != "" {
		tx.Where("user_name LIKE ?", model.LikeFilter(request.Filter.UserName))
	}
	if request.Filter.Role != "" {
		tx.Where("user_name LIKE ?", model.LikeFilter(request.Filter.Role))
	}
	tx.Count(&total)
	// 数据分页
	if request.PageInfo == nil {
		request.PageInfo = &model.PageInfo{
			Page:     1,
			PageSize: 10,
		}
	}
	err = tx.Offset(int((request.PageInfo.Page - 1) * request.PageInfo.PageSize)).Limit(int(request.PageInfo.PageSize)).Find(&users).Error
	return users, total, err
}

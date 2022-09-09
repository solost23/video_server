package forms

import (
	"video_server/pkg/models"
	"video_server/pkg/utils"
)

type RegisterForm struct {
	UserName  string `json:"userName" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
	Nickname  string `json:"nickname"`
	Role      string `json:"role" binding:"required,oneof=ADMIN USER"`
	Avatar    string `json:"avatar"`
	Introduce string `json:"introduce"`
}

type RegisterResponse struct {
}

type LoginForm struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Device   string `json:"device" comment:"设备类型" binding:"required,oneof=ios android web"`
}

type LogoutForm struct {
	Device string `json:"device" comment:"设备类型" binding:"required,oneof=ios android web"`
}

type LoginResponse struct {
	models.User
	IsFirstLogin uint   `json:"isFirstLogin"`
	Token        string `json:"token"`
}

type ListForm struct {
	utils.PageForm
	ID       uint   `form:"id"`
	UserName string `form:"userName"`
	Role     string `form:"role"`
}

type ListResponse struct {
	PageList *utils.PageList
	List     []ListRecord `json:"list"`
}

type ListRecord struct {
	ID           uint   `json:"id"`
	UserName     string `json:"user_name"`
	Nickname     string `json:"nickname"`
	Role         string `json:"role"`
	Avatar       string `json:"avatar"`
	Introduce    string `json:"introduce"`
	FansCount    int64  `json:"fansCount"`
	CommentCount int64  `json:"commentCount"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

type UserUpdateForm struct {
	UserName  string `json:"userName" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
	Nickname  string `json:"nickname"`
	Role      string `json:"role" binding:"required,oneof=ADMIN,USER"`
	Avatar    string `json:"avatar"`
	Introduce string `json:"introduce"`
}

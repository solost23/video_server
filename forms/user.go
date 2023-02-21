package forms

import (
	"video_server/pkg/models"
	"video_server/pkg/utils"
)

type RegisterForm struct {
	Username  *string `json:"username" binding:"required"`
	Password  *string `json:"password" binding:"required,min=8"`
	Nickname  *string `json:"nickname"`
	Role      *uint   `json:"role" binding:"required,oneof=0 1" comment:"0 管理员 1 普通用户"`
	Avatar    *string `json:"avatar"`
	Introduce *string `json:"introduce"`
}

type LoginForm struct {
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password" binding:"required,min=8"`
	Device   *string `json:"device" binding:"required,oneof=ios android web" comment:"设备类型"`
}

type LogoutForm struct {
	Device *string `json:"device" binding:"required,oneof=ios android web" comment:"设备类型"`
}

type LoginResponse struct {
	*models.User
	IsFirstLogin *uint   `json:"isFirstLogin"`
	Token        *string `json:"token"`
}

type ListForm struct {
	*utils.PageForm
	ID       *uint   `form:"id"`
	Username *string `form:"username"`
	Role     *uint   `form:"role"`
}

type ListResponse struct {
	PageList *utils.PageList
	List     []*ListRecord `json:"list"`
}

type ListRecord struct {
	ID           *uint   `json:"id"`
	Username     *string `json:"username"`
	Nickname     *string `json:"nickname"`
	Role         *uint   `json:"role"`
	Avatar       *string `json:"avatar"`
	Introduce    *string `json:"introduce"`
	FansCount    *int64  `json:"fansCount"`
	CommentCount *int64  `json:"commentCount"`
	CreatedAt    *string `json:"createdAt"`
	UpdatedAt    *string `json:"updatedAt"`
}

type UserUpdateForm struct {
	Username  *string `json:"username" binding:"required"`
	Password  *string `json:"password" binding:"required,min=8"`
	Nickname  *string `json:"nickname"`
	Role      *uint   `json:"role" binding:"required,oneof=0 1" comment:"0 管理员 1 普通用户"`
	Avatar    *string `json:"avatar"`
	Introduce *string `json:"introduce"`
}

type SearchForm struct {
	*utils.PageForm
	Keyword *string `form:"keyword"`
}

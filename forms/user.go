package forms

import "video_server/pkg/models"

type RegisterForm struct {
	UserName  string `json:"userName" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
	Nickname  string `json:"nickname"`
	Role      string `json:"role" binding:"required,oneof=ADMIN,USER"`
	Avatar    string `json:"avatar"`
	Introduce string `json:"introduce"`
}

type RegisterResponse struct {
}

type LoginForm struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	TokenStr string `json:"tokenStr"`
}

type ListForm struct {
	*models.PageForm
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	Role     string `json:"role"`
}

type ListResponse struct {
	PageList *models.PageList
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
package forms

import "video_server/pkg/utils"

type VideoInsertForm struct {
	CategoryId uint   `json:"categoryId" binding:"required"`
	Title      string `json:"title" binding:"required,max=20"`
	Introduce  string `json:"introduce" binding:"required,max=20"`
	ImageUrl   string `json:"imageUrl" binding:"required" comment:"oss上传之后生成的url,前端再传过来"`
	VideoUrl   string `json:"videoUrl" binding:"required" comment:"oss上传之后生成的url,前端再传过来"`
}

type VideoListForm struct {
	utils.PageForm
	CategoryName string `json:"categoryName"`
	UserName     string `json:"userName"`
	VideoTitle   string `json:"videoTitle"`
	Introduce    string `json:"introduce"`
}

type VideoListRecord struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"userId"`
	UserName     string `json:"userName"`
	CategoryID   uint   `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Title        string `json:"title"`
	Introduce    string `json:"introduce"`
	ImageUrl     string `json:"imageUrl"`
	VideoUrl     string `json:"videoUrl"`
	ThumbCount   int64  `json:"thumbCount"`
	CommentCount int64  `json:"commentCount"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type VideoListResponse struct {
	Records  []VideoListRecord `json:"records"`
	PageList *utils.PageList
}

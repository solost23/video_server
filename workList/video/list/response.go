package list

import "video_server/pkg/model"

type Response struct {
	List     []VideoInfo    `json:"list"`
	PageInfo model.PageInfo `json:"pageInfo"`
}

type VideoInfo struct {
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	ClassID      string `json:"classId"`
	Title        string `json:"title"`
	Introduce    string `json:"introduce"`
	ImageUrl     string `json:"imageUrl"`
	VideoUrl     string `json:"videoUrl"`
	ThumbCount   int64  `json:"thumbCount"`
	CommentCount int64  `json:"commentCount"`
	DeleteStatus string `json:"deleteStatus"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
}

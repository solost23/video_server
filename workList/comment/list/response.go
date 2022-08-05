package list

import "video_server/pkg/models"

type Response struct {
	List     []CommentInfo `json:"list"`
	PageInfo models.PageInfo
}

type CommentInfo struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	ParentID   string `json:"parentID"`
	ISThumb    string `json:"ISThumb"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

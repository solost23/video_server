package list

import "video_server/pkg/model"

type Response struct {
	List     []CommentInfo `json:"list"`
	PageInfo model.PageInfo
}

type CommentInfo struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	ParentID   string `json:"parentID"`
	ISThumb    string `json:"ISThumb"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

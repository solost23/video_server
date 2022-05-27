package list

import "video_server/pkg/model"

type Response struct {
	List     []CategoryInfo `json:"list"`
	PageInfo model.PageInfo `json:"pageInfo"`
}

type CategoryInfo struct {
	ID         string `json:"id"`
	UserID     string `json:"userID"`
	Title      string `json:"title"`
	Introduce  string `json:"introduce"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

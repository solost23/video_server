package list

import "video_server/pkg/model"

type Request struct {
	PageInfo *model.PageInfo `json:"pageInfo"`
	Filter   *Filter         `json:"filter"`
}

type Filter struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Role     string `json:"role"`
}

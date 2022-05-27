package list

import "video_server/pkg/model"

type Request struct {
	PageInfo *model.PageInfo `json:"pageInfo"`
	Filter   *Filter         `json:"filter"`
}

type Filter struct {
	UserID string `json:"userId"`
}

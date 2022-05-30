package list

import "video_server/pkg/model"

type Request struct {
	PageInfo *model.PageInfo `json:"pageInfo"`
	Filter   *Filter
}

type Filter struct {
	VideoID string `json:"videoId"`
}

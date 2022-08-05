package list

import "video_server/pkg/models"

type Request struct {
	PageInfo *models.PageInfo `json:"pageInfo"`
	Filter   *Filter
}

type Filter struct {
	VideoID string `json:"videoId"`
}

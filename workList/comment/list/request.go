package list

import "video_server/pkg/model"

type Request struct {
	PageInfo *model.PageInfo `json:"pageInfo"`
	VideoID  string          `json:"videoId"`
}

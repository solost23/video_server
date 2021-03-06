package list

import "video_server/pkg/model"

// 分类名，用户名，视频标题 搜索，并支持分页操作

type Request struct {
	PageInfo *model.PageInfo `json:"pageInfo"`
	Filter   *Filter         `json:"filter"`
}

type Filter struct {
	CategoryName string `json:"categoryName"`
	UserName     string `json:"userName"`
	VideoTitle   string `json:"videoTitle"`
}

package list

import "video_server/pkg/models"

type Request struct {
	PageInfo *models.PageInfo `json:"pageInfo"`
	Filter   *Filter          `json:"filter"`
}

type Filter struct {
	RoleName string `json:"RoleName"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

package forms

import (
	"video_server/pkg/utils"
)

type RoleInsertForm struct {
	RoleName string `json:"roleName" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required"`
}

type RoleListForm struct {
	utils.PageForm
	RoleName string `json:"roleName"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

type RoleListResponse struct {
	PageList *utils.PageList
	Records  []RoleInsertForm `json:"records"`
}

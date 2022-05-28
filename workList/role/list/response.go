package list

import "video_server/pkg/model"

type Response struct {
	PageInfo model.PageInfo `json:"pageInfo"`
	List     []*RoleInfo    `json:"list"`
}

type RoleInfo struct {
	RoleName string `json:"roleName"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

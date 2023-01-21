package routers

import (
	"video_server/forms"
	"video_server/pkg/response"
	"video_server/pkg/utils"
	"video_server/services"

	"github.com/gin-gonic/gin"
)

func roleInsert(c *gin.Context) {
	params := &forms.RoleInsertForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	err := (&services.Service{}).InsertRole(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.MessageSuccess(c, "成功", nil)
}

func roleDelete(c *gin.Context) {
	params := &forms.RoleInsertForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	err := (&services.Service{}).DeleteRole(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.MessageSuccess(c, "成功", nil)
}

func roleList(c *gin.Context) {
	params := &forms.RoleListForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	result, err := (&services.Service{}).ListRole(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.Success(c, result)
}

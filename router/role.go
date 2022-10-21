package router

import (
	"video_server/forms"
	"video_server/pkg/response"
	"video_server/pkg/utils"
	"video_server/services"

	"github.com/gin-gonic/gin"
)

// @Summary add roleAuth
// @Description add roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Param data body models.CasbinModel true "角色"
// @Accept json
// @Produce json
// @Success 200
// @Router /role/create [post]
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

// @Summary delete role
// @Description delete roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Param data body models.CasbinModel true "角色"
// @Accept json
// @Produce json
// @Success 200
// @Router /role/delete [post]
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

// @Summary get all roleAuth
// @Description get all roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /role/list [post]
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

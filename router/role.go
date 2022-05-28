package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	create_role "video_server/workList/role/create"
	delete_role "video_server/workList/role/delete"
	list_role "video_server/workList/role/list"
)

// @Summary add roleAuth
// @Description add roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Param data body model.CasbinModel true "角色"
// @Accept json
// @Produce json
// @Success 200
// @Router /role/create [post]
func addRoleAuth(c *gin.Context) {
	request := &create_role.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := create_role.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary delete role
// @Description delete roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Param data body model.CasbinModel true "角色"
// @Accept json
// @Produce json
// @Success 200
// @Router /role/delete [post]
func deleteRoleAuth(c *gin.Context) {
	request := &delete_role.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := delete_role.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary get all roleAuth
// @Description get all roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /role/list [post]
func getAllRoleAuth(c *gin.Context) {
	request := &list_role.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := list_role.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

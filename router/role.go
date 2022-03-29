package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"video_server/pkg/model"
	"video_server/workList"
)

// @Summary add roleAuth
// @Description add roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Param data body model.CasbinModel true "角色"
// @Accept json
// @Produce json
// @Success 200
// @Router /role [post]
func addRoleAuth(c *gin.Context) {
	var casbinModel = new(model.CasbinModel)
	if err := c.Bind(casbinModel); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := workList.NewWorkList(c).AddRoleAuth(casbinModel); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, casbinModel)
	return
}

// @Summary delete role
// @Description delete roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Param data body model.CasbinModel true "角色"
// @Accept json
// @Produce json
// @Success 200
// @Router /role [delete]
func deleteRoleAuth(c *gin.Context) {
	var casbinModel = new(model.CasbinModel)
	if err := c.Bind(casbinModel); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := workList.NewWorkList(c).DeleteRoleAuth(casbinModel); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "delete roleAuth success")
	return
}

// @Summary get all roleAuth
// @Description get all roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /role [get]
func getAllRoleAuth(c *gin.Context) {
	var casbinModel = new(model.CasbinModel)
	casbinModelList, err := workList.NewWorkList(c).GetAllRoleAuth(casbinModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, casbinModelList)
	return
}

// @Summary get roleAuth
// @Description get roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /role/{role_name} [get]
func getRoleAuth(c *gin.Context) {
	roleName := c.Param("role_name")
	c.Set("role_name", roleName)
	var casbinModel = new(model.CasbinModel)
	casbinModelList, err := workList.NewWorkList(c).GetRoleAuth(casbinModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, casbinModelList)
	return
}

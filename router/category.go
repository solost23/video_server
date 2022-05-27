package router

import (
	"github.com/gin-gonic/gin"
	"net/http"

	create_category "video_server/workList/category/create"
	list_category "video_server/workList/category/list"
	update_category "video_server/workList/category/update"
)

// @Summary create_class
// @Description add category
// @Tags Class
// @Security ApiKeyAuth
// @Param data body model.Class true "类别"
// @Accept json
// @Produce json
// @Success 200
// @Router /category/create [post]
func createCategory(c *gin.Context) {
	request := &create_category.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := create_category.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary update_class
// @Description update category
// @Tags Class
// @Security ApiKeyAuth
// @Param data body model.Class true "类别"
// @Accept json
// @Produce json
// @Success 200
// @Router /category/update [post]
func updateCategory(c *gin.Context) {
	request := &update_category.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := update_category.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary get user all category
// @Description get user all category
// @Tags Class
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /category/list [post]
func listCategory(c *gin.Context) {
	request := &list_category.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := list_category.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

package router

import (
	create_category "video_server/workList/category/create"
	list_category "video_server/workList/category/list"
	update_category "video_server/workList/category/update"

	"github.com/gin-gonic/gin"
)

// @Summary create_class
// @Description add category
// @Tags Class
// @Security ApiKeyAuth
// @Param data body models.Class true "类别"
// @Accept json
// @Produce json
// @Success 200
// @Router /category/create [post]
func categoryInsert(c *gin.Context) {
	request := &create_category.Request{}
	if err := c.ShouldBind(&request); err != nil {
		Render(c, err)
		return
	}
	data, err := create_category.NewActionWithCtx(c).Deal(request)
	if err != nil {
		Render(c, err)
		return
	}
	Render(c, err, data)
}

// @Summary update_class
// @Description update category
// @Tags Class
// @Security ApiKeyAuth
// @Param data body models.Class true "类别"
// @Accept json
// @Produce json
// @Success 200
// @Router /category/update [post]
func categoryUpdate(c *gin.Context) {
	request := &update_category.Request{}
	if err := c.ShouldBind(&request); err != nil {
		Render(c, err)
		return
	}
	data, err := update_category.NewActionWithCtx(c).Deal(request)
	if err != nil {
		Render(c, err)
		return
	}
	Render(c, err, data)
}

// @Summary get user all category
// @Description get user all category
// @Tags Class
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /category/list [post]
func categoryList(c *gin.Context) {
	request := &list_category.Request{}
	if err := c.ShouldBind(&request); err != nil {
		Render(c, err)
		return
	}
	data, err := list_category.NewActionWithCtx(c).Deal(request)
	if err != nil {
		Render(c, err)
		return
	}
	Render(c, err, data)
}

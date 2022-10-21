package router

import (
	"video_server/forms"
	"video_server/pkg/response"
	"video_server/pkg/utils"
	"video_server/services"

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
	params := &forms.CategoryInsertForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}

	err := (&services.Service{}).InsertCategory(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.MessageSuccess(c, "成功", nil)
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
	UIdForm := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}
	params := &forms.CategoryUpdateForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}

	err := (&services.Service{}).UpdateCategory(c, UIdForm.Id, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.MessageSuccess(c, "成功", nil)
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
	params := &forms.CategoryListForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Size == 0 {
		params.Size = 10
	}
	result, err := (&services.Service{}).ListCategory(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.Success(c, result)
}

func searchCategory(c *gin.Context) {
	params := &forms.SearchForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Size <= 0 {
		params.Size = 10
	}
	result, err := (&services.Service{}).SearchCategory(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.Success(c, result)
}

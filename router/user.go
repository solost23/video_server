package router

import (
	"video_server/forms"
	"video_server/pkg/response"
	"video_server/pkg/utils"
	"video_server/workList"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	params := &forms.RegisterForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	err := (&workList.UserService{}).Register(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.MessageSuccess(c, "成功", nil)
}

func uploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	result, err := (&workList.UserService{}).UploadAvatar(c, file)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.Success(c, result)
}

func login(c *gin.Context) {
	params := &forms.LoginForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	result, err := (&workList.UserService{}).Login(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.Success(c, result)
}

// @Summary get_user_info
// @Description get user info
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user/detail [post]
func userDetail(c *gin.Context) {
	UIdForm := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}
	result, err := (&workList.UserService{}).Detail(c, UIdForm.Id)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.Success(c, result)
}

// @Summary delete_user_info
// @Description delete user info
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user/delete [post]
func userDelete(c *gin.Context) {
	UIdForm := utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}

	err := (&workList.UserService{}).Delete(c, UIdForm.Id)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.MessageSuccess(c, "成功", nil)
}

// @Summary update_user_info
// @Description update user info
// @Tags User
// @Security ApiKeyAuth
// @Param data body models.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/update [post]
func userUpdate(c *gin.Context) {
	UIdForm := utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}
	params := &forms.UserUpdateForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}

	err := (&workList.UserService{}).Update(c, UIdForm.Id, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.MessageSuccess(c, "成功", nil)
}

// @Summary get_all_user_info
// @Description get all user info
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user/list [post]
func userList(c *gin.Context) {
	params := &forms.ListForm{}
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
	result, err := (&workList.UserService{}).List(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.Success(c, result)
}

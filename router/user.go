package router

import (
	"video_server/forms"
	"video_server/pkg/response"
	"video_server/pkg/utils"
	"video_server/services"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	params := &forms.RegisterForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	err := (&services.Service{}).Register(c, params)
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
	result, err := (&services.Service{}).UploadAvatar(c, file)
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
	result, err := (&services.Service{}).Login(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.Success(c, result)
}

func logout(c *gin.Context) {
	params := &forms.LogoutForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	err := (&services.Service{}).Logout(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.MessageSuccess(c, "成功", nil)
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
	result, err := (&services.Service{}).Detail(c, UIdForm.Id)
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

	err := (&services.Service{}).DeleteUser(c, UIdForm.Id)
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
	UIdForm := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}
	params := &forms.UserUpdateForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}

	err := (&services.Service{}).UpdateUser(c, UIdForm.Id, params)
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
	result, err := (&services.Service{}).ListUser(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.Success(c, result)
}

func searchUser(c *gin.Context) {
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
	result, err := (&services.Service{}).SearchUser(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.Success(c, result)
}

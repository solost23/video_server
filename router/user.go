package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	list_video "video_server/workList/user/List"
	delete_video "video_server/workList/user/delete"
	detail_video "video_server/workList/user/detail"
	login_video "video_server/workList/user/login"
	register_video "video_server/workList/user/register"
	update_video "video_server/workList/user/update"
)

// @Summary register
// @Description add user
// @Tags User
// @Param data body model.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /register [post]
func register(c *gin.Context) {
	request := &register_video.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := register_video.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary login
// @Description user login
// @Tags User
// @Param data body model.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /login [post]
func login(c *gin.Context) {
	request := &login_video.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := login_video.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary get_user_info
// @Description get user info
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user/detail [post]
func getUserInfo(c *gin.Context) {
	request := &detail_video.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := detail_video.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary delete_user_info
// @Description delete user info
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user/delete [post]
func deleteUserInfo(c *gin.Context) {
	request := &delete_video.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := delete_video.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary update_user_info
// @Description update user info
// @Tags User
// @Security ApiKeyAuth
// @Param data body model.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/update [post]
func updateUserInfo(c *gin.Context) {
	request := &update_video.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := update_video.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary get_all_user_info
// @Description get all user info
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user/list [post]
func getAllUserInfo(c *gin.Context) {
	request := &list_video.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := list_video.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

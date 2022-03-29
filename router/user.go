package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"video_server/pkg/model"
	"video_server/workList"
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
	var user = new(model.User)
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := workList.NewWorkList(c).Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
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
	var user = new(model.User)
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenStr, err := workList.NewWorkList(c).Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Auth": tokenStr,
	})
	return
}

// @Summary get_user_info
// @Description get user info
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user/{user_name} [get]
func getUserInfo(c *gin.Context) {
	var user = new(model.User)
	if err := workList.NewWorkList(c).GetUserInfo(user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

// @Summary delete_user_info
// @Description delete user info
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user/{user_name} [delete]
func deleteUserInfo(c *gin.Context) {
	var user = new(model.User)
	if err := workList.NewWorkList(c).DeleteUserInfo(user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

// @Summary update_user_info
// @Description update user info
// @Tags User
// @Security ApiKeyAuth
// @Param data body model.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/{user_name} [delete]
func updateUserInfo(c *gin.Context) {
	// 通过用户名去更新字段
	var user = new(model.User)
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	if err := workList.NewWorkList(c).UpdateUserInfo(user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

// @Summary get_all_user_info
// @Description get all user info
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user [get]
func getAllUserInfo(c *gin.Context) {
	var user = new(model.User)
	users, err := workList.NewWorkList(c).GETAllUserInfo(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
	return
}

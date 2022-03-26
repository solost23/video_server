package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"video_server/pkg/model"
	"video_server/workList"
)

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

func getUserInfo(c *gin.Context) {
	var user = new(model.User)
	if err := workList.NewWorkList(c).GetUserInfo(user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func deleteUserInfo(c *gin.Context) {
	var user = new(model.User)
	if err := workList.NewWorkList(c).DeleteUserInfo(user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

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

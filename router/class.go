package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"video_server/pkg/model"
	"video_server/workList"
)

func createClass(c *gin.Context) {
	var class = new(model.Class)
	if err := c.ShouldBind(class); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := workList.NewWorkList(c).CreateClass(class); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, class)
	return
}

func updateClass(c *gin.Context) {
	var class = new(model.Class)
	if err := c.ShouldBind(class); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := workList.NewWorkList(c).UpdateClass(class); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, class)
	return
}

func getUserAllClass(c *gin.Context) {
	var class = new(model.Class)
	userClasses, err := workList.NewWorkList(c).GetUserAllClass(class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userClasses)
	return
}

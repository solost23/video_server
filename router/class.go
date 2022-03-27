package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"video_server/pkg/model"
	"video_server/workList"
)

// PingExample godoc
// @Summary ping class
// @Schemes
// @Description add class
// @Tags Class
// @Accept json
// @Produce json
// @Success 200
// @Router /class/{user_name} [post]
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

// PingExample godoc
// @Summary ping class
// @Schemes
// @Description update class
// @Tags Class
// @Accept json
// @Produce json
// @Success 200
// @Router /class/{user_name}/{class_id} [put]
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

// PingExample godoc
// @Summary ping class
// @Schemes
// @Description get user all class
// @Tags Class
// @Accept json
// @Produce json
// @Success 200
// @Router /class/{user_name} [get]
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

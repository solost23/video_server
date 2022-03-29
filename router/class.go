package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"video_server/pkg/model"
	"video_server/workList"
)

// @Summary create_class
// @Description add class
// @Tags Class
// @Security ApiKeyAuth
// @Param data body model.Class true "类别"
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

// @Summary update_class
// @Description update class
// @Tags Class
// @Security ApiKeyAuth
// @Param data body model.Class true "类别"
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

// @Summary get user all class
// @Description get user all class
// @Tags Class
// @Security ApiKeyAuth
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

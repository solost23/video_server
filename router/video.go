package router

import (
	"net/http"
	"video_server/scheduler/video/create"
	delete2 "video_server/scheduler/video/delete"
	"video_server/scheduler/video/detail"
	list2 "video_server/scheduler/video/list"

	"github.com/gin-gonic/gin"
)

// @Summary add video
// @Description add video
// @Tags Video
// @Security ApiKeyAuth
// @Param data body model.Video true "视频"
// @Accept json
// @Produce json
// @Success 200
// @Router /video/create [post]
func createVideo(c *gin.Context) {
	request := &create.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := create.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary delete video
// @Description delete video
// @Tags Video
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /video/delete [post]
func deleteVideo(c *gin.Context) {
	request := &delete2.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := delete2.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary video detail
// @Description video detail
// @Tags Video
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /video/detail [get]
func videoDetail(c *gin.Context) {
	request := &detail.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := detail.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary video list
// @Description video list
// @Tags Video
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /video/list [post]
func list(c *gin.Context) {
	request := &list2.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := list2.NewActionWithCtx(c).Deal(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	create_video "video_server/workList/video/create"
	delete_video "video_server/workList/video/delete"
	detail_video "video_server/workList/video/detail"
	list_video "video_server/workList/video/list"
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
	request := &create_video.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := create_video.NewActionWithCtx(c).Deal(request)
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

// @Summary video detail
// @Description video detail
// @Tags Video
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /video/detail [get]
func videoDetail(c *gin.Context) {
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

// @Summary video list
// @Description video list
// @Tags Video
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /video/list [post]
func list(c *gin.Context) {
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

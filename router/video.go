package router

import (
	create2 "video_server/workList/video/create"
	delete3 "video_server/workList/video/delete"
	detail2 "video_server/workList/video/detail"
	list3 "video_server/workList/video/list"

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
	request := &create2.Request{}
	if err := c.ShouldBind(&request); err != nil {
		Render(c, err)
		return
	}
	data, err := create2.NewActionWithCtx(c).Deal(request)
	if err != nil {
		Render(c, err)
		return
	}
	Render(c, err, data)
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
	request := &delete3.Request{}
	if err := c.ShouldBind(&request); err != nil {
		Render(c, err)
		return
	}
	data, err := delete3.NewActionWithCtx(c).Deal(request)
	if err != nil {
		Render(c, err)
		return
	}
	Render(c, err, data)
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
	request := &detail2.Request{}
	if err := c.ShouldBind(&request); err != nil {
		Render(c, err)
		return
	}
	data, err := detail2.NewActionWithCtx(c).Deal(request)
	if err != nil {
		Render(c, err)
		return
	}
	Render(c, err, data)
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
	request := &list3.Request{}
	if err := c.ShouldBind(&request); err != nil {
		Render(c, err)
		return
	}
	data, err := list3.NewActionWithCtx(c).Deal(request)
	if err != nil {
		Render(c, err)
		return
	}
	Render(c, err, data)
}

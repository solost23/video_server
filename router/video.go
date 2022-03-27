package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"video_server/pkg/model"
	"video_server/workList"
)

// PingExample godoc
// @Summary ping video
// @Schemes
// @Description add video
// @Tags Video
// @Accept json
// @Produce json
// @Success 200
// @Router /video/{user_name}/{class_id} [post]
func createVideo(c *gin.Context) {
	var video = new(model.Video)
	if err := c.ShouldBind(video); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := workList.NewWorkList(c).CreateVideo(video); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, video)
	return
}

// PingExample godoc
// @Summary ping video
// @Schemes
// @Description delete video
// @Tags Video
// @Accept json
// @Produce json
// @Success 200
// @Router /video/{user_name}/{class_id}/{video_id} [delete]
func deleteVideo(c *gin.Context) {
	var video = new(model.Video)
	if err := workList.NewWorkList(c).DeleteVideo(video); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, video)
	return
}

// PingExample godoc
// @Summary ping video
// @Schemes
// @Description get video
// @Tags Video
// @Accept json
// @Produce json
// @Success 200
// @Router /video/{user_name}/{class_id}/{video_id} [get]
func getVideo(c *gin.Context) {
	var video = new(model.Video)
	if err := workList.NewWorkList(c).GetVideo(video); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, video)
	return
}

// PingExample godoc
// @Summary ping video
// @Schemes
// @Description get video by user_name and class_id
// @Tags Video
// @Accept json
// @Produce json
// @Success 200
// @Router /video/{user_name}/{class_id} [get]
func getVideoByUserNameAndClassID(c *gin.Context) {
	var video = new(model.Video)
	videos, err := workList.NewWorkList(c).GetVideoByUserNameAndClassID(video)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, videos)
	return
}

// PingExample godoc
// @Summary ping video
// @Schemes
// @Description get video by user_name
// @Tags Video
// @Accept json
// @Produce json
// @Success 200
// @Router /video/{user_name} [get]
func getVideoByUserName(c *gin.Context) {
	var video = new(model.Video)
	videos, err := workList.NewWorkList(c).GetVideoByUserName(video)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, videos)
	return
}

// PingExample godoc
// @Summary ping video
// @Schemes
// @Description get all video
// @Tags Video
// @Accept json
// @Produce json
// @Success 200
// @Router /video [get]
func getALLVideo(c *gin.Context) {
	var video = new(model.Video)
	videos, err := workList.NewWorkList(c).GetAllVideo(video)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, videos)
	return
}

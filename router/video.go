package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"video_server/pkg/model"
	"video_server/workList"
)

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

func deleteVideo(c *gin.Context) {
	var video = new(model.Video)
	if err := workList.NewWorkList(c).DeleteVideo(video); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, video)
	return
}

func getVideo(c *gin.Context) {
	var video = new(model.Video)
	if err := workList.NewWorkList(c).GetVideo(video); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, video)
	return
}

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

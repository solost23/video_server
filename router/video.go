package router

import (
	"video_server/forms"
	"video_server/pkg/response"
	"video_server/pkg/utils"
	"video_server/workList"

	"github.com/gin-gonic/gin"
)

// @Summary add video
// @Description add video
// @Tags Video
// @Security ApiKeyAuth
// @Param data body models.Video true "视频"
// @Accept json
// @Produce json
// @Success 200
// @Router /video/create [post]

func videoUploadImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	result, err := (&workList.VideoService{}).VideoUploadImg(c, file)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.Success(c, result)
}

func videoUploadVid(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	result, err := (&workList.VideoService{}).VideoUploadVid(c, file)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.Success(c, result)
}

func videoInsert(c *gin.Context) {
	params := &forms.VideoInsertForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if err := (&workList.VideoService{}).VideoInsert(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.MessageSuccess(c, "成功", nil)
}

// @Summary delete video
// @Description delete video
// @Tags Video
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /video/delete [post]
func videoDelete(c *gin.Context) {
	UIdForm := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if err := (&workList.VideoService{}).VideoDelete(c, UIdForm.Id); err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.MessageSuccess(c, "成功", nil)
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
	UIdForm := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}
	result, err := (&workList.VideoService{}).VideoDetail(c, UIdForm.Id)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.Success(c, result)
}

// @Summary video list
// @Description video list
// @Tags Video
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /video/list [post]
func videoList(c *gin.Context) {
	params := &forms.VideoListForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	result, err := (&workList.VideoService{}).VideoList(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.Success(c, result)
}

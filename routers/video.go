package routers

import (
	"video_server/forms"
	"video_server/pkg/response"
	"video_server/pkg/utils"
	"video_server/services"

	"github.com/gin-gonic/gin"
)

func videoUploadImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	result, err := (&services.Service{}).VideoUploadImg(c, file)
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

	result, err := (&services.Service{}).VideoUploadVid(c, file)
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
	id, err := (&services.Service{}).VideoInsert(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.Success(c, id)
}

func videoDelete(c *gin.Context) {
	UIdForm := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if err := (&services.Service{}).VideoDelete(c, UIdForm.Id); err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.MessageSuccess(c, "成功", nil)
}

func videoDetail(c *gin.Context) {
	UIdForm := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}
	result, err := (&services.Service{}).VideoDetail(c, UIdForm.Id)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.Success(c, result)
}

func videoList(c *gin.Context) {
	params := &forms.VideoListForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Size == 0 {
		params.Size = 10
	}
	result, err := (&services.Service{}).VideoList(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.Success(c, result)
}

func searchVideo(c *gin.Context) {
	params := &forms.SearchForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Size <= 0 {
		params.Size = 10
	}
	result, err := (&services.Service{}).SearchVideo(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}

	response.Success(c, result)
}

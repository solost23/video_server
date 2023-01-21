package routers

import (
	"video_server/forms"
	"video_server/pkg/response"
	"video_server/pkg/utils"
	"video_server/services"

	"github.com/gin-gonic/gin"
)

func commentCreate(c *gin.Context) {
	params := &forms.CommentCreateForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if err := (&services.Service{}).CommentInsert(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.MessageSuccess(c, "成功", nil)
}

func commentDelete(c *gin.Context) {
	UIdForm := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if err := (&services.Service{}).CommentDelete(c, UIdForm.Id); err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.MessageSuccess(c, "成功", nil)
}

func commentList(c *gin.Context) {
	params := &forms.CommentListForm{}
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

	result, err := (&services.Service{}).CommentList(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.Success(c, result)
}

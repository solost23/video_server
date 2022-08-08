package router

import (
	"video_server/forms"
	"video_server/pkg/response"
	"video_server/pkg/utils"
	"video_server/workList"

	"github.com/gin-gonic/gin"
)

// @Summary create comment
// @Description add comment
// @Tags Comment
// @Security ApiKeyAuth
// @Param data body models.Comment true "评论"
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/create [post]
func commentCreate(c *gin.Context) {
	params := &forms.CommentCreateForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if err := (&workList.CommentService{}).CommentInsert(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.MessageSuccess(c, "成功", nil)
}

// @Summary delete comment
// @Description delete comment
// @Tags Comment
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/delete [post]
func commentDelete(c *gin.Context) {
	UIdForm := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UIdForm); err != nil {
		response.Error(c, 2001, err)
		return
	}
	if err := (&workList.CommentService{}).CommentDelete(c, UIdForm.Id); err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.MessageSuccess(c, "成功", nil)
}

// @Summary get_comment_by_video_id
// @Description get comment by video id
// @Tags Comment
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/list [post]
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

	result, err := (&workList.CommentService{}).CommentList(c, params)
	if err != nil {
		response.Error(c, 2001, err)
		return
	}
	response.Success(c, result)
}

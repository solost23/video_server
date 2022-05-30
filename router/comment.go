package router

import (
	"github.com/gin-gonic/gin"

	create_comment "video_server/workList/comment/create"
	delete_comment "video_server/workList/comment/delete"
	list_comment "video_server/workList/comment/list"
)

// @Summary create comment
// @Description add comment
// @Tags Comment
// @Security ApiKeyAuth
// @Param data body model.Comment true "评论"
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/create [post]
func createComment(c *gin.Context) {
	request := &create_comment.Request{}
	if err := c.ShouldBind(&request); err != nil {
		Render(c, err)
		return
	}
	data, err := create_comment.NewActionWithCtx(c).Deal(request)
	if err != nil {
		Render(c, err)
		return
	}
	Render(c, err, data)
}

// @Summary delete comment
// @Description delete comment
// @Tags Comment
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/delete [post]
func deleteComment(c *gin.Context) {
	request := &delete_comment.Request{}
	if err := c.ShouldBind(&request); err != nil {
		Render(c, err)
		return
	}
	data, err := delete_comment.NewActionWithCtx(c).Deal(request)
	if err != nil {
		Render(c, err)
		return
	}
	Render(c, err, data)
}

// @Summary get_comment_by_video_id
// @Description get comment by video id
// @Tags Comment
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/list [post]
func getCommentByVideoID(c *gin.Context) {
	request := &list_comment.Request{}
	if err := c.ShouldBind(&request); err != nil {
		Render(c, err)
		return
	}
	data, err := list_comment.NewActionWithCtx(c).Deal(request)
	if err != nil {
		Render(c, err)
		return
	}
	Render(c, err, data)
}

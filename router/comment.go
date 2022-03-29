package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"video_server/pkg/model"
	"video_server/workList"
)

// @Summary create comment
// @Description add comment
// @Tags Comment
// @Security ApiKeyAuth
// @Param data body model.Comment true "评论"
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/{video_id} [post]
func createComment(c *gin.Context) {
	var comment = new(model.Comment)
	if err := c.ShouldBind(comment); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := workList.NewWorkList(c).CreateComment(comment); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, comment)
	return
}

// @Summary delete comment
// @Description delete comment
// @Tags Comment
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/{video_id}/{comment_id} [delete]
func deleteComment(c *gin.Context) {
	var comment = new(model.Comment)
	if err := workList.NewWorkList(c).DeleteComment(comment); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, comment)
	return
}

// @Summary get_comment_by_video_id
// @Description get comment by video id
// @Tags Comment
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/{video_id} [get]
func getCommentByVideoID(c *gin.Context) {
	var comment = new(model.Comment)
	comments, err := workList.NewWorkList(c).GetCommentByVideoID(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, comments)
	return
}

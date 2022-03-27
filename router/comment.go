package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"video_server/pkg/model"
	"video_server/workList"
)

// PingExample godoc
// @Summary ping comment
// @Schemes
// @Description add comment
// @Tags Comment
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

// PingExample godoc
// @Summary ping comment
// @Schemes
// @Description delete comment
// @Tags Comment
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

// PingExample godoc
// @Summary ping comment
// @Schemes
// @Description get comment by video_id
// @Tags Comment
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

package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"video_server/pkg/model"
	"video_server/workList"
)

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

func deleteComment(c *gin.Context) {
	var comment = new(model.Comment)
	if err := workList.NewWorkList(c).DeleteComment(comment); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, comment)
	return
}

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

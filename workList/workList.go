package workList

import "github.com/gin-gonic/gin"

type WorkList struct {
	ctx *gin.Context
}

func NewWorkList(ctx *gin.Context) *WorkList {
	return &WorkList{
		ctx: ctx,
	}
}

package forms

import "video_server/pkg/utils"

type CommentCreateForm struct {
	VideoID  *uint   `json:"videoId" binding:"required"`
	Content  *string `json:"content" binding:"required"`
	ParentID *uint   `json:"parentId"`
	Type     *uint   `json:"type" binding:"required,oneof=0 1" comment:"是否评论 0-点赞 1-评论"`
}

type CommentListForm struct {
	*utils.PageForm
	VideoId *uint `form:"videoId" binding:"required"`
}

type CommentListRecord struct {
	Id            *uint   `json:"id"`
	Content       *string `json:"content"`
	ParentId      *uint   `json:"parentId"`
	Type          *uint   `json:"type"`
	CreatedAt     *string `json:"createdAt"`
	UpdatedAt     *string `json:"updatedAt"`
	CreatorId     *uint   `json:"creatorId"`
	CreatorAvatar *string `json:"creatorAvatar"`
}

type CommentListResponse struct {
	Records  []*CommentListRecord `json:"records"`
	PageList *utils.PageList
}

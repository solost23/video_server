package forms

import "video_server/pkg/utils"

type CommentCreateForm struct {
	VideoID  uint   `json:"videoId" binding:"required"`
	Content  string `json:"content" binding:"required"`
	ParentID uint   `json:"parentId"`
	ISThumb  string `json:"ISThumb" binding:"required,oneof=ISTHUMB,ISCOMMENT"`
}

type CommentListForm struct {
	utils.PageForm
	VideoId uint `form:"videoId" binding:"required"`
}

type CommentListRecord struct {
	Id          uint   `json:"id"`
	Content     string `json:"content"`
	ParentId    uint   `json:"parentId"`
	ISThumb     string `json:"ISThumb"`
	CreatedAt   string `json:"createdAt"`
	UpdatedTime string `json:"updatedTime"`
}

type CommentListResponse struct {
	Records  []CommentListRecord `json:"records"`
	PageList *utils.PageList
}

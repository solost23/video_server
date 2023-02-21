package forms

import (
	"video_server/pkg/utils"
)

type CategoryInsertForm struct {
	Title     *string `json:"title" binging:"required,max=20"`
	Introduce *string `json:"introduce"`
}

type CategoryListForm struct {
	*utils.PageForm
	UserID    *uint   `json:"userId"`
	Title     *string `json:"title"`
	Introduce *string `json:"introduce"`
}

type CategoryListResponse struct {
	PageList *utils.PageList
	Records  []*CategoryListRecord `json:"records"`
}

type CategoryListRecord struct {
	Id        *uint   `json:"id"`
	Title     *string `json:"title"`
	Introduce *string `json:"introduce"`
}

type CategoryUpdateForm struct {
	Title     *string `json:"title"`
	Introduce *string `json:"introduce"`
}

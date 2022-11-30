package models

import (
	"fmt"
	"gorm.io/gorm"
)

type ListPageInput struct {
	Page int `comment:"当前页"`
	Size int `comment:"页长"`
}

type PageForm struct {
	UsePage bool `json:"-"`
	Page    int  `json:"page"`
	Size    int  `json:"pageSize"`
}

type PageList struct {
	Size    int   `json:"size"`
	Pages   int64 `json:"pages"`
	Total   int64 `json:"total"`
	Current int   `json:"current"`
}

// 生成模糊匹配字符串
func LikeFilter(value interface{}) string {
	return fmt.Sprintf("%%%v%%", value)
}

type CreatorBase struct {
	gorm.Model
	CreatorId uint `json:"creatorId" gorm:"column:creator_id;type:bigint unsigned;comment: 创建人ID"`
}

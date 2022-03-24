package model

type Comment struct {
	ID       string `gorm:"id;primary_key"`
	UserID   string `gorm:"user_id"`
	VideoID  string `gorm:"video_id"`
	Content  string `gorm:"content"`
	ParentID string `gorm:"parent_id"`
	ISThumb  string `gorm:"is_thumb;type:enum('ISTHUMB','ISCOMMENT');default:ISTHUMB"`
	// DeleteStatus string `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime int64 `gorm:"create_time"`
	UpdateTime int64 `gorm:"update_time"`
}

func (c *Comment) TableName() string {
	return "comment"
}

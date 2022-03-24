package model

type Video struct {
	ID           string `gorm:"id:primary_key"`
	UserID       string `gorm:"user_id"`
	ClassID      string `gorm:"class_id"`
	Title        string `gorm:"title"`
	Introduce    string `gorm:"introduce"`
	VideoUrl     string `gorm:"video_url"`
	ThumbCount   int64  `gorm:"thumb_count;default:0"`
	CommentCount int64  `gorm:"comment_count;default:0"`
	DeleteStatus string `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime   int64  `gorm:"create_time"`
	UpdateTime   int64  `gorm:"update_time"`
}

func (v *Video) TableName() string {
	return "video"
}

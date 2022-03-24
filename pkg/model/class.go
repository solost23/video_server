package model

type Class struct {
	ID        string `gorm:"id;primary_key"`
	UserID    string `gorm:"user_id"`
	Title     string `gorm:"title"`
	Introduce string `gorm:"introduce"`
	// DeleteStatus string `gorm:"type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime int64 `gorm:"create_time"`
	UpdateTime int64 `gorm:"update_time"`
}

func (c *Class) TableName() string {
	return "class"
}

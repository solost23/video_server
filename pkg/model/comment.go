package model

type Comment struct {
	ID       string `gorm:"id;primary_key"`
	VideoID  string `gorm:"video_id"`
	Content  string `gorm:"content" json:"content"`
	ParentID string `gorm:"parent_id" json:"parentId"`
	ISThumb  string `gorm:"is_thumb;type:enum('ISTHUMB','ISCOMMENT');default:ISTHUMB" json:"isThumb"`
	// DeleteStatus string `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime int64 `gorm:"create_time"`
	UpdateTime int64 `gorm:"update_time"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) Create() error {
	c.ID = NewUUID()
	c.CreateTime = GetCurrentTime()
	c.UpdateTime = GetCurrentTime()
	if err := dbConn.Table(c.TableName()).Create(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) Delete(commentID string) error {
	if err := dbConn.Table(c.TableName()).Where("id=?", commentID).Delete(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) DeleteByVideoIDAndCommentID(videoID, commentID string) error {
	if err := dbConn.Table(c.TableName()).Where("id=? AND video_id=?", commentID, videoID).Delete(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) FindByVideoID(videoID string) (res []*Comment, err error) {
	if err = dbConn.Table(c.TableName()).Where("video_id=?", videoID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (c *Comment) FindByID(commentID string) error {
	if err := dbConn.Table(c.TableName()).Where("id=?", commentID).First(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) DeleteByVideoID(videoID string) error {
	if err := dbConn.Table(c.TableName()).Where("video_id=?", videoID).Delete(c).Error; err != nil {
		return err
	}
	return nil
}

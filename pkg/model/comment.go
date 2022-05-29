package model

import "gorm.io/gorm"

type Comment struct {
	conn     *gorm.DB `gorm:"_"`
	ID       string   `gorm:"id;primary_key"`
	VideoID  string   `gorm:"video_id"`
	Content  string   `gorm:"content" json:"content"`
	ParentID string   `gorm:"parent_id" json:"parentId"`
	ISThumb  string   `gorm:"is_thumb;type:enum('ISTHUMB','ISCOMMENT');default:ISTHUMB" json:"isThumb"`
	// DeleteStatus string `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime int64 `gorm:"create_time"`
	UpdateTime int64 `gorm:"update_time"`
}

func NewComment(conn *gorm.DB) *Comment {
	return &Comment{
		conn: conn,
	}
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) Connection() *gorm.DB {
	return c.conn.Table(c.TableName())
}

func (c *Comment) Create(data *Comment) (err error) {
	err = c.Connection().Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Comment) Delete(commentID string) (comment *Comment, err error) {
	err = dbConn.Table(c.TableName()).Where("id = ?", commentID).Delete(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (c *Comment) FindByVideoID(videoID string) (comments []*Comment, err error) {
	err = dbConn.Table(c.TableName()).Where("video_id = ?", videoID).Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}

func (c *Comment) DeleteByVideoID(videoID string) (comment *Comment, err error) {
	err = dbConn.Table(c.TableName()).Where("video_id = ?", videoID).Delete(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (c *Comment) FindByID(ID string) (comment *Comment, err error) {
	err = c.Connection().Where("id = ?", ID).First(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

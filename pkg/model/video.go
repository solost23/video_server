package model

import (
	"errors"
	"gorm.io/gorm"
)

type Video struct {
	conn         *gorm.DB `gorm:"_"`
	ID           string   `gorm:"id:primary_key"`
	UserID       string   `gorm:"user_id"`
	ClassID      string   `gorm:"class_id"`
	Title        string   `gorm:"title" json:"title" form:"title"`
	Introduce    string   `gorm:"introduce" json:"introduce" form:"introduce"`
	ImageUrl     string   `gorm:"image_url" json:"image_url" form:"image_url"`
	VideoUrl     string   `gorm:"video_url"`
	ThumbCount   int64    `gorm:"thumb_count;default:0"`
	CommentCount int64    `gorm:"comment_count;default:0"`
	DeleteStatus string   `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime   int64    `gorm:"create_time"`
	UpdateTime   int64    `gorm:"update_time"`
}

func NewVideo(conn *gorm.DB) *Video {
	return &Video{
		conn: conn,
	}
}

func (v *Video) TableName() string {
	return "video"
}

func (v *Video) Connection() *gorm.DB {
	return v.conn.Table(v.TableName())
}

func (v *Video) Create(data *Video) error {
	if err := v.Connection().Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (v *Video) Delete(videoID string) error {
	err := v.Connection().Where("id = ? AND delete_status = ?", videoID, DELETENORMAL).
		Update("delete_status = ?", DELETEDEL).Error
	if err != nil {
		return err
	}
	return nil
}

func (v *Video) FindByVideoID(videoID string, deleteStatus string) (video *Video, err error) {
	err = dbConn.Table(v.TableName()).Where("id = ? AND delete_status = ?", videoID, deleteStatus).
		First(&video).Error
	if err != nil {
		return video, err
	}
	return video, nil
}

func (v *Video) FindByUserIDAndClassID(userID, classID, deleteStatus string) (res []*Video, err error) {
	if err = dbConn.Table(v.TableName()).Where("user_id=? AND class_id=? AND delete_status=?", userID, classID, deleteStatus).Find(&res).Error; err != nil {
		return res, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return res, nil
}

func (v *Video) FindByUserIDANDClassIDAndID(userID, classID, videoID, deleteStatus string) error {
	if err := v.Connection().Where("user_id=? AND class_id=? AND id=? AND delete_status=?", userID, classID, videoID, deleteStatus).First(v).Error; err != nil {
		return err
	}
	return nil
}

func (v *Video) FindByUserID(userID, deleteStatus string) (res []*Video, err error) {
	if err = dbConn.Table(v.TableName()).Where("user_id=? AND delete_status=?", userID, deleteStatus).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (v *Video) FindByClassID(classID string, deleteStatus string) (videos []*Video, err error) {
	if err = dbConn.Table(v.TableName()).Where("class_id=? AND delete_status=?", classID, deleteStatus).Find(&videos).Error; err != nil {
		return videos, err
	}
	return videos, nil
}

func (v *Video) Find(deleteStatus string) (res []*Video, err error) {
	if err = dbConn.Table(v.TableName()).Where("delete_status=?", deleteStatus).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

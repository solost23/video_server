package model

type Video struct {
	ID           string `gorm:"id:primary_key"`
	UserID       string `gorm:"user_id"`
	ClassID      string `gorm:"class_id"`
	Title        string `gorm:"title" json:"title" form:"title"`
	Introduce    string `gorm:"introduce" json:"introduce" form:"introduce"`
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

func (v *Video) Create() error {
	v.ID = NewUUID()
	v.CreateTime = GetCurrentTime()
	v.UpdateTime = GetCurrentTime()
	if err := dbConn.Table(v.TableName()).Create(v).Error; err != nil {
		return err
	}
	return nil
}

func (v *Video) Delete(videoID string) error {
	v.DeleteStatus = "DELETE_STATUS_DEL"
	if err := dbConn.Table(v.TableName()).Where("id=?", videoID).Save(v).Error; err != nil {
		return err
	}
	return nil
}

func (v *Video) FindByVideoID(videoID string, deleteStatus string) error {
	if err := dbConn.Table(v.TableName()).Where("id=? AND delete_status=?", videoID, deleteStatus).First(v).Error; err != nil {
		return err
	}
	return nil
}

func (v *Video) FindByUserIDAndClassID(userID, classID, deleteStatus string) (res []*Video, err error) {
	if err = dbConn.Table(v.TableName()).Where("user_id=? AND class_id=? AND delete_status=?", userID, classID, deleteStatus).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (v *Video) FindByUserIDANDClassIDAndID(userID, classID, videoID, deleteStatus string) error {
	if err := dbConn.Table(v.TableName()).Where("user_id=? AND class_id=? AND id=? AND delete_status=?", userID, classID, videoID, deleteStatus).First(v).Error; err != nil {
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

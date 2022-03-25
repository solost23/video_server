package model

type Class struct {
	ID        string `gorm:"id;primary_key"`
	UserID    string `gorm:"user_id"`
	Title     string `gorm:"title" json:"title"`
	Introduce string `gorm:"introduce" json:"introduce"`
	// DeleteStatus string `gorm:"type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime int64 `gorm:"create_time"`
	UpdateTime int64 `gorm:"update_time"`
}

func (c *Class) TableName() string {
	return "class"
}

// 增加分类
func (c *Class) Create() error {
	c.ID = NewUUID()
	c.CreateTime = GetCurrentTime()
	c.UpdateTime = GetCurrentTime()
	if err := dbConn.Table(c.TableName()).Create(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Class) Delete(id string) error {
	if err := dbConn.Table(c.TableName()).Where("id=?", id).Delete(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Class) Update(id string) error {
	c.UpdateTime = GetCurrentTime()
	if err := dbConn.Table(c.TableName()).Omit("user_id", "create_time").Where("id=?", id).Save(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Class) FindByUserID(userID string) ([]*Class, error) {
	var res []*Class
	if err := dbConn.Table(c.TableName()).Where("user_id=?", userID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (c *Class) FindByUserIDClassTitle(userID, classTitle string) error {
	if err := dbConn.Table(c.TableName()).Where("user_id=? AND title=?", userID, classTitle).First(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Class) FindByID(id string) error {
	if err := dbConn.Table(c.TableName()).Where("id=?", id).First(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Class) DeleteByUserID(userID string) error {
	if err := dbConn.Table(c.TableName()).Where("user_id=?", userID).Delete(c).Error; err != nil {
		return err
	}
	return nil
}

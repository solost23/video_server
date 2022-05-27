package model

import "gorm.io/gorm"

type Category struct {
	conn      *gorm.DB `gorm:"_"`
	ID        string   `gorm:"id;primary_key"`
	UserID    string   `gorm:"user_id"`
	Title     string   `gorm:"title" json:"title"`
	Introduce string   `gorm:"introduce" json:"introduce"`
	// DeleteStatus string `gorm:"type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime int64 `gorm:"create_time"`
	UpdateTime int64 `gorm:"update_time"`
}

func NewCategory(conn *gorm.DB) *Category {
	return &Category{
		conn: conn,
	}
}

func (c *Category) TableName() string {
	return "category"
}

func (c *Category) Connection() *gorm.DB {
	return c.conn.Table(c.TableName())
}

// 增加分类
func (c *Category) Create() error {
	c.ID = NewUUID()
	c.CreateTime = GetCurrentTime()
	c.UpdateTime = GetCurrentTime()
	if err := dbConn.Table(c.TableName()).Create(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) FindByClassName(className string) (class *Category, err error) {
	err = c.Connection().Where("title = ?", className).First(&class).Error
	if err != nil {
		return class, err
	}
	return class, nil
}

func (c *Category) Delete(id string) error {
	if err := dbConn.Table(c.TableName()).Where("id=?", id).Delete(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) Update(id string) error {
	c.UpdateTime = GetCurrentTime()
	if err := dbConn.Table(c.TableName()).Omit("user_id", "create_time").Where("id=?", id).Save(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) FindByUserID(userID string) ([]*Category, error) {
	var res []*Category
	if err := dbConn.Table(c.TableName()).Where("user_id=?", userID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (c *Category) FindByUserIDClassTitle(userID, classTitle string) error {
	if err := dbConn.Table(c.TableName()).Where("user_id=? AND title=?", userID, classTitle).First(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) FindByID(id string) error {
	if err := dbConn.Table(c.TableName()).Where("id=?", id).First(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Category) DeleteByUserID(userID string) error {
	if err := dbConn.Table(c.TableName()).Where("user_id=?", userID).Delete(c).Error; err != nil {
		return err
	}
	return nil
}

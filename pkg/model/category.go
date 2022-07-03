package model

import (
	"errors"

	"gorm.io/gorm"
)

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

func (t *Category) TableName() string {
	return "category"
}

func (t *Category) Connection() *gorm.DB {
	return t.conn.Table(t.TableName())
}

// 增加分类
func (t *Category) Create(data *Category) (err error) {
	err = t.Connection().Create(&data).Error
	return err
}

// 删
func (t *Category) Delete1(query interface{}, args ...interface{}) (err error) {
	err = t.Connection().Where(query, args...).Delete(&t).Error
	return err
}

// 改
func (t *Category) UpdateColumn(key string, value interface{}, query interface{}, args ...interface{}) (err error) {
	return t.Connection().Where(query, args...).Update(key, value).Error
}

// 更新全部数据
func (t *Category) Save(data *Category) (err error) {
	return t.Connection().Save(&data).Error
}

// 查单个数据
func (t *Category) WhereOne(query interface{}, args ...interface{}) (category *Category, err error) {
	err = t.Connection().Where(query, args...).First(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

// 查询所有
func (t *Category) WhereAll(query interface{}, args ...interface{}) (categories []*Category, err error) {
	err = t.Connection().Where(query, args...).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// 支持分页筛选
func (t *Category) PageList(params *ListPageInput, query interface{}, args ...interface{}) (categories []*Category, count int64, err error) {
	offset := (params.Page - 1) * params.Size

	err = t.Connection().Where(query, args...).Offset(offset).Limit(params.Size).Order("create_time DESC").Find(&categories).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	err = t.Connection().Where(query, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return categories, count, nil
}

func (c *Category) FindByClassName(className string) (class *Category, err error) {
	err = c.Connection().Where("title = ?", className).First(&class).Error
	if err != nil {
		return class, err
	}
	return class, nil
}

func (c *Category) Delete(id string) (category *Category, err error) {
	err = dbConn.Table(c.TableName()).Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *Category) Update(data *Category) (err error) {
	err = c.Connection().Omit("user_id", "create_time").Where("id = ?", data.ID).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) FindByUserID(userID string) (categorys []*Category, err error) {
	err = c.Connection().Where("user_id = ?", userID).Find(&categorys).Error
	if err != nil {
		return categorys, err
	}
	return categorys, nil
}

func (c *Category) FindByUserIDClassTitle(userID, classTitle string) (category *Category, err error) {
	err = c.Connection().Where("user_id = ? AND title = ?", userID, classTitle).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *Category) FindByID(id string) (category *Category, err error) {
	err = c.Connection().Where("id = ?", id).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *Category) DeleteByUserID(userID string) (category *Category, err error) {
	err = dbConn.Table(c.TableName()).Where("user_id=?", userID).Delete(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

package model

import "gorm.io/gorm"

type User struct {
	conn         *gorm.DB `gorm:"_"`
	ID           string   `gorm:"id;primary_key"`
	UserName     string   `gorm:"user_name" json:"user_name"`
	Password     string   `gorm:"password" json:"password"`
	Nickname     string   `gorm:"nickname" json:"nickname"`
	Role         string   `gorm:"role;type:enum('ADMIN','USER');default:USER" json:"role"`
	Avatar       string   `gorm:"avatar" json:"avatar"`
	Introduce    string   `gorm:"introduce" json:"introduce"`
	FansCount    int64    `gorm:"fans_count;default:0"`
	CommentCount int64    `gorm:"comment_count;default:0"`
	// DeleteStatus string `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime int64 `gorm:"create_time"`
	UpdateTime int64 `gorm:"update_time"`
}

func NewUser(conn *gorm.DB) *User {
	return &User{
		conn: conn,
	}
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Connection() *gorm.DB {
	return u.conn.Table(u.TableName())
}

// 增加用户
func (u *User) Create(data *User) (err error) {
	err = u.Connection().Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除用户
func (u *User) Delete(id string) (user *User, err error) {
	err = dbConn.Table(u.TableName()).Where("id=?", id).Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// 修改用户信息
func (u *User) Update(data *User) (err error) {
	err = u.Connection().Omit("id", "user_name", "fans_count", "comment_count", "create_time").Where("id=?", data.ID).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

// 显示单个用户信息
func (u *User) FindByID(id string) (user *User, err error) {
	err = u.Connection().Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// 显示所有用户信息
func (u *User) Find() (users []*User, err error) {
	err = dbConn.Table(u.TableName()).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *User) FindBYUserName(userName string) (user *User, err error) {
	err = u.Connection().Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

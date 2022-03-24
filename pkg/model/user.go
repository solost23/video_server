package model

type User struct {
	ID           string `gorm:"id;primary_key"`
	UserName     string `gorm:"user_name" json:"user_name"`
	Password     string `gorm:"password" json:"password"`
	Nickname     string `gorm:"nickname" json:"nickname"`
	Role         string `gorm:"role;type:enum('ADMIN','USER');default:USER" json:"role"`
	Avatar       string `gorm:"avatar" json:"avatar"`
	Introduce    string `gorm:"introduce" json:"introduce"`
	FansCount    int64  `gorm:"fans_count;default:0"`
	CommentCount int64  `gorm:"comment_count;default:0"`
	// DeleteStatus string `gorm:"delete_status;type:enum('DELETE_STATUS_NORMAL','DELETE_STATUS_DEL');default:DELETE_STATUS_NORMAL"`
	CreateTime int64 `gorm:"create_time"`
	UpdateTime int64 `gorm:"update_time"`
}

func (u *User) TableName() string {
	return "user"
}

// 增加用户
func (u *User) Create() error {
	u.ID = NewUUID()
	u.CreateTime = GetCurrentTime()
	u.UpdateTime = GetCurrentTime()
	if err := dbConn.Table(u.TableName()).Create(u).Error; err != nil {
		return err
	}
	return nil
}

// 删除用户
func (u *User) Delete(id string) error {
	if err := dbConn.Table(u.TableName()).Where("id=?", id).Delete(u).Error; err != nil {
		return err
	}
	return nil
}

// 修改用户信息
func (u *User) Update(id string) error {
	u.UpdateTime = GetCurrentTime()
	if err := dbConn.Table(u.TableName()).Omit("id", "user_name", "fans_count", "comment_count", "create_time").Where("id=?", id).Save(u).Error; err != nil {
		return err
	}
	return nil
}

// 显示单个用户信息
func (u *User) FindByID(id string) error {
	if err := dbConn.Table(u.TableName()).Where("id=?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

// 显示所有用户信息
func (u *User) Find() ([]*User, error) {
	var res []*User
	if err := dbConn.Table(u.TableName()).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (u *User) FindBYUserName(userName string) error {
	if err := dbConn.Table(u.TableName()).Where("user_name=?", userName).First(u).Error; err != nil {
		return err
	}
	return nil
}

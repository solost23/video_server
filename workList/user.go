package workList

import (
	"errors"
	"video_server/pkg/middleware"

	"gorm.io/gorm"

	"video_server/pkg/model"
)

func (w *WorkList) Register(user *model.User) error {
	// 检查当前用户是否存在，若存在，则返回错误
	// 若不存在，则创建
	if err := user.FindBYUserName(user.UserName); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	// 说明用户已经存在
	if user.ID != "" {
		return errors.New("user already exist")
	}
	// 判断用户角色
	if user.Role != model.ROLEADMIN && user.Role != model.ROLEUSER {
		return errors.New("role not exist")
	}
	user.Password = model.NewMd5(user.Password, model.SECRET)
	if err := user.Create(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) Login(user *model.User) (string, error) {
	// 查看有无用户，若没有直接报错
	// 若有，检查账户名密码，若有一个为错，则返回
	// 否则生成一个token
	userName := user.UserName
	role := user.Role
	password := user.Password
	if err := user.FindBYUserName(userName); err != nil {
		return "", err
	}
	if userName != user.UserName || model.NewMd5(password, model.SECRET) != user.Password {
		return "", errors.New("username or password err")
	}
	tokenStr, err := middleware.CreateToken(userName, role)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

func (w *WorkList) GetUserInfo(user *model.User) error {
	// 直接查找用户信息，找不到报错
	userName := w.ctx.Param("user_name")
	if err := user.FindBYUserName(userName); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) UpdateUserInfo(user *model.User) error {
	// 检查用户是否存在，若用户存在，则更新用户信息
	userName := w.ctx.Param("user_name")
	var tmpUser = new(model.User)
	if err := tmpUser.FindBYUserName(userName); err != nil {
		return err
	}
	if err := user.Update(tmpUser.ID); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteUserInfo(user *model.User) error {
	// 查看用户是否存在，若不存在，则返回错误
	// 若存在，则按照id删除
	userName := w.ctx.Param("user_name")
	if err := user.FindBYUserName(userName); err != nil {
		return err
	}
	if err := user.Delete(user.ID); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GETAllUserInfo(user *model.User) ([]*model.User, error) {
	// 直接查询所有用户
	users, err := user.Find()
	if err != nil {
		return users, err
	}
	return users, nil
}

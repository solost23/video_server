package workList

import (
	"errors"

	"gorm.io/gorm"

	"video_server/pkg/model"
)

func (w *WorkList) CreateClass(class *model.Class) error {
	// 先查询用户是否存在，再查询用户下此分类是否存在
	userName := w.ctx.Param("user_name")
	var user = new(model.User)
	if err := user.FindBYUserName(userName); err != nil {
		return err
	}
	var tmpClass = new(model.Class)
	if err := tmpClass.FindByUserIDClassTitle(user.ID, class.Title); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if tmpClass.ID != "" {
		return errors.New("user class exist")
	}
	class.UserID = user.ID
	if err := class.Create(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) UpdateClass(class *model.Class) error {
	// 校验用户是否存在，校验类型是否存在
	userName := w.ctx.Param("user_name")
	classID := w.ctx.Param("class_id")
	var user = new(model.User)
	if err := user.FindBYUserName(userName); err != nil {
		return err
	}
	var tmpClass = new(model.Class)
	if err := tmpClass.FindByID(classID); err != nil {
		return err
	}
	class.ID = classID
	class.UserID = user.ID
	class.CreateTime = tmpClass.CreateTime
	if err := class.Update(tmpClass.ID); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetUserAllClass(class *model.Class) (classes []*model.Class, err error) {
	// 通过用户名找到用户id,再通过id去获取class
	userName := w.ctx.Param("user_name")
	var user = new(model.User)
	if err = user.FindBYUserName(userName); err != nil {
		return classes, err
	}
	classes, err = class.FindByUserID(user.ID)
	if err != nil {
		return classes, err
	}
	return classes, nil
}

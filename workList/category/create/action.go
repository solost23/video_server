package create

import (
	"github.com/gin-gonic/gin"
	"video_server/workList"
)

type Action struct {
	workList.WorkList
}

func NewActionWithCtx(ctx *gin.Context) *Action {
	r := &Action{}
	r.Init(ctx)
	return r
}

func (a *Action) Deal(request *Request) (resp Response, err error) {
	return resp, err
}

//func (w *workList.WorkList) CreateClass(class *model.Class) error {
//	// 先查询用户是否存在，再查询用户下此分类是否存在
//	userName := w.ctx.Param("user_name")
//	var user = new(model.User)
//	if err := user.FindBYUserName(userName); err != nil {
//		return err
//	}
//	var tmpClass = new(model.Class)
//	if err := tmpClass.FindByUserIDClassTitle(user.ID, class.Title); err != nil {
//		if !errors.Is(err, gorm.ErrRecordNotFound) {
//			return err
//		}
//	}
//	if tmpClass.ID != "" {
//		return errors.New("user category exist")
//	}
//	class.UserID = user.ID
//	if err := class.Create(); err != nil {
//		return err
//	}
//	return nil
//}

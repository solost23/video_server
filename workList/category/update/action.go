package update

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

//func (w *workList.WorkList) UpdateClass(class *model.Class) error {
//	// 校验用户是否存在，校验类型是否存在
//	userName := w.ctx.Param("user_name")
//	classID := w.ctx.Param("class_id")
//	var user = new(model.User)
//	if err := user.FindBYUserName(userName); err != nil {
//		return err
//	}
//	var tmpClass = new(model.Class)
//	if err := tmpClass.FindByID(classID); err != nil {
//		return err
//	}
//	class.ID = classID
//	class.UserID = user.ID
//	class.CreateTime = tmpClass.CreateTime
//	if err := class.Update(tmpClass.ID); err != nil {
//		return err
//	}
//	return nil
//}

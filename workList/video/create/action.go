package create

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"os"
	"video_server/pkg/model"
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
	// 视频目录结构 video/user_name/class_name/filename
	// 拿到用户名，分类id，查询此文件是否存在，若不存在，则创建
	// 然后存储视频文件
	// 封装视频结构体，存储视频信息

	// 校验参数，查看该用户下是否有此分类
	if err = model.NewUser(a.GetMysqlConn()).FindByID(request.UserID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, err
		}
		return resp, err
	}
	if err = model.NewClass(a.GetMysqlConn()).FindByID(request.ClassID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, err
		}
		return resp, err
	}
	// 文件暂时保存在本地
	file, header, err := a.Ctx.Request.FormFile("file")
	if err != nil {
		return resp, err
	}
	filePath := fmt.Sprintf("%s/%s/%s", model.FilePath, request.UserID, request.ClassID)
	fileName := fmt.Sprintf("%s/%s/%s/%s", model.FilePath, request.UserID, request.ClassID, header.Filename)
	// 判断此文件是否存在,若存在，则直接返回错误
	if Exist(fileName) {
		return resp, errors.New("此视频已存在，请勿重复创建")
	}
	// 不存在，则先创建文件夹，再创建文件
	// 判断文件夹是否存在，若不存在，则创建
	if !Exist(filePath) {
		if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
			return resp, err
		}
	}
	fp, err := os.Create(fileName)
	if err != nil {
		return resp, err
	}
	defer func() {
		fp.Close()
	}()
	_, err = io.Copy(fp, file)
	if err != nil {
		return resp, err
	}

	// 封装视频结构体，保存数据
	if err = model.NewVideo(a.GetMysqlConn()).Create(a.buildCreateVideo(request, fileName)); err != nil {
		return resp, err
	}
	return resp, nil
}

func (a *Action) buildCreateVideo(request *Request, fileName string) (data *model.Video) {
	// 视频网址暂时采用本地的，后面改为前段上传
	data = &model.Video{
		ID:           model.NewUUID(),
		UserID:       request.UserID,
		ClassID:      request.ClassID,
		Title:        request.Title,
		Introduce:    request.Introduce,
		ImageUrl:     request.ImageUrl,
		VideoUrl:     fmt.Sprintf("http://127.0.0.1:8080/%s", fileName),
		ThumbCount:   0,
		CommentCount: 0,
		DeleteStatus: "DELETE_STATUS_NORMAL",
		CreateTime:   model.GetCurrentTime(),
		UpdateTime:   model.GetCurrentTime(),
	}
	return data
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

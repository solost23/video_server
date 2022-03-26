package workList

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gorm.io/gorm"

	"video_server/pkg/model"
)

func (w *WorkList) CreateVideo(video *model.Video) error {
	// 视频目录结构 video/user_name/class_name/filename
	// 拿到用户名，分类id，查询此文件是否存在，若不存在，则创建
	// 然后存储视频文件
	// 封装视频结构体，存储视频信息
	userName := w.ctx.Param("user_name")
	classID := w.ctx.Param("class_id")
	// 校验参数，查看该用户下是否有此分类
	var user = new(model.User)
	if err := user.FindBYUserName(userName); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	var class = new(model.Class)
	if err := class.FindByID(classID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	file, header, err := w.ctx.Request.FormFile("file")
	if err != nil {
		return err
	}
	filepath := fmt.Sprintf("%s/%s/%s", model.FilePath, userName, class.Title)
	filename := fmt.Sprintf("%s/%s/%s/%s", model.FilePath, userName, class.Title, header.Filename)
	// 判断此文件是否存在,若存在，则直接返回错误
	if Exist(filename) {
		return errors.New("此视频已存在，请勿重复创建")
	}
	// 不存在，则先创建文件夹，再创建文件
	// 判断文件夹是否存在，若不存在，则创建
	// 创建文件
	if !Exist(filepath) {
		if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
			return err
		}
	}
	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		fp.Close()
	}()
	_, err = io.Copy(fp, file)
	if err != nil {
		return err
	}

	// 封装视频结构体，保存数据
	video.UserID = user.ID
	video.ClassID = class.ID
	// 后期写到配置文件中
	video.VideoUrl = fmt.Sprintf("http://127.0.0.1:8080/%s", filename)
	if err := video.Create(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteVideo(video *model.Video) error {
	// 先查找视频信息,找不到则报错
	// 给视频的delete_status打上标记就可以
	// 注意:删视频的时候，视频下面的评论也要删除
	userName := w.ctx.Param("user_name")
	classID := w.ctx.Param("class_id")
	videoID := w.ctx.Param("video_id")
	var user = new(model.User)
	if err := user.FindBYUserName(userName); err != nil {
		return err
	}
	// 直接通过用户id,分类id，视频id查就可以
	if err := video.FindByUserIDANDClassIDAndID(user.ID, classID, videoID, model.DELETENORMAL); err != nil {
		return err
	}
	// 删视频
	if err := video.Delete(video.ID); err != nil {
		return err
	}
	// 删视频下的评论
	var comment = new(model.Comment)
	// 查看一下视频下是否有评论
	comments, err := comment.FindByVideoID(videoID)
	if err != nil {
		return err
	}
	if len(comments) <= 0 {
		return nil
	}
	if err := comment.DeleteByVideoID(videoID); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetVideo(video *model.Video) error {
	// 直接获取视频信息
	videoID := w.ctx.Param("video_id")
	if err := video.FindByVideoID(videoID, model.DELETENORMAL); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetVideoByUserNameAndClassID(video *model.Video) (videos []*model.Video, err error) {
	userName := w.ctx.Param("user_name")
	classID := w.ctx.Param("class_id")
	// 直接筛选
	var user = new(model.User)
	if err = user.FindBYUserName(userName); err != nil {
		return videos, err
	}
	videos, err = video.FindByUserIDAndClassID(user.ID, classID, model.DELETENORMAL)
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func (w *WorkList) GetVideoByUserName(video *model.Video) (videos []*model.Video, err error) {
	userName := w.ctx.Param("user_name")
	var user = new(model.User)
	if err = user.FindBYUserName(userName); err != nil {
		return videos, err
	}
	videos, err = video.FindByUserID(user.ID, model.DELETENORMAL)
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func (w *WorkList) GetAllVideo(video *model.Video) (videos []*model.Video, err error) {
	// 直接获取并返回
	videos, err = video.Find(model.DELETENORMAL)
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

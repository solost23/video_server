package deleteVideo

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/url"
	"os"
	"video_server/global"

	"video_server/pkg/models"
)

type Task struct {
	Ctx  context.Context
	Conn *gorm.DB
}

func (t *Task) DeleteVideo() {
	t.Ctx = context.Background()
	var err error
	t.Conn = global.DB
	if err = t.Deal(); err != nil {
		log.Printf("delete video: %v \n", err.Error())
	}
}

func (t *Task) Deal() (err error) {
	// 去数据库查询要删除的视频数据
	// 删除数据库视频数据
	// 删除本地视频数据
	var video = new(models.Video)
	var videos []*models.Video
	if err = t.Conn.Table(video.TableName()).Where("deleted_at != NULL").Find(&videos).Error; err != nil {
		return err
	}
	if len(videos) <= 0 {
		return errors.New("不存在要删除视频")
	}
	// 删除数据库中记录
	if err = t.Conn.Table(video.TableName()).Where("deleted_at != NULL").Delete(video).Error; err != nil {
		return err
	}
	var videoFilePath []string
	for _, item := range videos {
		videoUrl, _ := url.Parse(item.VideoUrl)
		videoFileName := videoUrl.Path
		videoFilePath = append(videoFilePath, videoFileName[1:len(videoFileName)])
	}
	// 根据路径去删除视频文件
	if err = t.DeleteVideoFile(videoFilePath); err != nil {
		return err
	}
	return nil
}

func (t *Task) DeleteVideoFile(videoFilePath []string) (err error) {
	ch := make(chan bool, len(videoFilePath))
	for _, videoFileName := range videoFilePath {
		// 删除文件
		go func(videoFileName string) {
			ch <- true
			os.Remove(videoFileName)
			fmt.Printf("删除文件%s成功 \n", videoFileName)
		}(videoFileName)
	}
	for i := 0; i < len(videoFilePath); i++ {
		<-ch
	}
	return nil
}

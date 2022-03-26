package scheduler

import (
	"github.com/robfig/cron/v3"

	"video_server/scheduler/deleteVideo"
)

// 删除数据库中视频数据记录
// 删除本地视频文件
func Run() {
	var task = new(deleteVideo.Task)
	c := cron.New()
	// 后期写入配置
	c.AddFunc("0 0 3 * * ?", func() {
		task.DeleteVideo()
	})
	c.Start()
	select {}
}

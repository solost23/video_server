package scheduler

import (
	"github.com/robfig/cron/v3"
	"video_server/config"

	"video_server/scheduler/deleteVideo"
)

// 删除数据库中视频数据记录
// 删除本地视频文件
func Run() {
	var task = new(deleteVideo.Task)
	c := cron.New()
	schedulerConfig := config.NewScheduler()
	c.AddFunc(schedulerConfig.CronTime, func() {
		task.DeleteVideo()
	})
	c.Start()
	select {}
}

package main

import (
	"fmt"
	"net/http"
	"time"
	"video_server/config"

	"video_server/router"
	"video_server/scheduler"
)

// @title video_server Swagger
// @version 1.0
// @description this is a video server
// @host localhost:8080
// @schemes http https
// @BasePath /
func main() {
	// 启动定时任务，每天晚上三点删除视频文件，降低用户删除请求的io操作过多
	fmt.Println("start scheduler")
	go scheduler.Run()
	service := config.NewProject()
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", service.ServiceAddr, service.ServicePort),
		Handler:      router.InitRouter(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}

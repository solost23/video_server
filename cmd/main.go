package main

import (
	"fmt"
	"net/http"
	"time"

	"video_server/router"
	"video_server/scheduler"
)

func main() {
	// 启动定时任务，每天晚上三点删除视频文件，降低用户删除请求的io操作过多
	fmt.Println("start scheduler")
	go scheduler.Run()
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router.InitRouter(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}

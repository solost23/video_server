package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"video_server/global"
	"video_server/global/initialize"
	"video_server/router"
	"video_server/scheduler"
)

func Run() {
	// 启动定时任务，每天晚上三点删除视频文件，降低用户删除请求的io操作过多
	go func() {
		fmt.Println("start scheduler")
		scheduler.Run()
	}()

	initialize.Initialize("./config/config.yml")
	server := &http.Server{
		Addr:         global.ServerConfig.Addr,
		Handler:      router.InitRouter(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s \n", err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			_ = server.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

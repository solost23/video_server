package server

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"video_server/config"
	"video_server/router"
	"video_server/scheduler"
)

func Run() {
	fmt.Println(viper.ConfigFileUsed())
	port := viper.Get("params.service_port").(int)
	// 启动定时任务，每天晚上三点删除视频文件，降低用户删除请求的io操作过多
	fmt.Println("start scheduler")
	go scheduler.Run()
	service := config.NewProject()
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", service.ServiceAddr, port),
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
			server.Close()
			time.Sleep(3 * time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}

package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"video/global"
	"video/global/initialize"
	"video/pkg/models"
	"video/routers"
	"video/scheduler"
)

func Run() {
	initialize.Initialize("./configs/config.yml")
	// 创建 model
	err := autoMigrate(global.DB, &models.CasbinRule{}, &models.Category{}, &models.Comment{}, &models.User{}, &models.UserComment{}, &models.Video{})
	if err != nil {
		panic(err)
	}
	// Version
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("video version: %s\n", global.ServerConfig.Version)
		os.Exit(0)
	}

	// HTTP init
	app := gin.New()
	routers.Setup(app)

	server := &http.Server{
		Addr:         global.ServerConfig.Addr,
		Handler:      app,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	// 启动定时任务，每天晚上三点删除视频文件，降低用户删除请求的io操作过多
	go func() {
		fmt.Println("start scheduler")
		scheduler.Run()
	}()
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			if err == http.ErrServerClosed {
				log.Panicf("Server closed")
			} else {
				log.Panicf("Server closed unexpect %s", err.Error())
			}
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

func autoMigrate(db *gorm.DB, dst ...interface{}) error {
	return db.AutoMigrate(dst...)
}

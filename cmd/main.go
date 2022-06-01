package main

import (
	"flag"
	"fmt"
	video_server_logger "github.com/solost23/tools/log"
	"github.com/spf13/viper"
	"os"
	"path"
	"video_server/server"
)

var (
	WebConfigPath = "config/config.yaml"
	version       = "__BUILD_VERSION__"
	execDir       string
	st, v, V      bool
)

// @title video_server Swagger
// @version 1.0
// @description this is a video server
// @host localhost:8080
// @schemes http https
// @BasePath /
func main() {
	flag.StringVar(&execDir, "d", ".", "项目目录")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本号")
	flag.BoolVar(&st, "s", false, "项目状态")
	flag.Parse()
	if v || V {
		fmt.Println(version)
		return
	}
	// 运行
	InitConfig()
	InitLogger()
	server.Run()
}

func InitConfig() {
	configPath := path.Join(execDir, WebConfigPath)
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("未找到配置文件，当前path:", configPath)
		os.Exit(1)
	}
}

func InitLogger() {
	// 默认已经正常load config到var cfg configs.Config
	// 使用自定义的log
	logger := video_server_logger.NewLogger(viper.GetString("log.runtime.path"))
	if logger == nil {
		fmt.Println("init logger failed")
		os.Exit(1)
	}
	ctxLogger := video_server_logger.NewLogger(viper.GetString("log.track.path"))
	if ctxLogger == nil {
		fmt.Println("init ctxLogger failed")
		os.Exit(1)
	}
}

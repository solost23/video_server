package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	Version    = viper.GetString("version")
	Video_path = viper.GetString("video.video_path")
	JwtKey     = viper.GetString("jwt.key")
	Md5        = viper.GetString("md5.secret")
)

type Project struct {
	ServiceName string
	ServiceAddr string
	ServicePort string
}

func NewProject() *Project {
	return &Project{
		ServiceName: viper.GetString("params.service_name"),
		ServicePort: viper.GetString("params.service_port"),
		ServiceAddr: viper.GetString("params.service_addr"),
	}
}

type Connections struct {
	Host     string
	UserName string
	Password string
	Port     string
	DB       string
	CasbinDB string
	Charset  string
}

func NewConnections() *Connections {
	return &Connections{
		Host:     viper.GetString("connections.mysql.video_server.host"),
		UserName: viper.GetString("connections.mysql.video_server.user"),
		Password: viper.GetString("connections.mysql.video_server.password"),
		Port:     viper.GetString("connections.mysql.video_server.port"),
		DB:       viper.GetString("connections.mysql.video_server.db"),
		CasbinDB: viper.GetString("connections.mysql.video_server.casbin_db"),
		Charset:  viper.GetString("connections.mysql.video_server.charset"),
	}
}

type Scheduler struct {
	CronTime string
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		CronTime: viper.GetString("scheduler.delete_video.cron_time"),
	}
}

func init() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("config read error")
	}
}

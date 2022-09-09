package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Version = viper.GetString("version")
	Md5     = viper.GetString("md5.secret")
)

type JWTConfig struct {
	Key      string
	Duration int64
}

func NewJWTConfig() *JWTConfig {
	return &JWTConfig{
		Key:      viper.GetString("jwt.key"),
		Duration: viper.GetInt64("jwt.duration"),
	}
}

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

type MysqlConfig struct {
	Host     string
	UserName string
	Password string
	Port     string
	DB       string
	CasbinDB string
	Charset  string
}

func NewMysqlConfig() *MysqlConfig {
	return &MysqlConfig{
		Host:     viper.GetString("connections.mysql.video_server.host"),
		UserName: viper.GetString("connections.mysql.video_server.user"),
		Password: viper.GetString("connections.mysql.video_server.password"),
		Port:     viper.GetString("connections.mysql.video_server.port"),
		DB:       viper.GetString("connections.mysql.video_server.db"),
		CasbinDB: viper.GetString("connections.mysql.video_server.casbin_db"),
		Charset:  viper.GetString("connections.mysql.video_server.charset"),
	}
}

type RedisConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host:     viper.GetString("connections.redis.video_server.host"),
		Port:     viper.GetString("connections.redis.video_server.port"),
		UserName: viper.GetString("connections.redis.video_server.user"),
		Password: viper.GetString("connections.redis.video_server.password"),
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

type MinioConfig struct {
	EndPoint        string
	AccessKeyID     string
	SecretAccessKey string
	UserSSL         bool
}

func NewMinio() *MinioConfig {
	return &MinioConfig{
		EndPoint:        viper.GetString("minio.end_point"),
		AccessKeyID:     viper.GetString("minio.access_key_id"),
		SecretAccessKey: viper.GetString("minio.secret_access_key"),
		UserSSL:         viper.GetBool("minio.user_ssl"),
	}
}

func init() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("config read error")
	}
}

package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go"
	"gorm.io/gorm"
	"video_server/config"
)

var (
	ServerConfig = &config.ServerConfig{}
	DB           *gorm.DB
	CasbinDB     *gorm.DB
	Minio        *minio.Client
	RedisMapPool = make(map[int]*redis.Client, 15)
)

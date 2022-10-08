package cache

import (
	"context"
	"video_server/global"

	"github.com/go-redis/redis/v8"
)

func RedisConnFactory(db int) (*redis.Client, error) {
	if global.RedisMapPool[db] != nil {
		return global.RedisMapPool[db], nil
	}
	global.RedisMapPool[db] = redis.NewClient(&redis.Options{
		Addr:     global.ServerConfig.RedisConfig.Addr,
		Password: global.ServerConfig.RedisConfig.Password,
		DB:       db,
	})

	_, err := global.RedisMapPool[db].Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return global.RedisMapPool[db], nil
}

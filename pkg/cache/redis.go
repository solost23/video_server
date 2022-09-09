package cache

import (
	"context"
	"fmt"
	"video_server/config"

	"github.com/go-redis/redis/v8"
)

var RedisMapPool = make(map[int]*redis.Client, 15)

func RedisConnFactory(db int) (*redis.Client, error) {
	redisConfig := config.NewRedisConfig()
	if RedisMapPool[db] != nil {
		return RedisMapPool[db], nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	RedisMapPool[db] = rdb
	return RedisMapPool[db], nil
}

package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"video_server/global"
)

const (
	provider = "consul"
)

func InitConfig(filePath string) {
	v := viper.New()
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

	// 从配置中心读取配置
	err := v.AddRemoteProvider(provider,
		fmt.Sprintf("%s:%d", global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port),
		global.ServerConfig.ConfigPath)
	if err != nil {
		panic(err)
	}

	v.SetConfigType("YAML")

	if err = v.ReadRemoteConfig(); err != nil {
		panic(err)
	}

	if err = v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
}

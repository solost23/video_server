package initialize

import (
	"github.com/spf13/viper"
	"video/global"
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
}

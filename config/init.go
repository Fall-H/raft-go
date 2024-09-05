package config

import (
	"github.com/spf13/viper"
	"log"
)

var GConfig Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AddConfigPath(".\\config")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	viper.Unmarshal(&GConfig)
}

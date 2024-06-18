package config

import (
	"blogServe/business/config"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Init struct{}

func InitConfig(configFilePath string, configType string) config.GlobalConfig {
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}
	if configType == "" {
		configType = "yaml"
	}
	v := viper.New()
	v.Debug()
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		fmt.Println("配置文件不存在")
		// 可以进行其他处理，如使用默认配置
	}
	v.AddConfigPath(".")
	v.SetConfigFile(configFilePath)
	v.SetConfigType(configType)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var confFig config.GlobalConfig
	err = v.Unmarshal(&confFig)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(confFig.Redis.Db)
	fmt.Println(confFig.Redis.Host)
	fmt.Println(confFig.Redis.Port)
	fmt.Println(confFig.Redis.Password)
	return confFig
}

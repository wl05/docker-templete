package config

import (
	"github.com/spf13/viper"
	"os"
	"fmt"
)

func initConfig() (err error) {
	configType := "yaml"
	defaultPath := "conf"
	v := viper.New()
	// 从default中读取默认的配置
	v.SetConfigName("default")
	v.AddConfigPath(defaultPath)
	v.SetConfigType(configType)
	err = v.ReadInConfig()
	if err != nil {
		return
	}

	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}
	env := os.Getenv("GO_ENV")

	// 根据配置的env读取相应的配置信息
	if env != "" {
		viper.SetConfigName(env)
		viper.AddConfigPath(defaultPath)
		viper.SetConfigType(configType)
		err = viper.ReadInConfig()
		if err != nil {
			return
		}
	}
	return
}

func init(){
	err := initConfig()
	if(err != nil){
		fmt.Println("init config failed",err)
		return
	}
}


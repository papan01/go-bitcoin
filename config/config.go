package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//初始化設定檔
func init() {
	setConfig("config", "json", "./config")
}

func setConfig(name string, extension string, path string) {
	viper.SetConfigName(name)
	viper.SetConfigType(extension)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

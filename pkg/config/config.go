package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	AppConfig *Config
)

type Config struct {
	AppName string
	Port    string
	Host    string
}

func InitConfig() {
	if AppConfig == nil {
		viper.SetConfigFile("config.toml")
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Printf("configuration file not found: %v", err)
			} else {
				fmt.Printf("error on reading the configuration file: %v", err)
			}
		}
		viper.AutomaticEnv()

		AppConfig = &Config{
			AppName: viper.GetString("app.name"),
			Port:    viper.GetString("app.port"),
			Host:    viper.GetString("app.host"),
		}
	}
}

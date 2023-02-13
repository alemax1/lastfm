package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("../../config")

	return viper.ReadInConfig()
}

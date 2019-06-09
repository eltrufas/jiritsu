package api

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

func InitConfig() (*Config, error) {
	viper.SetDefault("Port", "8080")
	config := &Config{
		Port: viper.GetString("Port"),
	}
	return config, nil
}

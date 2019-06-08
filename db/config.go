package db

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ConnectionString string
}

func InitConfig() (*Config, error) {
	config := &Config{
		ConnectionString: viper.GetString("ConnectionString"),
	}
	if config.ConnectionString == "" {
		return nil, fmt.Errorf("ConnectionString must be set")
	}
	return config, nil
}
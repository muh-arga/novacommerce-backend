package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.BindEnv("app.name", "APP_NAME"); err != nil {
		return nil, fmt.Errorf("bind app name: %w", err)
	}

	if err := viper.BindEnv("app.port", "APP_PORT"); err != nil {
		return nil, fmt.Errorf("bind app port: %w", err)
	}

	if err := viper.BindEnv("app.env", "APP_ENV"); err != nil {
		return nil, fmt.Errorf("bind app environment: %w", err)
	}

	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("decode config: %w", err)
	}

	return &config, nil
}

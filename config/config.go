package config

import (
	"github.com/adimurianto/go-gin-boilerplate/infra/logger"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// SetupConfig configuration
func SetupConfig() error {
	var configuration *Configuration
	viper.AutomaticEnv()

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logger.Warnf("Tidak ada file .env, menggunakan environment variables")
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Errorf("Error saat decode config, %v", err)
		return err
	}

	return nil
}

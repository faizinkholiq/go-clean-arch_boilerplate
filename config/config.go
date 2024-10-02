package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppPort   string `mapstructure:"app_port"`
	DBHost    string `mapstructure:"db_host"`
	DBUser    string `mapstructure:"db_user"`
	DBPass    string `mapstructure:"db_pass"`
	DBName    string `mapstructure:"db_name"`
	RedisHost string `mapstructure:"redis_host"`
}

var AppConfig *Config

func LoadConfig() error {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")

	err := config.ReadInConfig()
	if err != nil {
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}

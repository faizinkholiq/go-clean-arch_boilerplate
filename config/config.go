package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort   string `json:"app_port"`
	DBHost    string `json:"db_host"`
	DBUser    string `json:"db_user"`
	DBPass    string `json:"db_pass"`
	DBName    string `json:"db_name"`
	RedisHost string `json:"redis_host"`
}

var App *Config

func Load() error {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("/app/config")

	err := config.ReadInConfig()
	if err != nil {
		return err
	}

	if err := viper.Unmarshal(&App); err != nil {
		return err
	}

	log.Println("here is port: ", config.GetString("app_port"))
	log.Println("here is DBHost: ", App.DBHost)

	return nil
}

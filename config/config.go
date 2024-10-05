package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB     ConfigDB     `json:"db"`
	Redis  ConfigRedis  `json:"redis"`
	Server ConfigServer `json:"server"`
}

type ConfigDB struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type ConfigRedis struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type ConfigServer struct {
	Port string `json:"port"`
}

func LoadConfig() (config Config, er error) {
	viperConf := viper.New()

	viperConf.SetConfigName("config")
	viperConf.SetConfigType("json")
	viperConf.AddConfigPath("/app/config")

	err := viperConf.ReadInConfig()
	if err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

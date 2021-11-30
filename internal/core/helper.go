package core

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Config(conf string) string {
	// viper
	viper.AddConfigPath("/app/internal/conf")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return viper.GetString(conf)
}
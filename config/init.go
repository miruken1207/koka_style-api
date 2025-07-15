package config

import (
	"log"

	"github.com/spf13/viper"
)

var JwtSecret []byte

func Init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	JwtSecret = []byte(viper.GetString("jwt_secret"))
}

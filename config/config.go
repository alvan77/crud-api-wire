package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBName string `mapstructure:"DB_NAME"`
}

func LoadConfig() (config Config, err error) {

	viper.SetConfigName(".env")

	viper.SetConfigType("env")

	viper.AddConfigPath("./")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&config)
	return
}
